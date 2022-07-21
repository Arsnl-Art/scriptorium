package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// ILogger represent microservice application
// standard logger interface
type ILogger interface {
	logrus.FieldLogger
	SetLevel(level string) error
}

// Logger represent standard logger structure
type Logger struct {
	*logrus.Logger
}

// NewLogger create new logger instance
func NewLogger() ILogger {
	logger := logrus.New()
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}

	logger.SetFormatter(formatter)
	return &Logger{logger}
}

// SetLevel set logger instance. Could be Trace,
// Debug, Info, Warning, Error, Fatal and Panic
func (l *Logger) SetLevel(level string) error {
	parsed, err := logrus.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("parse log level from string: %w", err)
	}

	l.Logger.SetLevel(parsed)
	return nil
}
