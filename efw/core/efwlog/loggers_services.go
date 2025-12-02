package efwlog

func LogChat(args ...any) {
	LogByLoggerName("chat", args...)
}

func LogEvent(args ...any) {
	LogByLoggerName("events", args...)
}

func LogCoraza(args ...any) {
	LogByLoggerName("coraza", args...)
}

func LogRequest(args ...any) {
	LogByLoggerName("request", args...)
}

func LogResponse(args ...any) {
	LogByLoggerName("response", args...)
}

func LogGameRequest(args ...any) {
	LogByLoggerName("game_request", args...)
}

func LogEFWRequest(args ...any) {
	LogByLoggerName("efw_request", args...)
}

func LogEFWProxyRequest(args ...any) {
	LogByLoggerName("efw_proxy_request", args...)
}
