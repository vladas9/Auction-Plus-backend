package services

import (
	"database/sql"

	r "github.com/vladas9/backend-practice/internal/repository"
)

type Service struct {
	store *r.Store
}

func NewService(db *sql.DB) *Service {
	return &Service{r.NewStore(db)}
}

var ImageDir = "./public/img/"
