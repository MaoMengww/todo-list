package logger

import (
	"context"
	"os"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

func InitLogger(){
	
	//配置Encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 日志等级大写
		EncodeTime:     zapcore.ISO8601TimeEncoder,  // 时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 定义日志级别优先级
	// infoLevel: 记录 Info 及以上
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.InfoLevel
	})
	// errorLevel: 只记录 Error 及以上
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.ErrorLevel
	})

	//多核心，不同地方输出日志
	cores := [...]zapcore.Core{
		zapcore.NewCore(encoder, os.Stdout, errorLevel),    //写入终端
		zapcore.NewCore(
			encoder,
			getWriteSyncer("./logs/info.log"),
			infoLevel,
		),	
		zapcore.NewCore(
			encoder,
			getWriteSyncer("./logs/error.log"),
			errorLevel,
		),												//写入文件
	}


	//合并核心
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func Sync() {
	if logger != nil {
		_ = logger.Sync()
	}
}


//创建日志轮转器，写入文件/进行轮转
func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    100,  // MB
		MaxBackups: 30,   // 最大备份数
		MaxAge:     30,   // 保存天数
		Compress:   true, // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		return logger.With("trace_id", span.SpanContext().TraceID().String())
	}
	return logger
}




//封装日志方法
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}


func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}	



