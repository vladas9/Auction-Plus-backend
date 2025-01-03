package models

import (
	//"github.com/google/uuid"
	"time"
)

type UserModel struct {
	BaseModel
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Image          string    `json:"img_src"`
	Address        string    `json:"address"`
	PhoneNumber    string    `json:"phone_number"`
	UserType       string    `json:"user_type"`
	RegisteredDate time.Time `json:"registered_date"`
}
