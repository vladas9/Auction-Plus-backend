package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can load env variables")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	u.SetupLogger("log-files/logs.log")
	server := NewServer(fmt.Sprintf("%s:%s", host, port))
	server.Run()
}
