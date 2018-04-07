package geniuslurker

import (
	"log"
	"os"
)

var (
	// InfoGeniusLogger is an INFO level logger
	InfoGeniusLogger *log.Logger
)

// InitLoggers initializes loggers
func InitLoggers() {
	logFile, _ := os.OpenFile("genius.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	InfoGeniusLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
