package logger

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	config "../../v1/config"
	"github.com/Sirupsen/logrus"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
)

var logger *logrus.Logger

func init() {
	stdOut := "file"
	logFile := "logs/app.log"

	var logSettings config.LogSettings
	logSettings.Level = "debug"

	logger = logrus.New()
	logger.Level, _ = logrus.ParseLevel(logSettings.Level)
	logger.Formatter = new(logrus.JSONFormatter)

	switch strings.ToLower(logSettings.StdOut) {
	case "console":
		stdOut = "console"
	}

	if len(logSettings.FileName) > 0 {
		logFile = logSettings.FileName
	}

	logger.Out = os.Stdout
	if stdOut == "file" {
		dir := filepath.Dir(logFile)
		if len(dir) > 0 {
			err := os.MkdirAll(dir, 777)
			if err != nil {
				logger.Info("Logger stdout is set to console.")
			}
		}
		writer, err := rotatelogs.New(
			logFile+".%Y%m%d%H%M", // rotation pattern
			rotatelogs.WithLinkName(logFile),
			rotatelogs.WithMaxAge(24*time.Hour),
			rotatelogs.WithRotationTime(time.Hour),
		)
		if err != nil {
			logger.Info("Logger stdout is set to console.")
		}
		logger.Hooks.Add(lfshook.NewHook(lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.PanicLevel: writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
		}))
	}
}
