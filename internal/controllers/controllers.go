package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/vladas9/backend-practice/internal/utils"
	"net/http"
	"reflect"
)

func writeJSON(w http.ResponseWriter, status int, v any) *ApiError {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		aErr := &ApiError{fmt.Sprintf("Encoding of object of type %v failed", reflect.TypeOf(v)), 500}
		utils.Logger.Error(aErr)
		return aErr
	}
	return nil
}

type ApiError struct {
	ErrorMsg string `json:"error"`
	Status   int
}

func (e ApiError) Error() string {
	return e.ErrorMsg
}
