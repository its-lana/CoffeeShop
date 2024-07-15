package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type loggerWrapper struct {
	lw *logrus.Logger
}

func (logger *loggerWrapper) Errorf(format string, args ...any) {
	logger.lw.Errorf(format, args...)
}

func (logger *loggerWrapper) Info(args ...any) {
	logger.lw.Info(args...)
}
func (logger *loggerWrapper) WithField(key string, value any) (entry Logger) {
	entry = &LoggerEntry{logger.lw.WithField(key, value)}
	return
}
func (logger *loggerWrapper) WithFields(args map[string]any) (entry Logger) {
	entry = &LoggerEntry{logger.lw.WithFields(args)}
	return
}

func NewLogger() *loggerWrapper {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
		ForceColors:     true,
	})
	log.SetOutput(os.Stdout)

	return &loggerWrapper{log}
}
