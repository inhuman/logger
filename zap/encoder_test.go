package zap

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestEncoder(t *testing.T) {
	logger, _ := New(WithEncoderConstructor(func(cfg zapcore.EncoderConfig) zapcore.Encoder {
		return &PrettyJSONEncoder{Encoder: zapcore.NewJSONEncoder(cfg)}
	}))

	testFunc1(logger)
}

func testFunc1(logger *zap.Logger) {
	testFunc2(logger)
}

func testFunc2(logger *zap.Logger) {
	testFunc3(logger)
}

func testFunc3(logger *zap.Logger) {
	testFunc4(logger)
}

func testFunc4(logger *zap.Logger) {
	logger.Error("test err")
}
