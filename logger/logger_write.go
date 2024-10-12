package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"
)

type LogEntry struct {
	Level     string      `json:"level"`
	Timestamp time.Time   `json:"timestamp"`
	Location  string      `json:"location"`
	Title     string      `json:"title"`
	Keys      interface{} `json:"keys"`
}

func (l *Logger) write(level string, title string, keys ...any) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
	}

	e := LogEntry{
		Level:     level,
		Timestamp: time.Now(),
		Location:  fmt.Sprintf("%s:%d", file, line),
		Title:     title,
		Keys:      keys,
	}

	j, err := json.Marshal(e)
	if err != nil {
		return
	}

	fmt.Println(string(j))

	if l.createFile {
		l.writeInFile(j)
	}
}

func (l *Logger) writeInFile(jsonData []byte) {
	fileHandle, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	defer func(fileHandle *os.File) {
		_ = fileHandle.Close()
	}(fileHandle)

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

	_, err = fileHandle.WriteString(fmt.Sprintf("%s", jsonData))
}
