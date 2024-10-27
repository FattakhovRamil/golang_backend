package logger

import (
	"log"
	"os"
)

type Logger interface {
	LogInfo(message string)
	LogError(message string, err error)
}

type FileLogger struct {
	file *os.File
}

func NewFileLogger(filename string) *FileLogger {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	return &FileLogger{file: file}
}

func (l *FileLogger) LogInfo(message string) {
	log.SetOutput(l.file)
	log.Println("[INFO]:", message)
}

func (l *FileLogger) LogError(message string, err error) {
	log.SetOutput(l.file)
	log.Println("[ERROR]:", message, err)
}
