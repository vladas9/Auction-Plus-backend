package services

import (
	"fmt"

	"github.com/vladas9/backend-practice/internal/errors"
	r "github.com/vladas9/backend-practice/internal/repository"
	p "github.com/vladas9/backend-practice/pkg/postgres"
)

type Service struct {
	store *r.Store
}

var Host, Port string

func NewService(host, port string) *Service {
	Host = host
	Port = port
	return &Service{r.NewStore(p.DB)}
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
