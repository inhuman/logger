package main

import (
	toolsZap "github.com/inhuman/logger/zap"
	"go.uber.org/zap"
)

func main() {
	zapLogger, atomicLevel := toolsZap.New()

	zapLogger.Info("Информационное сообщение", zap.String("дополнительное поле", "значение"))
	zapLogger.Debug("Отладочное сообщение, не будет отображено, т.к. лог левел по умолчанию INFO")

	atomicLevel.SetLevel(zap.DebugLevel)
	zapLogger.Debug("Отладочное сообщение, будет отображено, т.к. поменяли лог левел на DEBUG")
	atomicLevel.SetLevel(zap.InfoLevel)

	loggerWithConstantField := zapLogger.With(zap.String("константное поле", "значение"))

	// теперь этот логгер будет всегда с константным полем
	loggerWithConstantField.Info("Информационное сообщение")
}

// Вывод:
// {"L":"INFO","T":"2023-05-16T12:31:50.771172206+03:00","C":"base/main.go:11","M":"Информационное сообщение","дополнительное поле":"значение"}
// {"L":"DEBUG","T":"2023-05-16T12:31:50.771215832+03:00","C":"base/main.go:15","M":"Отладочное сообщение, будет отображено, т.к. поменяли лог левел на DEBUG"}
// {"L":"INFO","T":"2023-05-16T12:31:50.771224682+03:00","C":"base/main.go:21","M":"Информационное сообщение","константное поле":"значение"}
