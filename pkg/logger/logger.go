package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func SetuoLogger(){
	//设置日志传入级别
	level := zap.NewAtomicLevel()
	level.SetLevel(zap.DebugLevel)

	//编码器
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:           "message",          				//消息
		LevelKey:             "level",							//级别
		TimeKey: 			  "time",							//时间
		NameKey:              "logger",							//日志器名称
		CallerKey: 			  "caller",							//调用者名称
		StacktraceKey:  	  "stacktrace",						//堆栈跟踪
		LineEnding:  		  zapcore.DefaultLineEnding,		//每行结束符号
		EncodeLevel:    	  zapcore.CapitalColorLevelEncoder,	//级别编码器
		EncodeTime: 		  CustomTimeEncoder,				//时间编码器
		EncodeDuration:       zapcore.StringDurationEncoder,	//时间段编码器
		EncodeCaller:         zapcore.FullCallerEncoder,		//调用者编码器
		ConsoleSeparator:     "",								//控制台分割符号
	})

	//多核心，不同地方输出日志
	cores := [...]zapcore.Core{
		zapcore.NewCore(encoder, os.Stdout, level),    //写入终端
		zapcore.NewCore(
			encoder,
			getwritesync(),
			level,
		),												//写入文件
	}


	//合并核心
	Logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller()).Sugar()
	defer func(Logger *zap.SugaredLogger){
		_ = Logger.Sync()
	}(Logger)
}


//创建日志轮转器，写入文件/进行轮转
func getwritesync() zapcore.WriteSyncer{
	lumberjackLogger := &lumberjack.Logger{
		Filename: 	global.Config.ZapConfig.Filename,   //文件名
		MaxSize:    global.Config.ZapConfig.MaxSize,    //最大文件MB大小   结构体显式访问（还有匿名访问）
		MaxAge:     global.Config.ZapConfig.MaxAge, 	//最大文件保存天数
		MaxBackups: global.Config.ZapConfig.MaxBackups,	//最大备份文件
		LocalTime:  true,								//本地时间
	}
	return zapcore.AddSync(lumberjackLogger)
}

//时间编码器
func CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder){
	encoder.AppendString(t.Format("2006-01-01 15:04:05"))
}