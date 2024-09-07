package server

import (
	"encoding/json"
	"fmt"
	u "github.com/vladas9/backend-practice/internal/utils"
	"net/http"
	"reflect"
)

func WriteJSON(w http.ResponseWriter, status int, v any) *apiError {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		aErr := &apiError{fmt.Sprintf("Encoding of object of type %v failed", reflect.TypeOf(v)), 500}
		u.Logger.Error(aErr)
		return aErr
	}
	return nil
}

type apiFunc func(w http.ResponseWriter, r *http.Request) *apiError

func (fn apiFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if apiErr := fn(w, r); apiErr != nil {
		u.Logger.Error(reflect.ValueOf(apiErr))
		//WriteJSON(w, apiErr.Status, apiErr.Error())
		http.Error(w, apiErr.Error(), apiErr.Status)
	}
}

type apiError struct {
	ErrorMsg string `json:"error"`
	Status   int
}

func (e apiError) Error() string {
	return e.ErrorMsg
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
	s.Router.Handle("GET /api/lots", apiFunc(s.handleGetLotList))
	s.Router.Handle("/api/lot/{id}", apiFunc(s.handleLot))
	s.Router.Handle("POST /api/auth", apiFunc(s.handleAuth))
	s.Router.Handle("GET /api/user/{id}", apiFunc(s.handleUser))
	u.Logger.Info("Registered Routes")

	u.Logger.Info("Started server on", s.ListenAddr)
	u.Logger.Error(http.ListenAndServe(s.ListenAddr, s.Router))
}

// Handlers

// Returns a list for the main page
func (s *Server) handleGetLotList(w http.ResponseWriter, r *http.Request) *apiError {
	return WriteJSON(w, http.StatusOK, generateDummyAuctions())
}

// handles different operations on lots
func (s *Server) handleLot(w http.ResponseWriter, r *http.Request) *apiError {
	idStr := r.PathValue("id")
	_ = idStr // do something with the id
	switch r.Method {
	case http.MethodGet:
		// GetLot controller
		return WriteJSON(w, http.StatusOK, generateDummyAuctions()[0])
	case http.MethodPost:
		// ModifyLot controller
		fallthrough
	case http.MethodPut:
		// AddLot controller
		fallthrough
	case http.MethodDelete:
		// DeleteLot controller
		return &apiError{fmt.Sprintf("Method %v not implemented", r.Method), http.StatusNotImplemented} // use auth controller
	default:
		return &apiError{fmt.Sprintf("Method %v not supported", r.Method), http.StatusMethodNotAllowed}
	}
}

// handles authentication via POST
func (s *Server) handleAuth(w http.ResponseWriter, r *http.Request) *apiError {
	// TODO: extract values from form data, error checking, etc...
	return &apiError{"Authentication not implemented", http.StatusNotImplemented} // use auth controller
}

func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) *apiError {
	idStr := r.PathValue("id")
	_ = idStr
	return WriteJSON(w, http.StatusOK, generateDummyUser())
}
