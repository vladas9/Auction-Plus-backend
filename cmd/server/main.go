package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	//repo "github.com/vladas9/backend-practice/internal/repository"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
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

type UnitWorker interface{}
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
	s.Router.HandleFunc("/api/getalllots", makeHandlerFunc(s.handleGetAllLots))
	s.Router.HandleFunc("/api/getlotbyid", makeHandlerFunc(s.handleGetLotById))
	s.Router.HandleFunc("/api/deletelot", makeHandlerFunc(s.handleDeleteLot))
	s.Router.HandleFunc("/api/createlot", makeHandlerFunc(s.handleCreateLot))
	u.Logger.Info("Registered Routes")

	u.Logger.Info("Started server on", s.ListenAddr)
	u.Logger.Error(http.ListenAndServe(s.ListenAddr, s.Router))
}

func main() {
	u.SetupLogger("log-files/logs.log")
	server := NewServer("localhost:1169")
	server.Run()
}

// Handlers
func (s *Server) handleGetAllLots(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("There are no lots to show")
}
func (s *Server) handleGetLotById(w http.ResponseWriter, r *http.Request) error {
	_ = w
	return fmt.Errorf("No lot with id '%s'", r.URL.Query().Get("id"))
}
func (s *Server) handleCreateLot(w http.ResponseWriter, r *http.Request) error {
	_ = w
	_ = r
	return fmt.Errorf("This is not implemented yet")
}
func (s *Server) handleDeleteLot(w http.ResponseWriter, r *http.Request) error {
	_ = w
	_ = r
	return fmt.Errorf("This is not implemented yet")
}
