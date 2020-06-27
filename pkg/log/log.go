package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger builds and returns new Logger
func NewLogger(debug bool) *logrus.Logger {
	var l *logrus.Logger

	hostname, _ := os.Hostname()

	l = logrus.New()
	if debug {
		l.Level = logrus.DebugLevel
	}
	l.Formatter = &Formatter{host: hostname}
	l.Out = os.Stdout

	return l
}
