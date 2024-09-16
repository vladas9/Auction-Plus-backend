package controllers

import (
	"encoding/json"
	"fmt"
	m "github.com/vladas9/backend-practice/internal/models"
	"net/http"
	"os"
	"strings"
)

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) error {
	user := &m.UserModel{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return fmt.Errorf("Decoding failed(Login): %w", err)
	}

	storedUser, err := c.service.CheckUser(user)
	if err != nil {
		return fmt.Errorf("Login failed: %w", err)
	}

	return WriteJSON(w, http.StatusOK, Response{
		"auth_token": storedUser.ID,
		"img_src":    fmt.Sprintf("http://%s:%s/api/img/%s", Host, Port, storedUser.Image),
		"user_type":  storedUser.UserType,
	})
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) error {

	user := &m.UserModel{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return fmt.Errorf("Decoding failed(Register): %w", err)
	}

	// TODO: validate fields
	storedUser, err := c.service.CreateUser(user)
	if err != nil {
		return fmt.Errorf("Registration failed: %w", err)
	}

	return WriteJSON(w, http.StatusOK, Response{
		"auth_token": storedUser.ID,
		"img_src":    fmt.Sprintf("http://%s:%s/api/img/%s", Host, Port, storedUser.Image),
		"user_type":  storedUser.UserType,
	})
}

func (c *Controller) ImageHandler(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/api/img/")
	imagePath := fmt.Sprintf("%s%s.png", "./public/img/", id)

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return &ApiError{"Image not found:", http.StatusNotFound}
	}

	http.ServeFile(w, r, imagePath)
	return nil
}
