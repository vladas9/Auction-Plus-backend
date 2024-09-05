package main

import (
	"github.com/vladas9/backend-practice/internal/utils"
)

func main() {
	utils.SetupLogger()
	utils.Logger.Info("Hello")
}
