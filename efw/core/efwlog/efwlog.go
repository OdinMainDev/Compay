package efwlog

import (
	"fmt"
	"runtime"

	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

type EfwLog struct {
	FileName string
	Logger   *logging.Logger

	// Additional fields
	MaxSize    int
	MaxBackups int
	MaxAge     int

	EnableConsoleOutput bool
	DisableCallerInfo   bool
}

// FileName -> Logger
var allLoggers = []*EfwLog{}

var fileLogFormat = logging.MustStringFormatter("[%{time:02.01.2006 15:04:05.000}%{message}")

const (
	logsDirectory = "resources/logs"
	logs          = "resources/logs"
)

var ErrorLoggerNotExists = func(loggerName string) error {
	return fmt.Errorf("%s %s", "[EFW: LOGGER] NotExists:", loggerName)
}

func LogByLoggerName(name string, args ...any) error {
	for _, suffix := range []string{"", ".log"} {
		var handler *EfwLog
		var exists bool
		for _, v := range allLoggers {
			if v.FileName == name+suffix {
				exists = true
				handler = v
				break
			}
		}
		if !exists {
			continue
		}
		if handler.Logger == nil {
			continue
		}
		info := ""
		if !handler.DisableCallerInfo {
			info = " " + getCallerInfo(0)
		}
		message := fmt.Sprint(args...)
		handler.Logger.Info(logFormatter(info, message))
		return nil
	}

	return ErrorLoggerNotExists(name)
}

func getCallerInfo(callerSkip int, args ...string) string {
	_, file, line, ok := runtime.Caller(callerSkip + 2)
	if !ok {
		return ""
	}

	if len(args) == 1 {
		return fmt.Sprintf("[%s] %s:%d", args[0], file, line)
	}

	return fmt.Sprintf("%s:%d", file, line)
}

func createFileLogger(efwLog *EfwLog, name string) logging.Backend {
	name = getLogFile(name)
	return logging.NewLogBackend(&lumberjack.Logger{
		Filename:   name,
		MaxSize:    efwLog.MaxSize,    // megabytes
		MaxBackups: efwLog.MaxBackups, // count of backups files
		MaxAge:     efwLog.MaxAge,     // days lifetime of backups files (0 - unlimited)
	}, "", 0)
}

func getLogFile(fileName string) string {
	return fileName
}

func logFormatter(info, message string) any {
	return fmt.Sprintf("%s] %s", info, message)
}
