package logger

import (
	"os"

	logrus "github.com/m1t0k/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout
}

func Debug(message string) {
	logger.Debug(message)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

func Info(message string) {
	logger.Info(message)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args)
}

func Warn(message string) {
	logger.Warn(message)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args)
}

func Error(message string) {
	logger.Error(message)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args)
}

func Fatal(message string) {
	logger.Fatal(message)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args)
}

func Panic(message string) {
	logger.Panic(message)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args)
}
