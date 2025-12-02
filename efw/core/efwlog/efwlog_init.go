package efwlog

import (
	"os"
	"path/filepath"

	"github.com/op/go-logging"
)

func Init() error {
	for _, v := range allLoggers {
		initFileLogger(v)
	}
	return nil
}

func initFileLogger(efwLog *EfwLog) {
	customLogsDir := os.Getenv(logsDirectory)
	var backend logging.Backend
	if customLogsDir == "" {
		backend = createFileLogger(efwLog, filepath.Join(logs, efwLog.FileName))
	} else {
		backend = createFileLogger(efwLog, filepath.Join(customLogsDir, efwLog.FileName))
	}
	fileFormatter := logging.NewBackendFormatter(backend, fileLogFormat)
	fileLoggerLeveled := logging.AddModuleLevel(fileFormatter)
	fileLoggerLeveled.SetLevel(logging.DEBUG, "")

	if efwLog.EnableConsoleOutput {
		outBackend := logging.NewLogBackend(os.Stdout, "", 0)
		outFormatter := logging.NewBackendFormatter(outBackend, fileLogFormat)
		outLoggerLeveled := logging.AddModuleLevel(outFormatter)
		outLoggerLeveled.SetLevel(logging.DEBUG, "")
		multiBackend := logging.MultiLogger(outLoggerLeveled, fileLoggerLeveled)

		efwLog.Logger.SetBackend(multiBackend)
	} else {
		efwLog.Logger.SetBackend(fileLoggerLeveled)
	}
}
