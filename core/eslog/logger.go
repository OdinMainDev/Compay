package eslog

import (
	"compay/core/utils"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm/logger"
)

const (
	logsDirectory              = "storage/logs"
	logs                       = "storage/logs"
	errorLogFileName           = "error.log"
	debugLogFileName           = "debug.log"
	infoLogFileName            = "info.log"
	chatLogFileName            = "chat.log"
	ginLogFileName             = "gin.log"
	gormLogFileName            = "gorm.log"
	gormErrorLogFileName       = "gorm_error.log"
	gormDebugLogFileName       = "gorm_debug.log"
	requestsLogFileName        = "requests.log"
	responsesLogFileName       = "responses.log"
	esRequestsLogFileName      = "es_requests.log"
	esProxyRequestsLogFileName = "es_proxy_requests.log"
)

var errorLogger = logging.MustGetLogger("error")
var debugLogger = logging.MustGetLogger("debug")
var infoLogger = logging.MustGetLogger("info")
var chatLogger = logging.MustGetLogger("chat")
var ginLogger = logging.MustGetLogger("gin")
var gormLogger = logging.MustGetLogger("gorm")
var gormErrorLogger = logging.MustGetLogger("gorm_error")
var gormDebugLogger = logging.MustGetLogger("gorm_debug")
var requestsLogger = logging.MustGetLogger("requests")
var responsesLogger = logging.MustGetLogger("responses")
var esRequestsLogger = logging.MustGetLogger("esRequests")
var esProxyRequestsLogger = logging.MustGetLogger("esProxyRequests")

var fileLogFormat = logging.MustStringFormatter("[%{time:15:04:05.000} %{message}")

func Initialize() {
	os.Setenv("TZ", "Europe/Moscow")
	initFileLogger(errorLogFileName, errorLogger)
	initFileLogger(debugLogFileName, debugLogger)
	initFileLogger(infoLogFileName, infoLogger)
	initFileLogger(chatLogFileName, chatLogger)
	initFileLogger(ginLogFileName, ginLogger)
	initFileLogger(gormLogFileName, gormLogger)
	initFileLogger(gormErrorLogFileName, gormErrorLogger)
	initFileLogger(gormDebugLogFileName, gormDebugLogger)
	initFileLogger(requestsLogFileName, requestsLogger)
	initFileLogger(responsesLogFileName, responsesLogger)
	initFileLogger(esRequestsLogFileName, esRequestsLogger)
	initFileLogger(esProxyRequestsLogFileName, esProxyRequestsLogger)
}

func initFileLogger(logFileName string, fileLogger *logging.Logger) {
	customLogsDir := os.Getenv(logsDirectory)
	var backend logging.Backend
	if customLogsDir == "" {
		backend = createFileLogger(filepath.Join(logs, logFileName), 1)
	} else {
		backend = createFileLogger(filepath.Join(customLogsDir, logFileName), 1)
	}
	fileFormatter := logging.NewBackendFormatter(backend, fileLogFormat)
	fileLoggerLeveled := logging.AddModuleLevel(fileFormatter)
	fileLoggerLeveled.SetLevel(logging.DEBUG, "")

	if os.Getenv("APP_DEBUG") == "true" &&
		!utils.InArray(
			logFileName,
			[]string{
				esRequestsLogFileName,
				esProxyRequestsLogFileName,
				requestsLogFileName,
				responsesLogFileName,
			},
		) {
		outBackend := logging.NewLogBackend(os.Stdout, "", 0)
		outFormatter := logging.NewBackendFormatter(outBackend, fileLogFormat)
		outLoggerLeveled := logging.AddModuleLevel(outFormatter)
		outLoggerLeveled.SetLevel(logging.DEBUG, "")
		multiBackend := logging.MultiLogger(outLoggerLeveled, fileLoggerLeveled)

		fileLogger.SetBackend(multiBackend)
	} else {
		fileLogger.SetBackend(fileLoggerLeveled)
	}
}

func createFileLogger(name string, _ int) logging.Backend {
	name = getLogFile(name)
	return logging.NewLogBackend(&lumberjack.Logger{
		Filename:   name,
		MaxSize:    100, // megabytes
		MaxBackups: 500,
		MaxAge:     500, //days
	}, "", 0)
}

func getLogFile(fileName string) string {
	return fileName
}

func PanicLog(args ...interface{}) {
	errorLogger.Panic(args...)
	log.Panic(args...)
}

func Log(args ...interface{}) {
	info := getCallerInfo()
	message := fmt.Sprint(args...)
	infoLogger.Info(logFormatter(info, message))
}

func LogError(args ...interface{}) {
	info := getCallerInfo()
	message := fmt.Sprint(args...)
	errorLogger.Error(logFormatter(info, message))
}

func LogDebug(args ...interface{}) {
	info := getCallerInfo()
	message := fmt.Sprint(args...)
	debugLogger.Info(logFormatter(info, message))
}

func LogChat(args ...interface{}) {
	info := getCallerInfo()
	message := fmt.Sprint(args...)
	chatLogger.Info(logFormatter(info, message))
}

func logGin(args ...interface{}) {
	message := fmt.Sprint(args...)
	ginLogger.Info(logFormatter("GIN", message))
}

func logGorm(args ...interface{}) {
	message := fmt.Sprint(args...)
	gormLogger.Info(logFormatter("GORM", message))
}

func logErrorGorm(args ...interface{}) {
	message := fmt.Sprint(args...)
	gormErrorLogger.Info(logFormatter("GORM", message))
}

func logDebugGorm(args ...interface{}) {
	message := fmt.Sprint(args...)
	gormDebugLogger.Info(logFormatter("GORM", message))
}

func LogRequest(args ...interface{}) {
	message := fmt.Sprint(args...)
	requestsLogger.Info(logFormatter("REQUEST", message))
}

func LogESRequest(args ...interface{}) {
	message := fmt.Sprint(args...)
	esRequestsLogger.Info(logFormatter("", message))
}

func LogESProxyRequest(args ...interface{}) {
	message := fmt.Sprint(args...)
	esProxyRequestsLogger.Info(logFormatter("", message))
}

func LogResponse(args ...interface{}) {
	message := fmt.Sprint(args...)
	responsesLogger.Info(logFormatter("GIN", message))
}

func logFormatter(info, message string) interface{} {
	return fmt.Sprintf("%s] %s", info, message)
}

func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", file, line)
}

type EsGinWriter struct{}

func NewEsGinWriter() io.Writer {
	return &EsGinWriter{}
}

func (w *EsGinWriter) Write(p []byte) (n int, err error) {
	logGin(string(p)) // Используйте ваш метод Log или другой, как вам нужно
	return len(p), nil
}

type EsGormLogger struct{}

func (gl *EsGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return gl
}

func (gl *EsGormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	logGorm(msg, data)
}

func (gl *EsGormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	logErrorGorm(msg, data)
}

func (gl *EsGormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	logErrorGorm(msg, data)
}

func (gl *EsGormLogger) Trace(_ context.Context, start time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(start)
	sql, rows := fc()
	if err != nil {
		logDebugGorm(fmt.Sprintf("%s | Rows affected: %d | Duration: %s | Error: %s", sql, rows, elapsed, err.Error()))
	} else {
		logDebugGorm(fmt.Sprintf("%s | Rows affected: %d | Duration: %s", sql, rows, elapsed))
	}
}
