package main

import (
	"net/http"

	"github.com/vladas9/backend-practice/internal/errors"
	"golang.org/x/net/websocket"

	c "github.com/vladas9/backend-practice/internal/controllers"
	u "github.com/vladas9/backend-practice/internal/utils"
	p "github.com/vladas9/backend-practice/pkg/postgres"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func (fn apiFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		u.Logger.Error(err)
		if apiErr, ok := err.(*errors.ApiError); ok {
			c.WriteJSON(w, apiErr.Status, *apiErr)
		} else {
			c.WriteJSON(w, http.StatusInternalServerError, errors.ApiError{
				Status:   http.StatusInternalServerError,
				ErrorMsg: "Internal server error",
			})
		}
	}
}

type Server struct {
	ListenAddr      string
	Router          *http.ServeMux
	EventController c.EventController
}

func NewServer(addr string) *Server {
	err := p.ConnectDB()
	if err != nil {
		u.Logger.Error("connecting db: ", err.Error())
		panic(err)
	}
	mux := http.NewServeMux()

	return &Server{
		ListenAddr:      addr,
		Router:          mux,
		EventController: c.NewEventController(),
	}
}

func (s *Server) Run() {
	//TODO: Add controllers links

	s.Router.Handle("POST /api/user/register", apiFunc(c.Register))
	s.Router.Handle("POST /api/user/login", apiFunc(c.Login))
	s.Router.Handle("GET /api/user/data", apiFunc(c.UserData))
	s.Router.Handle("GET /api/user/profile-data", apiFunc(c.ProfileData))
	s.Router.Handle("GET /api/img/", apiFunc(c.ImageHandler))
	s.Router.Handle("POST /api/bid/post", apiFunc(c.AddBid))
	s.Router.Handle("GET /api/bids/table", apiFunc(c.BidTable))
	s.Router.Handle("GET /api/auctions/table", apiFunc(c.AuctionTable))
	s.Router.Handle("GET /api/auctions/cards", apiFunc(c.GetAuctions))
	s.Router.Handle("GET /api/auction/{id}", apiFunc(c.GetAuction))
	s.Router.Handle("POST /api/auction/post", apiFunc(c.AddAuction))
	s.Router.Handle("/api/auction/ws/{id}", websocket.Handler(s.EventController.AuctionEvents))
	// s.Router.Handle("POST /api/payments", apiFunc(controllers.))
	// s.Router.Handle("GET /api/notifications", apiFunc(controllers.))
	// s.Router.Handle("PUT /api/notifications/{id}/read", apiFunc(controllers.))
	// u.Logger.Info("Registered Routes")

	u.Logger.Info("Started server on", s.ListenAddr)
	u.Logger.Error(http.ListenAndServe(s.ListenAddr, CORS(s.Router)))
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
