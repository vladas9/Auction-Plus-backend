package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/vladas9/backend-practice/internal/models"
)

// Login handles the HTTP request for user sign-in.
// It reads and decodes the request body into a UserModel instance.
// It then calls a service with the user information and returns an appropriate response based on the service result.
// If decoding fails, it returns a 400 Bad Request status with an error message.
// If successful, it returns a 200 OK status with a success message.
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) *ApiError {
	user := &m.UserModel{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return &ApiError{fmt.Sprintf("Decoding failed(login): %s", err), http.StatusBadRequest}
	}

	id, err := c.userService.CheckUser(user)
	if err != nil {
		return &ApiError{fmt.Sprintf("Login failed: %s", err.Error()), http.StatusNotFound}
	}

	return writeJSON(w, http.StatusOK, Response{"id": id})
}

// Register	handles the HTTP request for user sign-up.
// It reads and decodes the request body into a UserModel instance.
// It then calls a service with the user information and returns an appropriate response based on the service result.
// If decoding fails, it returns a 400 Bad Request status with an error message.
// If successful, it returns a 200 OK status with a success message.
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) *ApiError {

	user := &m.UserModel{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return &ApiError{fmt.Sprintf("Decoding failed(Register): %s", err), http.StatusBadRequest}
	}

	id, err := c.userService.CreateUser(user)
	if err != nil {
		return &ApiError{fmt.Sprintf("Registration failed: %s", err.Error()), http.StatusNotAcceptable}
	}

	return writeJSON(w, http.StatusOK, Response{
		"id": id,
	})
}
