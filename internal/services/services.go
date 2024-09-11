package services

import (
	"database/sql"
	uw "github.com/vladas9/backend-practice/internal/unitofwork"
)

type UserService struct {
	unitW *uw.UnitOfWork
}

func NewUserService(db *sql.DB) (*UserService, error) {
	unit, err := uw.NewUnitOfWork(db)
	if err != nil {
		return nil, err
	}
	return &UserService{unit}, nil
}
