package logger

import (
	log "github.com/m1t0k/logrus"
)

type Logger struct {
}

func Debug(message string) {
	log.Debug(message)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args)
}

func Info(message string) {
	log.Info(message)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args)
}

func Warn(message string) {
	log.Warn(message)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args)
}

func Error(message string) {
	log.Error(message)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args)
}

func Fatal(message string) {
	log.Fatal(message)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args)
}

func Panic(message string) {
	log.Panic(message)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args)
}
