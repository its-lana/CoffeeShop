package logger

var Log Logger

type Logger interface {
	Errorf(format string, args ...any)
	Info(args ...any)
	WithField(key string, value any) Logger
	WithFields(map[string]any) Logger
}

func SetLogger(log Logger) {
	Log = log
}
