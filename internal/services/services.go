package services

import (
	"database/sql"
	"fmt"

	"github.com/vladas9/backend-practice/internal/errors"
	r "github.com/vladas9/backend-practice/internal/repository"
)

type Service struct {
	store *r.Store
}

var Host, Port string

func NewService(db *sql.DB, host, port string) *Service {
	Host = host
	Port = port
	return &Service{r.NewStore(db)}
}

type Response map[string]interface{}

type Problems map[string]string

type Validator interface {
	Validate() Problems
}

func (p Problems) toErr() error {
	return errors.NotValid(p, fmt.Errorf("params not valid"))
}

var ImageDir = "./public/img/"
