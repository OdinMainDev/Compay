package efwlog

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm/logger"
)

func logGorm(args ...any) {
	LogByLoggerName("gorm", args...)
}

func logGormError(args ...any) {
	LogByLoggerName("gorm_error", args...)
}

func logGormDebug(args ...any) {
	LogByLoggerName("gorm_debug", args...)
}

type EsGormLogger struct{}

func (gl *EsGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return gl
}

func (gl *EsGormLogger) Info(_ context.Context, msg string, data ...any) {
	logGorm(msg, data)
}

func (gl *EsGormLogger) Warn(_ context.Context, msg string, data ...any) {
	logGormError(msg, data)
}

func (gl *EsGormLogger) Error(_ context.Context, msg string, data ...any) {
	logGormError(msg, data)
}

func (gl *EsGormLogger) Trace(_ context.Context, start time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(start)
	sql, rows := fc()
	if err != nil {
		message := fmt.Sprintf(
			"######################## HAS ERROR ########################\n"+
				"Rows affected: %d\n"+
				"Duration: %d\n"+
				"Error: %s\n"+
				"%s\n"+
				"######################## END ERROR ########################\n",
			rows, elapsed, err.Error(), sql)
		logGormDebug(message)
	} else {
		message := fmt.Sprintf(
			"Rows affected: %d\n"+
				"Duration: %v\n"+
				"SQL: %s\n",
			rows, elapsed, sql)
		logGormDebug(message)
	}
}
