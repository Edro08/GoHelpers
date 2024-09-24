package logger

import "os"

func (l *Logger) Info(title string, keys ...any) {
	l.writeLog(tagInfo, title, keys)
}

func (l *Logger) Warn(title string, keys ...any) {
	l.writeLog(tagWarning, title, keys)
}

func (l *Logger) Error(title string, keys ...any) {
	l.writeLog(tagError, title, keys)
}

func (l *Logger) Fatal(title string, keys ...any) {
	l.writeLog(tagFatal, title, keys)
	os.Exit(1)
}

// Close Function to add "]" to the end of the file before shutting down the service
func (l *Logger) Close() {
	fileHandle, _ := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if _, err := fileHandle.WriteString("\n]"); err != nil {
	}
	_ = fileHandle.Close()
}
