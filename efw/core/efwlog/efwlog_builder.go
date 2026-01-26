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

func (that *EfwLog) SetMaxSize(maxSize int) *EfwLog {
	if maxSize <= 0 {
		maxSize = 100
		fmt.Println("MaxSize is less than 0, set to 100 mb")
	}

	if maxSize > 10000 {
		maxSize = 10000
		fmt.Println("MaxSize is more than 10000 mb, set to 10000 mb")
	}
	that.MaxSize = maxSize
	return that
}

func (that *EfwLog) SetMaxBackups(maxBackups int) *EfwLog {
	if maxBackups <= 0 {
		maxBackups = 100
		fmt.Println("MaxBackups is less than 0, set to 100 backups")
	}

	if maxBackups > 10000 {
		maxBackups = 10000
		fmt.Println("MaxBackups is more than 10000 backups, set to 10000 backups")
	}
	that.MaxBackups = maxBackups
	return that
}

func (that *EfwLog) SetMaxAge(maxAge int) *EfwLog {
	if maxAge <= 0 {
		maxAge = 365
		fmt.Println("MaxAge is less than 0, set to 365 days")
	}

	if maxAge > 10000 {
		maxAge = 10000
		fmt.Println("MaxAge is more than 10000 days, set to 10000 days")
	}
	that.MaxAge = maxAge
	return that
}

func (that *EfwLog) SetEnableConsoleOutput(enable bool) *EfwLog {
	that.EnableConsoleOutput = enable
	return that
}

func (that *EfwLog) SetDisableCallerInfo(disable bool) *EfwLog {
	that.DisableCallerInfo = disable
	return that
}

func (that *EfwLog) SaveLogger() {
	for i, logger := range allLoggers {
		if logger.FileName == that.FileName {
			allLoggers[i] = that
			return
		}
	}

	allLoggers = append(allLoggers, that)
}
