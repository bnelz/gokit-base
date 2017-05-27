package logger

import (
	"encoding/json"
	"io"
	"os"
	"strconv"

	gklog "github.com/go-kit/kit/log"
)

// herbertLogger is a logging service wrapping gokit and custom logging logic
type herbertLogger struct {
	log    gklog.Logger
	writer io.Writer
	level  LogLevel
}

// LogLevel represents the atreides logging level
type LogLevel int

const (
	VERBOSE LogLevel = iota + 1
	ERROR
)

// NewHerbertFormatLogger returns a wrapped gokit logger
func NewHerbertFormatLogger(log gklog.Logger, file string, level LogLevel) gklog.Logger {
	f, _ := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	return &herbertLogger{log: log, writer: f, level: level}
}

// Log writes a log entry
func (l *herbertLogger) Log(keyvals ...interface{}) error {

	var logData map[string]string
	logData = make(map[string]string)

	logData["channel"] = "golang"

	var hasError = false
	for i := 0; i < len(keyvals); i += 2 {
		key := keyvals[i].(string)
		val := keyvals[i+1]
		switch val.(type) {
		case string:
			logData[key] = val.(string)
		case int:
			logData[key] = strconv.Itoa(val.(int))
		case bool:
			logData[key] = strconv.FormatBool(val.(bool))
		case error:
			hasError = true
			logData[key] = val.(error).Error()
		}
	}

	serialized, _ := json.Marshal(logData)

	if l.level == VERBOSE || (l.level == ERROR && hasError) {
		l.writer.Write(append(serialized, '\n'))
	}

	return l.log.Log(keyvals...)
}
