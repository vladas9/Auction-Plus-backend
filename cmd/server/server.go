package server

import (
	"net/http"

	c "github.com/vladas9/backend-practice/internal/controllers"
	u "github.com/vladas9/backend-practice/internal/utils"
	p "github.com/vladas9/backend-practice/pkg/postgres"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func (fn apiFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		u.Logger.Error(err)
		if apiErr, ok := err.(*c.ApiError); ok {
			c.WriteJSON(w, apiErr.Status, *apiErr)
		} else {
			c.WriteJSON(w, http.StatusInternalServerError, c.ApiError{
				Status:   http.StatusInternalServerError,
				ErrorMsg: "Internal server error",
			})
		}
	}
}

type Server struct {
	ListenAddr  string
	Router      *http.ServeMux
	Controllers *c.Controller
}

func NewServer(addr string) *Server {
	db, err := p.ConnectDB()
	if err != nil {
		u.Logger.Error("connecting db: ", err.Error())
	}
	mux := http.NewServeMux()

	controller := c.NewController(db)

	return &Server{
		ListenAddr:  addr,
		Router:      mux,
		Controllers: controller,
	}
}

func (s *Server) Run() {
	//TODO: Add controllers links

	s.Router.Handle("POST /api/users/register", apiFunc(s.Controllers.Register))
	s.Router.Handle("POST /api/users/login", apiFunc(s.Controllers.Login))
	s.Router.Handle("GET /api/img/", apiFunc(s.Controllers.ImageHandler))
	// s.Router.Handle("GET /api/users/{id}", apiFunc(controllers.))
	// s.Router.Handle("PUT /api/users/{id}", apiFunc(controllers.))
	// s.Router.Handle("GET /api/items", apiFunc(controllers.))
	// s.Router.Handle("GET /api/items/{id}", apiFunc(controllers.))
	// s.Router.Handle("POST /api/items", apiFunc(controllers.))
	// s.Router.Handle("GET /api/auctions", apiFunc(controllers.))
	// s.Router.Handle("POST /api/auctions", apiFunc(controllers.))
	// s.Router.Handle("GET /api/auctions/{id}", apiFunc(controllers.))
	// s.Router.Handle("PUT /api/auction/{id}/bid", apiFunc(controllers.))
	// s.Router.Handle("", apiFunc(controllers.)) //TODO: WebSocket end point
	// s.Router.Handle("POST /api/bids", apiFunc(controllers.))
	// s.Router.Handle("POST /api/payments", apiFunc(controllers.))
	// s.Router.Handle("GET /api/notifications", apiFunc(controllers.))
	// s.Router.Handle("PUT /api/notifications/{id}/read", apiFunc(controllers.))
	// u.Logger.Info("Registered Routes")

	u.Logger.Info("Started server on", s.ListenAddr)
	u.Logger.Error(http.ListenAndServe(s.ListenAddr, s.Router))
}
