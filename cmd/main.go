package main

import (
	"github.com/vladas9/backend-practice/internal/utils"
)

func main() {
	utils.SetupLogger("log-files/logs.log")
	utils.Logger.Info("Hello")
}
