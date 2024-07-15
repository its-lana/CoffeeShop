package logger

import "github.com/sirupsen/logrus"

type LoggerEntry struct {
	entry *logrus.Entry
}

func (logger *LoggerEntry) Errorf(format string, args ...any) {
	logger.entry.Errorf(format, args...)
}
func (logger *LoggerEntry) Info(args ...any) {
	logger.entry.Info(args...)
}
func (logger *LoggerEntry) WithField(key string, value any) (entry Logger) {
	entry = &LoggerEntry{logger.entry.WithField(key, value)}
	return
}
func (logger *LoggerEntry) WithFields(args map[string]any) (entry Logger) {
	entry = &LoggerEntry{logger.entry.WithFields(args)}
	return
}
