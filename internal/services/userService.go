package services

import (
	"github.com/google/uuid"
	"github.com/vladas9/backend-practice/internal/errors"
	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func (s *Service) CreateUser(user *m.UserModel) (*m.UserModel, error) {
	var err error

	if user.Password, err = u.HashPassword(user.Password); err != nil {
		return nil, errors.Internal(err)
	}

	imageUUID := uuid.New().String()

	if err = u.DecodeAndSaveImage(user.Image, ImageDir, imageUUID); err != nil {
		return nil, errors.Internal(err)
	}

	user.Image = imageUUID

	err = s.store.WithTx(func(stx *r.StoreTx) error {
		user.ID, err = stx.UserRepo().Insert(user)
		return err
	})

	if err != nil {
		return nil, errors.Conflict("Account already exists", err)
	}

	return user, nil
}

func (s *Service) CheckUser(user *m.UserModel) (storedUser *m.UserModel, err error) {
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		storedUser, err = stx.UserRepo().GetByEmail(user.Email)
		return err
	})

	if err != nil {
		return nil, errors.NotFound("Account with given email does not exist", err)
	}

	if err = u.CompareHashPassword(user.Password, storedUser.Password); err != nil {
		return nil, errors.Unauthorized("Email or password is wrong", err)
	}

	return storedUser, nil
}

func (s *Service) GetUserData(id uuid.UUID) (storedUser *m.UserModel, err error) {
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		storedUser, err = stx.UserRepo().GetById(id)
		return err
	})
	if err != nil {
		return nil, errors.Internal(err)
	}

	return storedUser, nil
}
