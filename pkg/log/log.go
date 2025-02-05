package log

import (
	"easy-gin-template/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

func InitLog() *zap.Logger {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	level := getLevel(config.CFG.Log.Level)
	core := zapcore.NewCore(encoder, writerSyncer, level)

	return zap.New(core, zap.AddCaller())
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.CFG.Log.FilePath,
		MaxSize:    config.CFG.Log.MaxSize,
		MaxBackups: config.CFG.Log.MaxBackups,
		MaxAge:     config.CFG.Log.MaxAge,
		Compress:   config.CFG.Log.Compress,
	}

	ws := io.MultiWriter(lumberJackLogger, os.Stdout)
	return zapcore.AddSync(ws)
}

func getEncoder() zapcore.Encoder {
	customEncodeTime := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(time.DateTime))
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     customEncodeTime,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}
