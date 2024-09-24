package errors

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	ErrorMsg any `json:"error"`
	Status   int `json:"status"`
}

func (e *ApiError) Error() string {
	return fmt.Sprint(e.ErrorMsg)
}


type ErrMesage interface{}

type Validator interface {
	Validate() error
}

type ErrNotFound 

func NonValid(msg ErrMesage, err error) error {
	return &ApiError{Status: http.StatusBadRequest, ErrorMsg: msg}
}

func NotFound(msg ErrMesage) error {
	return &ApiError{Status: http.StatusNotFound, ErrorMsg: msg}
}

func NotAllowed(msg ErrMesage) error {
	return &ApiError{Status: http.StatusMethodNotAllowed}
}

func Internal(err error) error {
	return fmt.Errorf("Internal Error: %w", err)
}
