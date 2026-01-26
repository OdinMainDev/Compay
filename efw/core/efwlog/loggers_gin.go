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

type EfwLogWriter struct{}

func NewEfwLogWriter() io.Writer {
	return &EfwLogWriter{}
}

func (w *EfwLogWriter) Write(p []byte) (n int, err error) {
	logGin(string(p))
	return len(p), nil
}
