package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//DOCS
//Sign In Controller that reads body data from front and Decode it in userInfo
//Than it send userInfo to service and wait for responce
//Based on responce it will send a status ok or error

func SignIn(w http.ResponseWriter, r *http.Request) *ApiError {

	var userInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		return &ApiError{fmt.Sprintf("Failed SignIn: %s", err), http.StatusBadRequest}
	}

	//TODO: call service for this using userInfo
	return writeJSON(w, http.StatusOK, data)
}

//DOCS
//Sign Up Controller that reads body data from front and Decode it in userInfo
//Than it send userInfo to service and wait for responce
//Based on responce it will send a status ok or error

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

	//TODO: call service for this using userInfo

	return writeJSON(w, http.StatusOK, data)
}
