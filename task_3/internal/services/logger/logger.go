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

func NewLogger(file *os.File) *FileLogger {
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
