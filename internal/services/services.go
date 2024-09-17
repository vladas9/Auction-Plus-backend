package services

import (
	"database/sql"
	"fmt"

	r "github.com/vladas9/backend-practice/internal/repository"
)

type Service struct {
	store *r.Store
}

func NewService(db *sql.DB) *Service {
	return &Service{r.NewStore(db)}
}

const ImageDir = "./public/img/"

type ServiceErrorType string

const (
	InternalError   ServiceErrorType = "Internal Error"
	ValidationError                  = "Validation Error"
	RetrievalError                   = "Retrieival Error"
)

type ServiceError struct {
	ErrorType ServiceErrorType
	ErrorMsg  any
}

func (err *ServiceError) Error() string {
	return fmt.Sprint(err.ErrorMsg)
}

func serviceError(errType ServiceErrorType, err error) error {
	if err == nil {
		return nil
	}
	return &ServiceError{
		ErrorType: errType,
		ErrorMsg:  err.Error(),
	}
}
