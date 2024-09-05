package models

type User struct {
  ID int64 `json:"id"`
  Username string `json:"username"` 
  Email string `json:"email"`
  Password string `json:"-"`
  Address string `json:""`
  PhoneNumber string `json:"phone_number"`
  UserType string `json:"user_type"`
  RegisteredDate string `json:"registered_date"`
}
