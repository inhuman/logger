package main

import (
	loggerZap "github.com/inhuman/logger/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	zapLogger, atomicLevel := loggerZap.New(
		loggerZap.WithProductionEncodingConfig(),
		loggerZap.WithLogLevel(zapcore.WarnLevel),
	)

	zapLogger.Info("Информационное сообщение, не будет отображено, т.к. лог левел WARN",
		zap.String("дополнительное инфо поле", "значение"))

	zapLogger.Debug("Отладочное сообщение, не будет отображено, т.к. лог левел WARN",
		zap.String("дополнительное дебаг поле", "значение"))

	zapLogger.Debug("Предупреждающее сообщение, будет отображено, т.к. лог левел WARN",
		zap.String("дополнительное варн поле", "значение"))

	atomicLevel.SetLevel(zap.DebugLevel)
	zapLogger.Debug("Отладочное сообщение, будет отображено, т.к. поменяли лог левел на DEBUG")
	atomicLevel.SetLevel(zap.WarnLevel)

	loggerWithConstantField := zapLogger.With(zap.String("константное поле", "значение"))

	// теперь этот логгер будет всегда с константным полем
	loggerWithConstantField.Warn("Предупреждающее сообщение")
}

// Вывод:
// {"level":"debug","ts":"2023-05-16T12:37:01.885671972+03:00","caller":"prod_encoding/main.go:25","msg":"Отладочное сообщение, будет отображено, т.к. поменяли лог левел на DEBUG"}
// {"level":"warn","ts":"2023-05-16T12:37:01.885714988+03:00","caller":"prod_encoding/main.go:31","msg":"Предупреждающее сообщение","константное поле":"значение"}
