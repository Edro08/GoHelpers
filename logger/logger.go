package logger

import (
	"os"
)

// ILogger
// ------------------------------------------------------------------------------------------------
type ILogger interface {
	Info(title string, keys ...any)
	Warn(title string, keys ...any)
	Error(title string, keys ...any)
	Fatal(title string, keys ...any)
}

// Implementation Methods
// ------------------------------------------------------------------------------------------------
const (
	tagInfo    = "INFO"
	tagWarning = "WARN"
	tagError   = "ERROR"
	tagFatal   = "FATAL"
)

func (l *Logger) Info(title string, keys ...any) {
	l.write(tagInfo, title, keys)
}

func (l *Logger) Warn(title string, keys ...any) {
	l.write(tagWarning, title, keys)
}

func (l *Logger) Error(title string, keys ...any) {
	l.write(tagError, title, keys)
}

func (l *Logger) Fatal(title string, keys ...any) {
	l.write(tagFatal, title, keys)
	os.Exit(1)
}
