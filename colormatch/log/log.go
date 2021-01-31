package log

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	clogger *zap.Logger
	logger  *zap.Logger
)

// ColoredLogger log with color support
func ColoredLogger() *zap.Logger {
	if clogger != nil {
		clogger.Debug("using global logger")
		return clogger
	}
	// https://github.com/uber-go/zap/pull/307#issuecomment-504794011
	// conf := zap.NewProductionEncoderConfig()
	conf := zap.NewDevelopmentEncoderConfig()
	conf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	clogger = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(conf),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	))
	return clogger
}

// Logger logger with colored support disabled
func Logger() *zap.Logger {
	if logger != nil {
		logger.Debug("using global logger")
		return logger
	}
	logger, _ := zap.NewProduction()
	return logger
}
