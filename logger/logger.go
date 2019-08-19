package logger

import (
	"github.com/gelraen/appengine-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime/debug"
)

// NewLogger returns a new Logger instance
func NewLogger() *logrus.Logger {
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

	log.SetReportCaller(true)
	log.SetFormatter(&appengine.Formatter{})
	log.SetOutput(out)

	return log
}

// SendError send error to Stackdriver Error Reporting
func SendError(logger *logrus.Logger, err error) {
	logger.Errorf("%v\n\n%s\n\n%+v", err, string(debug.Stack()), err)
}
