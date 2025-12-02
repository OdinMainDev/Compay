package efwlog

import (
	"fmt"
	"strings"

	"github.com/op/go-logging"
)

var ErrorLoggerBuildFail = func(loggerName string) error {
	return fmt.Errorf("%s %s", "[EFW: LOGGER BUILDER] BadLoggerName:", loggerName)
}

func NewEfwLog(name string) *EfwLog {
	fileName := name
	if !strings.HasSuffix(fileName, ".log") {
		fileName += ".log"
	}

	return &EfwLog{
		FileName: fileName,
		Logger:   logging.MustGetLogger(name),

		MaxSize:    100,
		MaxBackups: 100,
		MaxAge:     365,

		EnableConsoleOutput: false,
	}
}

func (e *EfwLog) SetMaxSize(maxSize int) *EfwLog {
	if maxSize <= 0 {
		maxSize = 100
		fmt.Println("MaxSize is less than 0, set to 100 mb")
	}

	if maxSize > 10000 {
		maxSize = 10000
		fmt.Println("MaxSize is more than 10000 mb, set to 10000 mb")
	}
	e.MaxSize = maxSize
	return e
}

func (e *EfwLog) SetMaxBackups(maxBackups int) *EfwLog {
	if maxBackups <= 0 {
		maxBackups = 100
		fmt.Println("MaxBackups is less than 0, set to 100 backups")
	}

	if maxBackups > 10000 {
		maxBackups = 10000
		fmt.Println("MaxBackups is more than 10000 backups, set to 10000 backups")
	}
	e.MaxBackups = maxBackups
	return e
}

func (e *EfwLog) SetMaxAge(maxAge int) *EfwLog {
	if maxAge <= 0 {
		maxAge = 365
		fmt.Println("MaxAge is less than 0, set to 365 days")
	}

	if maxAge > 10000 {
		maxAge = 10000
		fmt.Println("MaxAge is more than 10000 days, set to 10000 days")
	}
	e.MaxAge = maxAge
	return e
}

func (e *EfwLog) SetEnableConsoleOutput(enable bool) *EfwLog {
	e.EnableConsoleOutput = enable
	return e
}

func (e *EfwLog) SaveLogger() error {
	for _, logger := range allLoggers {
		if logger.FileName == e.FileName {
			return ErrorLoggerBuildFail(e.FileName)
		}
	}

	allLoggers = append(allLoggers, e)
	return nil
}
