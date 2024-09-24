package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"
)

// writeLog writes a log with title and N values any
func (l *Logger) writeLog(typeLog string, title string, keys ...any) {
	// Get the location of the file that calls the logger
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
	}

	entry := LogEntry{
		Level:     typeLog,
		Timestamp: time.Now(),
		Location:  fmt.Sprintf("%s:%d", file, line),
		Title:     title,
		Keys:      keys,
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		return
	}

	// Print the log to the console
	fmt.Println(string(jsonData))

	fileHandle, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	// Check if isFirstLog
	if l.isFirstLog {
		if _, err := fileHandle.WriteString("[\n"); err != nil {
			return
		}
		l.isFirstLog = false
	} else {
		if _, err := fileHandle.WriteString(",\n"); err != nil {
			return
		}
	}

	defer func(fileHandle *os.File) {
		_ = fileHandle.Close()
	}(fileHandle)

	// Write JSON record to file
	_, err = fileHandle.WriteString(fmt.Sprintf("%s", jsonData))
}
