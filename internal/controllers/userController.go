package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
	u "github.com/vladas9/backend-practice/internal/utils"
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

	response, err := createResponse(storedUser.ID.String(), storedUser.UserType, storedUser.Image)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, response)
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

	response, err := createResponse(storedUser.ID.String(), storedUser.UserType, storedUser.Image)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, response)
}

func (c *Controller) UserData(w http.ResponseWriter, r *http.Request) error {
	userId, err := u.ExtractUserIDFromToken(r, JwtSecret)
	if err != nil {
		return err
	}
	user, err := c.service.GetUserData(userId)
	if err != nil {
		return fmt.Errorf("Failed getting user data: %s", err)
	}

	return WriteJSON(w, http.StatusOK, &Response{
		"img_src":   fmt.Sprintf("http://%s:%s/api/img/%s", Host, Port, user.Image),
		"user_type": user.UserType,
	})
}

func (c *Controller) ProfileData(w http.ResponseWriter, r *http.Request) error {
	// userId, err := u.ExtractUserIDFromToken(r, JwtSecret)
	// if err != nil {
	// 	return err
	// }

	userId, err := uuid.Parse("44b23f51-b8b4-4d73-aec8-aa1b3930d923")
	if err != nil {
		return fmt.Errorf("Failed to parse UUID: %v", err)
	}

	// Get user data
	user, err := c.service.GetUserData(userId)
	if err != nil {
		return fmt.Errorf("failed getting user data: %s", err)
	}

	// Get user statistics
	stats, err := c.service.GetUserStats(userId)
	if err != nil {
		return fmt.Errorf("failed to get user stats: %v", err)
	}

	return WriteJSON(w, http.StatusOK, &Response{
		"img_src":       fmt.Sprintf("http://%s:%s/api/img/%s", Host, Port, user.Image),
		"username":      user.Username,
		"email":         user.Email,
		"phone_number":  user.PhoneNumber,
		"creation_date": user.RegisteredDate,
		"stats":         stats,
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

	token, err := u.GenerateJWT(id, userType, JwtSecret)
	if err != nil {
		return Response{}, fmt.Errorf("Token generation failed: %w", err)
	}

	return Response{
		"auth_token": token,
		"img_src":    imgSrc,
	}, nil
}
