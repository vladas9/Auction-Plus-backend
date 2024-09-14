package controllers

import (
	"encoding/json"
	"fmt"
	m "github.com/vladas9/backend-practice/internal/models"
	"net/http"
	"os"
	"strings"
)

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) *ApiError {
	user := &m.UserModel{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return &ApiError{fmt.Sprintf("Decoding failed(login): %s", err), http.StatusBadRequest}
	}

	storedUser, err := c.userService.CheckUser(user)
	if err != nil {
		return &ApiError{fmt.Sprintf("Login failed: %s", err.Error()), http.StatusNotFound}
	}

	return writeJSON(w, http.StatusOK, Response{
		"auth_token": storedUser.ID,
		"img_src":    "/api/img/" + storedUser.Image,
		"user_type":  storedUser.UserType,
	})
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) *ApiError {

	user := &m.UserModel{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return &ApiError{fmt.Sprintf("Decoding failed(Register): %s", err), http.StatusBadRequest}
	}

	storedUser, err := c.userService.CreateUser(user)
	if err != nil {
		return &ApiError{fmt.Sprintf("Registration failed: %s", err.Error()), http.StatusNotAcceptable}
	}

	return writeJSON(w, http.StatusOK, Response{
		"auth_token": storedUser.ID,
		"img_src":    "/api/img/" + storedUser.Image,
		"user_type":  storedUser.UserType,
	})
}

func (c *Controller) ImageHandler(w http.ResponseWriter, r *http.Request) *ApiError {
	id := strings.TrimPrefix(r.URL.Path, "/api/img/")
	imagePath := fmt.Sprintf("%s%s.png", "./public/img/", id)

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return &ApiError{fmt.Sprintf("Image not found: %s", err.Error()), http.StatusNotAcceptable}
	}

	http.ServeFile(w, r, imagePath)
	return nil
}
