package utils

import (
	"log"
	"os"
)

const (
	LogFile = "../../log-files/logs.log"
)

type LoggerType struct {
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

var Logger *LoggerType

func NewLogger(file *os.File) *LoggerType {
	return &LoggerType{
		info:  log.New(file, "INFO: ", log.LstdFlags),
		warn:  log.New(file, "WARN: ", log.LstdFlags),
		error: log.New(file, "ERROR: ", log.LstdFlags),
	}
}

func (l *LoggerType) Info(v ...any) {
	l.info.Println(v...)
}

func (l *LoggerType) Warn(v ...any) {
	l.warn.Println(v...)
}

func (l *LoggerType) Error(v ...any) {
	l.error.Println(v...)
}

func SetupLogger(path string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	Logger = NewLogger(file)
}
