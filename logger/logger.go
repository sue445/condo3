package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

// NewLogger returns a new Logger instance
func NewLogger() *logrus.Logger {
	return newLogger(os.Stdout)
}

// NewErrorLogger returns a new Logger instance for error logging
func NewErrorLogger() *logrus.Logger {
	return newLogger(os.Stderr)
}

func newLogger(out io.Writer) *logrus.Logger {
	log := logrus.New()

	if os.Getenv("LOG_LEVEL") == "" {
		log.Level = logrus.DebugLevel
	} else {
		level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))

		if err != nil {
			panic(err)
		}

		log.Level = level
	}

	// c.f. https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = out

	return log
}
