package qlog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

var (
	sugaredLogger *zap.SugaredLogger
	core          zapcore.Core
)

func logger() *zap.SugaredLogger {
	if sugaredLogger == nil {
		panic("日志未初始化")
	}
	return sugaredLogger
}

func Setup(log *Conf) error {
	encoder := makeEncoder()

	var cores []zapcore.Core
	for _, logCore := range log.LogCores {
		w := makeWriteSyncer(logCore.LumberJack)
		cores = append(cores, zapcore.NewCore(encoder, w, logCore.Levels))
	}
	if log.Debug {
		w := os.Stdout
		cores = append(cores, zapcore.NewCore(encoder, w, Levels{
			zapcore.DebugLevel,
			zapcore.InfoLevel,
			zapcore.WarnLevel,
			zapcore.ErrorLevel,
			zapcore.DPanicLevel,
			zapcore.PanicLevel,
			zapcore.FatalLevel,
		}))
	}

	if len(cores) == 0 {
		core = zapcore.NewNopCore()
	} else {
		core = zapcore.NewTee(cores...)
	}

	sugaredLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
	return nil
}

func makeWriteSyncer(jack LumberJack) zapcore.WriteSyncer {
	lumberJackLogger := lumberjack.Logger{
		Filename:   filepath.Join(jack.Path, jack.Name),
		MaxSize:    jack.MaxSizeInMB,
		MaxBackups: jack.MaxBackups,
		MaxAge:     jack.MaxAgeInDay,
		Compress:   jack.Compress,
	}
	return zapcore.AddSync(&lumberJackLogger)
}

func makeEncoder() zapcore.Encoder {
	// 编码器配置
	c := zap.NewProductionEncoderConfig()
	// 指定时间编码器
	c.EncodeTime = zapcore.RFC3339TimeEncoder
	// 日志级别用大写
	c.EncodeLevel = zapcore.CapitalLevelEncoder
	// 编码器
	return zapcore.NewConsoleEncoder(c)
}

func Debug(args ...any) {
	logger().Debug(args...)
}

func Debugf(template string, args ...any) {
	logger().Debugf(template, args...)
}

func Info(args ...any) {
	logger().Info(args...)
}

func Infof(template string, args ...any) {
	logger().Infof(template, args...)
}

func Warn(args ...any) {
	logger().Warn(args...)
}

func Warnf(template string, args ...any) {
	logger().Warnf(template, args...)
}

func Error(args ...any) {
	logger().Error(args...)
}

func Errorf(template string, args ...any) {
	logger().Errorf(template, args...)
}
