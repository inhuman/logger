package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(opts ...Option) (*zap.Logger, zap.AtomicLevel) {
	options := defaultOptions()

	for _, opt := range opts {
		opt(options)
	}

	options.encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(options.encoderTimeLayout)

	logger := zap.New(
		zapcore.NewCore(
			options.encoderConstructor(options.encoderCfg),
			zapcore.Lock(options.logFile),
			options.atomicLogLevel,
		),
		options.zapOptions...,
	)

	return logger, options.atomicLogLevel
}
