package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/blacksheepaul/templateToGo/core/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	instance Logger
	once     sync.Once
)

type Logger interface {
	Debug(fields ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Info(fields ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warn(fields ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Error(fields ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatal(fields ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
}

var ilogger Logger

func GetLogger() Logger {
	return ilogger
}

func SetZapLogger(cfg config.Config) Logger {
	// Setting log level
	level, err := zap.ParseAtomicLevel(cfg.Log.Level)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse log level: %s", err))
	}

	// Setting log output
	logDir := filepath.Dir(cfg.Log.Path)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0744)
		if err != nil {
			panic(fmt.Sprintf("Failed to create log directory: %s", err))
		}
	}

	// Setting log format
	encoderConfig := zapcore.EncoderConfig{
		LevelKey: "lvl", TimeKey: "ts",
		MessageKey: "msg", CallerKey: "src",
		NameKey: "logger", StacktraceKey: "stack", // not used for now
		//
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Setting log rotation
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Log.Path,
		MaxSize:    cfg.Log.Rotation.MaxSize, //MB
		MaxBackups: cfg.Log.Rotation.MaxBackups,
		MaxAge:     cfg.Log.Rotation.MaxAge, // days
	})

	c := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		level,
	)

	logger := zap.New(c,
		zap.AddCaller(),
	).
		Sugar()
	defer logger.Sync()

	ilogger = logger

	return ilogger
}
