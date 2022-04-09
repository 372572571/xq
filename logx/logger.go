package logx

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Logger struct {
		log *zap.SugaredLogger
	}
)

func NewLog(cfg Config) *Logger {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		TimeKey:       "ts",
		StacktraceKey: "stacktrace",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,

		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	var level = zapcore.InfoLevel
	switch cfg.Level {
	case "debug":
		level = zap.DebugLevel
	case "warn":
		level = zap.WarnLevel
	default:
		level = zap.InfoLevel
	}

	// 判断回调
	infoLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= level
	})

	errorLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == zapcore.ErrorLevel
	})

	panicLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zapcore.DPanicLevel
	})

	list_core := []zapcore.Core{
		// 控制台日志
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), infoLevel)}

	// 获取 info 日志文件
	if cfg.Info != "" {
		list_core = append(list_core,
			zapcore.NewCore(encoder, zapcore.AddSync(getWriter(cfg.Info)), infoLevel))
	}

	// 获取错误日志文件
	if cfg.Error != "" {
		list_core = append(list_core,
			zapcore.NewCore(encoder, zapcore.AddSync(getWriter(cfg.Error)), errorLevel))
	}

	if cfg.Panic != "" {
		list_core = append(list_core,
			zapcore.NewCore(encoder, zapcore.AddSync(getWriter(cfg.Panic)), panicLevel))
	}

	// 创建具体的Logger
	core := zapcore.NewTee(list_core...)

	// 显示调用打印日志的是哪一行的code行数
	// zap.AddCallerSkip(1)
	log := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(panicLevel),
	) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	logger := Logger{}
	logger.log = log.Sugar()
	return &logger
}

func getWriter(filename string) io.Writer {

	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s.log", filename),
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

}

func (l Logger) Desugar(args ...interface{}) {
	l.log.Desugar()
}

func (l Logger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l Logger) Debugf(template string, args ...interface{}) {
	l.log.Debugf(template, args...)
}
func (l Logger) Debugw(message string, args ...interface{}) {
	l.log.Debugw(message, args...)
}

func (l Logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l Logger) Infof(template string, args ...interface{}) {
	l.log.Infof(template, args...)
}

func (l Logger) Infow(message string, args ...interface{}) {
	l.log.Infow(message, args...)
}

func (l Logger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

func (l Logger) Warnf(template string, args ...interface{}) {
	l.log.Warnf(template, args...)
}

func (l Logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l Logger) Errorf(template string, args ...interface{}) {
	l.log.Errorf(template, args...)
}

func (l Logger) DPanic(args ...interface{}) {
	l.log.DPanic(args...)
}

func (l Logger) DPanicf(template string, args ...interface{}) {
	l.log.DPanicf(template, args...)
}

func (l Logger) Panic(args ...interface{}) {
	l.log.Panic(args...)
}

func (l Logger) Panicf(template string, args ...interface{}) {
	l.log.Panicf(template, args...)
}

func (l Logger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

func (l Logger) Fatalf(template string, args ...interface{}) {
	l.log.Fatalf(template, args...)
}
