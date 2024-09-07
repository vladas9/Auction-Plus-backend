package server

import (
	"encoding/json"
	"fmt"
	u "github.com/vladas9/backend-practice/internal/utils"
	"net/http"
	"reflect"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error
type apiError struct {
	Error string `json:"error"`
}

func makeHandlerFunc(f apiFunc) http.HandlerFunc {
	u.Logger.Info("Created handler from ", reflect.TypeOf(f))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			apiErr := apiError{Error: err.Error()}
			u.Logger.Error(reflect.ValueOf(apiErr))
			WriteJSON(w, http.StatusAccepted, apiErr)
		}
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
	s.Router.HandleFunc("GET /api/lots", makeHandlerFunc(s.handleGetLotList))
	s.Router.HandleFunc("/api/lot/{id}", makeHandlerFunc(s.handleLot))
	s.Router.HandleFunc("POST /api/auth", makeHandlerFunc(s.handleAuth))
	s.Router.HandleFunc("GET /api/user/{id}", makeHandlerFunc(s.handleUser))
	u.Logger.Info("Registered Routes")

	u.Logger.Info("Started server on", s.ListenAddr)
	u.Logger.Error(http.ListenAndServe(s.ListenAddr, s.Router))
}

// Handlers

// Returns a list for the main page
func (s *Server) handleGetLotList(w http.ResponseWriter, r *http.Request) error { //return WriteJSON(w http.ResponseWriter, status int, v any)

	return WriteJSON(w, http.StatusOK, generateDummyAuctions())
}

// handles different operations on lots
func (s *Server) handleLot(w http.ResponseWriter, r *http.Request) error { //return WriteJSON(w http.ResponseWriter, status int, v any)
	idStr := r.PathValue("id")
	_ = idStr // do something with the id
	switch r.Method {
	case http.MethodGet:
		// GetLot controller
		return WriteJSON(w, http.StatusOK, generateDummyAuctions()[0])
	case http.MethodPost:
		// ModifyLot controller
	case http.MethodPut:
		// AddLot controller
	case http.MethodDelete:
		// DeleteLot controller
	default:
		return fmt.Errorf("Method %v not supported", r.Method)
	}
	return nil
}

// handles authentication via POST
func (s *Server) handleAuth(w http.ResponseWriter, r *http.Request) error {
	// TODO: extract values from form data, error checking, etc...
	return fmt.Errorf("Auth not implemented") // use auth controller
}

func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) error {
	idStr := r.PathValue("id")
	_ = idStr
	return WriteJSON(w, http.StatusOK, generateDummyUser())
}
