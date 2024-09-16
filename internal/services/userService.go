package services

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
	u "github.com/vladas9/backend-practice/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateUser(user *m.UserModel) (*m.UserModel, error) {
	var err error

	if user.Password, err = u.HashPassword(user.Password); err != nil {
		return nil, err
	}

	if err != nil {
	}

	imageUUID := uuid.New().String()

	if err = u.DecodeAndSaveImage(user.Image, ImageDir, imageUUID); err != nil {
		return nil, err
	}

	user.Image = imageUUID

	err = s.store.WithTx(func(stx *r.StoreTx) error {
		user.ID, err = stx.UserRepo().Insert(user)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("Faled to find user: %v", err.Error())
	}

	return user, nil
}

func (s *Service) CheckUser(user *m.UserModel) (storedUser *m.UserModel, err error) {
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		storedUser, err = stx.UserRepo().GetByEmail(user.Email)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to find user: %v", err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return nil, fmt.Errorf("Invalid password: %v", err.Error())
	}

	return storedUser, nil
}
