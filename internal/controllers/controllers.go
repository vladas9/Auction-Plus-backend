// Package controllers provides HTTP handlers for user authentication operations.
// It includes handlers for user sign-in and sign-up processes.
// Each handler reads the request body, decodes it into a Model, and interacts with a service to process the request.
// Based on the service's response, the handlers return appropriate HTTP responses.
package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	s "github.com/vladas9/backend-practice/internal/services"
	u "github.com/vladas9/backend-practice/internal/utils"
)

type Controller struct {
	service *s.Service
}

var Host, Port string
var JwtSecret []byte

func NewController(db *sql.DB) *Controller {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can load env variables")
	}
	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	JwtSecret = []byte(os.Getenv("JWTKEY"))

	return &Controller{s.NewService(db, Host, Port)}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		aErr := fmt.Errorf("Encoding of object of type %v failed", reflect.TypeOf(v))
		u.Logger.Error(aErr)
		return aErr
	}
	return nil
}

type Response map[string]interface{}

func atoi(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	return strconv.Atoi(str)
}

func atodec(str string) (decimal.Decimal, error) {
	if str == "" {
		return decimal.Zero, nil
	}
	return decimal.NewFromString(str)
}
