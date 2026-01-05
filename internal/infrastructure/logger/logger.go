package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *zap.Logger
	sugarLogger  *zap.SugaredLogger
)

// Init инициализирует глобальный логгер
func Init(level string, development bool) error {
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "fatal":
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	var config zap.Config
	if development {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	config.Level = zap.NewAtomicLevelAt(zapLevel)
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	logger, err := config.Build()
	if err != nil {
		return err
	}

	globalLogger = logger
	sugarLogger = logger.Sugar()

	return nil
}

// Get возвращает логгер
func Get() *zap.Logger {
	if globalLogger == nil {
		// Фолбэк на логгер по умолчанию
		logger, _ := zap.NewProduction()
		return logger
	}
	return globalLogger
}

// Sugar возвращает SugaredLogger
func Sugar() *zap.SugaredLogger {
	if sugarLogger == nil {
		logger, _ := zap.NewProduction()
		return logger.Sugar()
	}
	return sugarLogger
}

// Sync синхронизирует логгер
func Sync() error {
	if globalLogger != nil {
		return globalLogger.Sync()
	}
	return nil
}

// WithContext создает логгер с контекстными полями
func WithContext(fields ...zap.Field) *zap.Logger {
	return Get().With(fields...)
}
