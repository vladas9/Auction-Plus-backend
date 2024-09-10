package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SignIn(w http.ResponseWriter, r *http.Request) *ApiError {

	var userInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		return &ApiError{fmt.Sprintf("Failed SignIn: %s", err), http.StatusBadRequest}
	}

	//TODO: call service for this using userInfo and w
	return writeJSON(w, http.StatusOK, data)
}

func SignUp(w http.ResponseWriter, r *http.Request) *ApiError {

	var userInfo struct {
		Email            string    `json:"email"`
		Password         string    `json:"password"`
		Address          string    `json:"address"`
		PhoneNumber      string    `json:"phone_number"`
		UserType         string    `json:"user_type"`
		RegistrationDate time.Time `json:"restered_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		return &ApiError{fmt.Sprintf("Failed SignUp: %s", err), http.StatusBadRequest}
	}

	return writeJSON(w, http.StatusOK, data)
}
