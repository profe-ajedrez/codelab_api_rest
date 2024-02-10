package logadapter

import (
	"context"
	"log"
	"os"

	sqldblogger "github.com/simukti/sqldb-logger"
)

var Log Logger = Logger{log.New(os.Stdout, "", log.LstdFlags)}

type Logger struct {
	*log.Logger
}

func (l *Logger) Debug(args ...interface{}) {
	l.Logger.Println(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.Logger.Println(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.Logger.Println(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Logger.Println(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.Logger.Println(args...)
	os.Exit(1)
}

func (l *Logger) Panic(args ...interface{}) {
	panic(args)
}

func (l *Logger) Log(_ context.Context, level sqldblogger.Level, msg string, data map[string]interface{}) {

	switch level {
	case sqldblogger.LevelError:
		l.Error(msg)
	case sqldblogger.LevelInfo:
		l.Info(msg)
	case sqldblogger.LevelDebug:
		l.Debug(msg)
	case sqldblogger.LevelTrace:
		l.Trace(msg)
	default:
		l.Error(msg)
	}

}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

func (l *Logger) Trace(format string, args ...interface{}) {
	l.Logger.Println(args...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
	os.Exit(1)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	panic(args)
}

func NewSimpleLogger() *Logger {
	return &Logger{log.New(os.Stdout, "", log.LstdFlags)}
}

type SimpleLoggerAdapter struct {
	Logger *Logger
}

func NewSimpleAdapter(logger *SimpleLoggerAdapter) sqldblogger.Logger {
	return NewSimpleLogger()
}
