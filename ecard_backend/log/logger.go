package log

import (
	"github.com/siddontang/go/log"
)

var Logger *log.Logger

func Init(logPath string, loglevel string, filesize int, filenum int) {
	level := log.LevelInfo
	if loglevel == "Debug" {
		level = log.LevelDebug
	} else if loglevel == "Warn" {
		level = log.LevelWarn
	} else if loglevel == "Error" {
		level = log.LevelError
	}
	maxBytes := filesize * 1024 * 1024

	handler, _ := log.NewRotatingFileHandler(logPath, maxBytes, filenum)
	Logger = log.NewDefault(handler)

	Logger.SetLevel(level)
}

func Debug(format string, v ...interface{}) {
	Logger.Output(2, log.LevelDebug, format, v...)
}

func Info(format string, v ...interface{}) {
	Logger.Output(2, log.LevelInfo, format, v...)
}

func Warn(format string, v ...interface{}) {
	Logger.Output(2, log.LevelWarn, format, v...)
}

func Error(format string, v ...interface{}) {
	Logger.Output(2, log.LevelError, format, v...)
}
