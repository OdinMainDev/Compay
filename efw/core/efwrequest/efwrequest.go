package efwrequest

import (
	"strings"
	"time"
)

type RequestHelper struct {
	Url           string
	Method        string
	IsLogDisabled bool

	IsJsonRequest bool
	HasCookies    bool
	Headers       map[string]string
	Body          []byte

	IsProxyEnable bool
	ProxyType     string
	ProxyData     Proxy
	isProxyCheck  bool

	RequestTimeout time.Duration
}

func (that *RequestHelper) getMethod() string {
	method := strings.ToUpper(that.Method)
	if method != "POST" {
		method = "GET"
	}
	return method
}
