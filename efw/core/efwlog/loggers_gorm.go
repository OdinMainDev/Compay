package efwlog

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm/logger"
)

func logGorm(args ...any) {
	LogByLoggerName("gorm", 0, args...)
}

func logGormError(args ...any) {
	LogByLoggerName("gorm_error", 0, args...)
}

func logGormDebug(args ...any) {
	LogByLoggerName("gorm_debug", 0, args...)
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
		logGormError(FmtGormError(rows, elapsed, err.Error(), sql))
	} else {
		logGormDebug(FmtGormDebug(rows, elapsed, sql))
	}
}

func FmtGormError(rows int64, elapsed time.Duration, err, sql string) string {
	return fmt.Sprintf(
		"Rows affected: %d\n"+
			"Duration: %d\n"+
			"Error: %s\n"+
			"%s\n",
		rows, elapsed, err, sql)
}

func FmtGormDebug(rows int64, elapsed time.Duration, sql string) string {
	return fmt.Sprintf(
		"Rows affected: %d\n"+
			"Duration: %v\n"+
			"SQL: %s\n",
		rows, elapsed, sql)
}
