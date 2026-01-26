package efwlog

func LogCoraza(args ...any) {
	LogByLoggerName("coraza", 1, args...)
}

func LogEFWRequest(args ...any) {
	LogByLoggerName("efw_request", 1, args...)
}

func LogEFWProxyRequest(args ...any) {
	LogByLoggerName("efw_proxy_request", 1, args...)
}
