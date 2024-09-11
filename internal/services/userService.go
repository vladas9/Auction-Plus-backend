package services

import (
	"fmt"

	m "github.com/vladas9/backend-practice/internal/models"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func (s *UserService) CreateUser(user *m.UserModel) error {

	u.Logger.Info("Creating user: ", user.Username)
	if err := s.uow.BeginTransaction(); err != nil {
		return fmt.Errorf("Faled to create user: %v", err.Error())
	}

	if err := s.uow.UserRepo.Insert(user); err != nil {
		s.uow.Rollback()
		return fmt.Errorf("Faled to create user: %v", err.Error())
	}
	if err := s.uow.Commit(); err != nil {
		return fmt.Errorf("Faled to create user: %v", err.Error())
	}
	return nil
}
