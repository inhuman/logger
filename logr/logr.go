package logr

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

type Logger interface {
	GetSink() logr.LogSink
	WithSink(sink logr.LogSink) logr.Logger
	Enabled() bool
	Info(msg string, keysAndValues ...interface{})
	Error(err error, msg string, keysAndValues ...interface{})
	V(level int) logr.Logger
	WithValues(keysAndValues ...interface{}) logr.Logger
	WithName(name string) logr.Logger
	WithCallDepth(depth int) logr.Logger
	WithCallStackHelper() (func(), logr.Logger)
	IsZero() bool
}

func FromZap(logger *zap.Logger) logr.Logger {
	return zapr.NewLoggerWithOptions(logger)
}
