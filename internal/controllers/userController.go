package controllers

import (
	"encoding/json"
	"fmt"
	m "github.com/vladas9/backend-practice/internal/models"
	u "github.com/vladas9/backend-practice/internal/utils"
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

	u.Logger.Info(fmt.Sprintf("User %s loged in", storedUser.Username))

	token, err := createResponse(storedUser.ID.String(), storedUser.UserType, storedUser.Image)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, Response{
		"auth_token": token,
	})
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) error {

	user := &m.UserModel{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return fmt.Errorf("Decoding failed(Register): %w", err)
	}

	storedUser, err := c.service.CreateUser(user)
	if err != nil {
		return fmt.Errorf("Registration failed: %w", err)
	}

	u.Logger.Info(fmt.Sprintf("User %s created", storedUser.Username))

	token, err := createResponse(storedUser.ID.String(), storedUser.UserType, storedUser.Image)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, Response{
		"auth_token": token,
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

func createResponse(id, userType, image string) (Response, error) {
	imgSrc := fmt.Sprintf("http://%s:%s/api/img/%s", Host, Port, image)

	token, err := u.GenerateJWT(id, userType, imgSrc, JwtSecret)
	if err != nil {
		return Response{}, fmt.Errorf("Token generation failed: %w", err)
	}

	return Response{
		"auth_token": token,
	}, nil
}
