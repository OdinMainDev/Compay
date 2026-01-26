package efwlog

import (
	"io"
)

func logGin(args ...any) {
	LogByLoggerName("gin", 0, args...)
}

func LogGinError(args ...any) {
	LogByLoggerName("gin_error", 3, args...)
}

type EfwGinLogWriter struct{}

func NewEfwGinLogWriter() io.Writer {
	return &EfwGinLogWriter{}
}

func (w *EfwGinLogWriter) Write(p []byte) (n int, err error) {
	logGin(string(p))
	return len(p), nil
}
