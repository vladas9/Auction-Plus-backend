package services

import (
	"fmt"

	"github.com/vladas9/backend-practice/internal/errors"
)

var Host, Port string

func InitService(host, port string) {
	Host = host
	Port = port
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
