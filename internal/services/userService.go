package services

import (
	m "github.com/vladas9/backend-practice/internal/models"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func CreateUser(user *m.UserModel) error {
	u.Logger.Info("Creating user: ", user.Username)
}
