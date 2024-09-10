package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/vladas9/backend-practice/internal/models"
	"net/http"
)

//DOCS
//Sign In Controller that reads body data from front and Decode it in userInfo
//Than it send userInfo to service and wait for responce
//Based on responce it will send a status ok or error

func SignIn(w http.ResponseWriter, r *http.Request) *ApiError {

	var user models.UserModel

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return &ApiError{fmt.Sprintf("Failed SignIn: %s", err), http.StatusBadRequest}
	}

	//TODO: call service for this using user
	return writeJSON(w, http.StatusOK, "Sign in successfully")
}

//DOCS
//Sign Up Controller that reads body data from front and Decode it in userInfo
//Than it send userInfo to service and wait for responce
//Based on responce it will send a status ok or error  with a message

func SignUp(w http.ResponseWriter, r *http.Request) *ApiError {

	var user models.UserModel

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return &ApiError{fmt.Sprintf("Failed SignUp: %s", err), http.StatusBadRequest}
	}

	//TODO: call service for this using user

	return writeJSON(w, http.StatusOK, "Sing in successfully")
}
