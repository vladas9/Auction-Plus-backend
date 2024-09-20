package services

import (
	"database/sql"

	r "github.com/vladas9/backend-practice/internal/repository"
)

type Service struct {
	store *r.Store
}

var Host, Port string

func NewService(db *sql.DB, host, port string) *Service {
	Host = host
	Port = port
	return &Service{r.NewStore(db)}
}

type Response map[string]interface{}

type Problems map[string]string

type Validator interface {
	Validate() Problems
}

var ImageDir = "./public/img/"
