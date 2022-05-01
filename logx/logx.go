package logx

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

var l *zap.Logger

func init() {
	log := NewLog(Config{
		Enable: true,
		Level:  "debug", // 输出等级  debug info warn
	})
	l = log.log.Desugar()
}

func SetOut(log *Logger) {
	l = log.log.Desugar()
}

func Debug(v ...interface{}) {
	l.Sugar().Debug(v)
}

func Debugf(format string, v ...interface{}) {
	l.Sugar().Debugf(format, v)
}

func Info(v ...interface{}) {
	l.Sugar().Info(v)
}

func Infof(format string, v ...interface{}) {
	l.Sugar().Infof(format, v)
}

func Warn(v ...interface{}) {
	l.Sugar().Warn(v)
}

func Warnf(format string, v ...interface{}) {
	l.Sugar().Warnf(format, v)
}

func Error(v ...interface{}) {
	l.Sugar().Error(v)
}

func Errorf(format string, v ...interface{}) {
	l.Sugar().Errorf(format, v)
}

func DPanic(v ...interface{}) {
	l.Sugar().DPanic(v)
}

func DPanicf(format string, v ...interface{}) {
	l.Sugar().DPanicf(format, v)
}

func Printf(format string, v ...interface{}) {
	l.Sugar().Infof(format, v)
}

func Print(v ...interface{}) {
	l.Sugar().Info(v...)
}

func Println(v ...interface{}) {
	l.Sugar().Info(v)
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	l.Sugar().Error(v)
	os.Exit(1)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	l.Sugar().Errorf(format, v...)
	os.Exit(1)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	l.Sugar().Error(v...)
	os.Exit(1)
}

// Panic is equivalent to l.Print() followed by a call to panic().
func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Sugar().Panic(s)
	panic(s)
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Sugar().Panicf(s)
	panic(s)
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.Sugar().Panic(s)
	panic(s)
}
