package logger

import (
	"os"
	"strings"
	"time"
)

// Severity tags.
const (
	tagInfo    = "INFO"
	tagWarning = "WARN"
	tagError   = "ERROR"
	tagFatal   = "FATAL"
)

// LogEntry represents a log entry
type LogEntry struct {
	Level     string      `json:"level"`
	Timestamp time.Time   `json:"timestamp"`
	Location  string      `json:"location"`
	Title     string      `json:"title"`
	Keys      interface{} `json:"keys"`
}

type ILogger interface {
	Info(title string, keys ...any)
	Warn(title string, keys ...any)
	Error(title string, keys ...any)
	Fatal(title string, keys ...any)
	Close()
}

type Logger struct {
	filename   string
	isFirstLog bool
}

// NewLogger Function to start the logger and create the logs folder
func NewLogger() *Logger {
	_ = os.MkdirAll("logs", os.ModePerm)
	timeReplace := strings.Replace(time.Now().String(), ":", "-", -1)
	return &Logger{
		filename:   "logs/logs " + timeReplace + ".json",
		isFirstLog: true,
	}
}
