package logger

import (
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"encoding/binary"
	"encoding/hex"
	"github.com/gelraen/appengine-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"runtime/debug"
)

// NewLogger returns a new Logger instance
func NewLogger() *logrus.Logger {
	return newLogger(os.Stderr)
}

// NewRequestLogger returns a new Logger instance with http request
func NewRequestLogger(r *http.Request) *logrus.Entry {
	logger := newLogger(os.Stderr)
	return WithRequest(logger, r)
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

// WithRequest add traceID and spanID to entry
func WithRequest(logger *logrus.Logger, r *http.Request) *logrus.Entry {
	var format = propagation.HTTPFormat{}

	sc, ok := format.SpanContextFromRequest(r)

	if !ok {
		logger.Warn("FAILED: SpanContextFromRequest")
		return logger.WithFields(logrus.Fields{})
	}

	return logger.WithFields(logrus.Fields{
		"logging.googleapis.com/trace":  hex.EncodeToString(sc.TraceID[:]),
		"logging.googleapis.com/spanId": binary.BigEndian.Uint64(sc.SpanID[:]),
	})
}
