package models

import (
	"github.com/google/uuid"
	"time"
)

type UserModel struct {
	UserId         uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"-"`
	Address        string    `json:"address"`
	PhoneNumber    string    `json:"phone_number"`
	UserType       string    `json:"user_type"`
	RegisteredDate time.Time `json:"registered_date"`
}
