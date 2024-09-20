# logger

### Go библиотека logger

По сути это обертка над пакетом go.uber.org/zap, но так же есть конвертер в logr.Logr интерфейс

### Поддерживаемые уровни логирование (log level)

**DEBUG** (debug) - логи как правило, объемные и обычно отключаются в проде

**INFO** (info) - лог левел по умолчанию

**WARN** (warn) - более важен, чем Info, но не нуждаются в отдельном человеческом внимании

**ERROR** (error) - имеют высокий приоритет, на такие ошибки нужно обращать внимание

**DPANIC** (dpanic) - особенно важные ошибки, дев огружении логгер паникует после вывода сообщения

**PANIC** (panic) - критичные ошибки, после вывода сообщения логгер бросает панику

**FATAL** (fatal) - фатальные ошибки, после вывода сообщения вызывается os.Exit(1)

### Использование

Логгер создается с atomic log level, т.е. лог левел может быть изменен в любой момент, это потокобезопасно

Значения по умолчанию:
```go
func defaultOptions() *zapOpts {
    opts := &zapOpts{
        atomicLogLevel:     zap.NewAtomicLevel(), // default INFO
        encoderCfg:         zap.NewDevelopmentEncoderConfig(),
        encoderConstructor: zapcore.NewJSONEncoder,
        encoderTimeLayout:  time.RFC3339Nano,
        logFile:            os.Stdout,
        zapOptions: []zap.Option{
            zap.AddCaller(),
            zap.AddStacktrace(zapcore.ErrorLevel),
        },
    }
        
    return opts
}
```

#### [Без опций](https://github.com/inhuman/logger/-/tree/master/examples/zap/base) 

```go
package main

import (
	loggerZap "github.com/inhuman/logger/zap"
	"go.uber.org/zap"
)

func main() {
	zapLogger, atomicLevel := loggerZap.New()

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
// {"L":"INFO","T":"2023-05-16T11:48:10.416+0300","C":"base/main.go:11","M":"Информационное сообщение","дополнительное поле":"значение"}
// {"L":"DEBUG","T":"2023-05-16T11:48:10.416+0300","C":"base/main.go:15","M":"Отладочное сообщение, будет отображено, т.к. поменяли лог левел на DEBUG"}
// {"L":"INFO","T":"2023-05-16T11:48:10.416+0300","C":"base/main.go:21","M":"Информационное сообщение","константное поле":"значение"}
```

### [С опциями](https://github.com/inhuman/logger/-/tree/master/examples/zap/prod_encoding)
```go
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
```

Список всех опций можно посмотреть в файле [options.go](https://github.com/inhuman/logger/-/tree/master/zap/options.go)