package services

import (
	"database/sql"
	"fmt"

	"github.com/vladas9/backend-practice/internal/errors"
	r "github.com/vladas9/backend-practice/internal/repository"
)

var ImageDir = "./public/img/"

type Service struct {
	store *r.Store
	host  string
	port  string
}

type Response map[string]interface{}

type Problems map[string]string

type Validator interface {
	Validate() Problems
}

func NewService(db *sql.DB, host, port string) *Service {
	return &Service{r.NewStore(db), host, port}
}

func (p Problems) toErr() error {
	return errors.NotValid(p, fmt.Errorf("params not valid"))
}
