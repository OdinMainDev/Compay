package efwlog

import (
	"log"
)

func Panic(args ...any) {
	LogByLoggerName("panic", args...)
	log.Panic(args...)
}

func Log(args ...any) {
	LogByLoggerName("info", args...)
}

func LogError(args ...any) {
	LogByLoggerName("error", args...)
}

func LogWarning(args ...any) {
	LogByLoggerName("warning", args...)
}

func LogDebug(args ...any) {
	LogByLoggerName("debug", args...)
}
