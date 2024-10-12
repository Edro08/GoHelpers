package logger

import (
	"os"
	"strings"
	"time"
)

type Logger struct {
	createFile bool
	filename   string
	isFirstLog bool
}

func NewLogger(createFile bool) *Logger {
	l := &Logger{
		createFile: createFile,
		isFirstLog: true,
	}
	if l.createFile {
		_ = os.MkdirAll("logs", os.ModePerm)
		timeReplace := strings.Replace(time.Now().String(), ":", "-", -1)
		l.filename = "logs/logs " + timeReplace + ".json"
	}
	return l
}

func (l *Logger) Close() {
	if l.createFile {
		fileHandle, _ := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		_, _ = fileHandle.WriteString("\n]")
		_ = fileHandle.Close()
	}
}
