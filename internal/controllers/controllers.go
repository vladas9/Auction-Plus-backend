// Package controllers provides HTTP handlers for user authentication operations.
// It includes handlers for user sign-in and sign-up processes.
// Each handler reads the request body, decodes it into a Model, and interacts with a service to process the request.
// Based on the service's response, the handlers return appropriate HTTP responses.
package controllers

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	s "github.com/vladas9/backend-practice/internal/services"
)

type Controller struct {
	service   *s.Service
	host      string
	port      string
	jwtSecret []byte
}

type Response map[string]interface{}

func NewController(db *sql.DB) *Controller {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot load env variables")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	jwtSecret := []byte(os.Getenv("JWTKEY"))

	return &Controller{
		service:   s.NewService(db, host, port),
		host:      host,
		port:      port,
		jwtSecret: jwtSecret,
	}
}
