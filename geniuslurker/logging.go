package geniuslurker

import (
	"log"
	"os"
)

var (
	// InfoLogger is an INFO level logger
	InfoLogger *log.Logger
	// ErrorLogger is an INFO level logger
	ErrorLogger *log.Logger
)

// InitLoggers initializes loggers
func InitLoggers() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
