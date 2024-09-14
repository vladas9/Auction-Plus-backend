package services

import (
	"database/sql"
	uw "github.com/vladas9/backend-practice/internal/unitofwork"
)

type UserService struct {
	uow *uw.UnitOfWork
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{uw.NewUnitOfWork(db)}
}

var ImageDir = "./public/img/"
