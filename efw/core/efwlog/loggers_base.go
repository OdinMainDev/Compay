package efwlog

import (
	"log"
)

func Panic(args ...any) {
	LogByLoggerName("panic", 1, args...)
	log.Panic(args...)
}

func Log(args ...any) {
	LogByLoggerName("info", 1, args...)
}

func LogError(args ...any) {
	LogByLoggerName("error", 1, args...)
}

func LogWarning(args ...any) {
	LogByLoggerName("warning", 1, args...)
}

func LogDebug(args ...any) {
	LogByLoggerName("debug", 1, args...)
}
