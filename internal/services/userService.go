package services

import (
	"fmt"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) CreateUser(user *m.UserModel) (uuid.UUID, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Faled to create user: %v", err.Error())
	}
	user.Password = string(hashedPassword)

	if err := s.uow.BeginTransaction(); err != nil {
		return uuid.Nil, fmt.Errorf("Faled to create user: %v", err.Error())
	}

	id, err := s.uow.UserRepo.Insert(user)
	if err != nil {
		s.uow.Rollback()
		return uuid.Nil, fmt.Errorf("Faled to create user: %v", err.Error())
	}
	if err := s.uow.Commit(); err != nil {
		return uuid.Nil, fmt.Errorf("Faled to create user: %v", err.Error())
	}
	return id, nil
}
