package server

import (
	u "github.com/vladas9/backend-practice/internal/utils"
	"github.com/vladas9/controllers"
	"net/http"
	"reflect"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) *controllers.ApiError

func (fn apiFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if apiErr := fn(w, r); apiErr != nil {
		u.Logger.Error(reflect.ValueOf(apiErr))
		//WriteJSON(w, apiErr.Status, apiErr.Error())
		http.Error(w, apiErr.Error(), apiErr.Status)
	}
}

type UnitWorker interface{} // dummy temp type, delete later
type Server struct {
	ListenAddr string
	Router     *http.ServeMux
	UnitWorker UnitWorker
}

func NewServer(addr string) *Server {

	mux := http.NewServeMux()
	return &Server{
		ListenAddr: addr,
		Router:     mux,
	}
}

func (s *Server) Run() {
	s.Router.Handle("GET /api/lots", apiFunc(controllers.handleGetLotList))
	s.Router.Handle("/api/lot/{id}", apiFunc(controllers.handleLot))
	s.Router.Handle("POST /api/auth", apiFunc(controllers.handleAuth))
	s.Router.Handle("GET /api/user/{id}", apiFunc(controllers.handleUser))
	u.Logger.Info("Registered Routes")

	u.Logger.Info("Started server on", s.ListenAddr)
	u.Logger.Error(http.ListenAndServe(s.ListenAddr, s.Router))
}
