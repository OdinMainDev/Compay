package efwlog

import (
	"io"
)

func logGin(args ...any) {
	LogByLoggerName("gin", args...)
}

type EsGinWriter struct{}

func NewEsGinWriter() io.Writer {
	return &EsGinWriter{}
}

func (w *EsGinWriter) Write(p []byte) (n int, err error) {
	logGin(string(p)) // Используйте ваш метод Log или другой, как вам нужно
	return len(p), nil
}
