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
	return &Service{r.NewStore(db)}
}

type Response map[string]interface{}

var ImageDir = "./public/img/"
