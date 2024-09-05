package logger

import (
	"log"
	"os"
)

const (
	InfoLogFile  = "../../log-files/info.log"
	ErrorLogFile = "../../log-files/error.log"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func InitInfoLogger(file *os.File) {
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InitErrorLogger(file *os.File) {
	ErrorLogger = log.New(file, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func init() {
	infoFile, err := os.OpenFile(InfoLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer infoFile.Close()

	errorFile, err := os.OpenFile(ErrorLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

}
