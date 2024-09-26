package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/vladas9/backend-practice/internal/errors"
	m "github.com/vladas9/backend-practice/internal/models"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) error {
	user := &m.UserModel{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return errors.NotValid(err.Error(), err)
	}

	storedUser, err := c.service.CheckUser(user)
	if err != nil {
		return err
	}

	u.Logger.Info(fmt.Sprintf("User %s logged in", storedUser.Username))

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
		return err
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
		return err
	}

	return WriteJSON(w, http.StatusOK, &Response{
		"img_src":   fmt.Sprintf("http://%s:%s/api/img/%s", Host, Port, user.Image),
		"user_type": user.UserType,
	})
}

func (c *Controller) ImageHandler(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/api/img/")
	imagePath := fmt.Sprintf("%s%s.png", "./public/img/", id)

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return errors.NotFound("Image not found:", err)
	}

	http.ServeFile(w, r, imagePath)
	return nil
}

func createResponse(id, userType, image string) (Response, error) {
	imgSrc := fmt.Sprintf("http://%s:%s/api/img/%s", Host, Port, image)

	token, err := u.GenerateJWT(id, userType, JwtSecret)
	if err != nil {
		return Response{}, err
	}

	return Response{
		"auth_token": token,
		"img_src":    imgSrc,
	}, nil
}
