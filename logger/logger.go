package logger

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// NewLogger returns a new Logger instance
func NewLogger() *logrus.Logger {
	return newLogger(os.Stdout)
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

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// WithErrorLocation add error location for Stackdriver Error Reporting
func WithErrorLocation(logger *logrus.Logger, err error) *logrus.Entry {
	errorWithStack, ok := err.(stackTracer)

	if !ok {
		errorWithStack = errors.WithStack(err).(stackTracer)
	}

	frame := errorWithStack.StackTrace()[0]

	sourceFullPath := strings.Split(fmt.Sprintf("%+s", frame), "\n\t")[1]
	lineNumber, _ := strconv.Atoi(fmt.Sprintf("%d", frame))

	return logger.WithFields(logrus.Fields{
		"context": logrus.Fields{
			"reportLocation": logrus.Fields{
				"filePath":     sourceFullPath,
				"lineNumber":   lineNumber,
				"functionName": fmt.Sprintf("%n", frame),
			},
		},
	})
}
