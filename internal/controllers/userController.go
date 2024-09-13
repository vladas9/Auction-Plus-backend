package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
	"github.com/vladas9/backend-practice/internal/utils"
)

// Login handles the HTTP request for user sign-in.
// It reads and decodes the request body into a UserModel instance.
// It then calls a service with the user information and returns an appropriate response based on the service result.
// If decoding fails, it returns a 400 Bad Request status with an error message.
// If successful, it returns a 200 OK status with a success message.
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) *ApiError {
	var user *m.UserModel

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return &ApiError{fmt.Sprintf("SignIn failed: %s", err), http.StatusBadRequest}
	}

	err := c.userService.CreateUser(user)

	if err != nil {
		//TODO: error
	}

	return writeJSON(w, http.StatusOK, "Sign-in successful")
}

// Register	handles the HTTP request for user sign-up.
// It reads and decodes the request body into a UserModel instance.
// It then calls a service with the user information and returns an appropriate response based on the service result.
// If decoding fails, it returns a 400 Bad Request status with an error message.
// If successful, it returns a 200 OK status with a success message.
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) *ApiError {

	user := &m.UserModel{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return &ApiError{fmt.Sprintf("Registration failed: %s", err.Error()), http.StatusBadRequest}
	}

	user.ID = uuid.New()
	user.RegisteredDate = time.Now()
	utils.Logger.Info(user.RegisteredDate)

	if err := c.userService.CreateUser(user); err != nil {
		return &ApiError{fmt.Sprintf("Registration failed: %s", err.Error()), http.StatusNotAcceptable}
	}

	return writeJSON(w, http.StatusOK, "Sign-up successful")
}
