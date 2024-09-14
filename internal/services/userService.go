package services

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

func (s *UserService) CreateUser(user *m.UserModel) (uuid.UUID, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Faled to hash password: %v", err.Error())
	}
	user.Password = string(hashedPassword)

	decodedImage, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(user.Image, "data:image/png;base64,"))
	if err != nil {
		return uuid.Nil, fmt.Errorf("Faled to decode image: %v", err.Error())
	}

	imageUUID := uuid.New().String()

	imagePath := fmt.Sprintf("%s%s.png", ImageDir, imageUUID)
	if err = os.WriteFile(imagePath, decodedImage, 0644); err != nil {
		return uuid.Nil, fmt.Errorf("Faled to save image: %v", err.Error())
	}
	user.Image = imageUUID

	if err := s.uow.BeginTransaction(); err != nil {
		return uuid.Nil, fmt.Errorf("Faled to find user: %v", err.Error())
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

func (s *UserService) CheckUser(user *m.UserModel) (uuid.UUID, error) {
	if err := s.uow.BeginTransaction(); err != nil {
		return uuid.Nil, fmt.Errorf("Failed to start transaction: %v", err.Error())
	}

	storedUser, err := s.uow.UserRepo.GetByEmail(user.Email)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Failed to find user: %v", err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return uuid.Nil, fmt.Errorf("Invalid password: %v", err.Error())
	}

	return storedUser.ID, nil
}
