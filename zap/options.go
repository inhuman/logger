package zap

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapOpts struct {
	atomicLogLevel     zap.AtomicLevel
	encoderCfg         zapcore.EncoderConfig
	encoderConstructor func(cfg zapcore.EncoderConfig) zapcore.Encoder
	encoderTimeLayout  string
	logFile            *os.File
	zapOptions         []zap.Option
}

func defaultOptions() *zapOpts {
	conf := zap.NewDevelopmentEncoderConfig()
	conf.FunctionKey = "F"

	opts := &zapOpts{
		atomicLogLevel:     zap.NewAtomicLevel(),
		encoderCfg:         conf,
		encoderConstructor: defaultEncoderConstructor(),
		encoderTimeLayout:  time.RFC3339Nano,
		logFile:            os.Stdout,
		zapOptions: []zap.Option{
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
		},
	}

	return opts
}

func defaultEncoderConstructor() func(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return func(cfg zapcore.EncoderConfig) zapcore.Encoder {
		cfg.FunctionKey = "func"

		return &PrettyJSONEncoder{Encoder: zapcore.NewJSONEncoder(cfg)}
	}
}

type Option func(zo *zapOpts)

// WithProductionEncodingConfig returns an opinionated EncoderConfig for
// production environments.
func WithProductionEncodingConfig() Option {
	return func(zo *zapOpts) {
		zo.encoderCfg = zap.NewProductionEncoderConfig()
	}
}

// WithDevelopmentEncodingConfig returns an opinionated EncoderConfig for
// production environments.
func WithDevelopmentEncodingConfig() Option {
	return func(zo *zapOpts) {
		zo.encoderCfg = zap.NewDevelopmentEncoderConfig()
	}
}

// WithLogLevel alters the logging level
func WithLogLevel(level zapcore.Level) Option {
	return func(zo *zapOpts) {
		zo.atomicLogLevel.SetLevel(level)
	}
}

// WithTimeEncoderOfLayout alters the time field layout
func WithTimeEncoderOfLayout(layout string) Option {
	return func(zo *zapOpts) {
		zo.encoderTimeLayout = layout
	}
}

// WithLogFile alters the output (default os.Stdout)
func WithLogFile(file *os.File) Option {
	return func(zo *zapOpts) {
		zo.logFile = file
	}
}

// WithZapOptions sets go.uber.org/zap constructor options
// see https://github.com/uber-go/zap/blob/master/options.go
func WithZapOptions(zapOptions ...zap.Option) Option {
	return func(zo *zapOpts) {
		zo.zapOptions = zapOptions
	}
}

// WithEncoderConstructor set log encoder constructor
func WithEncoderConstructor(constructor func(cfg zapcore.EncoderConfig) zapcore.Encoder) Option {
	return func(zo *zapOpts) {
		zo.encoderConstructor = constructor
	}
}
