package main

import (
	s "github.com/vladas9/backend-practice/cmd/server"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func main() {
	u.SetupLogger("log-files/logs.log")
	server := s.NewServer("localhost:1169")
	server.Run()
}
