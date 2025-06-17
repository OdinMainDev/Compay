package config

import (
	"fmt"
	"sync"
)

var Proxies = map[string]string{}
var ProxiesMutex = &sync.Mutex{}

var ProxyMaxCheckCount = 5
var ProxyCheckTarget = "http://checkip.amazonaws.com"

var RequestTimeout = 30

var ErrorInvalidProxyFormat = func(proxy string) error {
	return fmt.Errorf("%s: %s", "InvalidProxyFormat", proxy)
}

var ErrorProxyInvalidAuthFormat = fmt.Errorf("%s", "InvalidAuthFormat")
var ErrorProxyInvalidHostFormat = fmt.Errorf("%s", "InvalidHostFormat")
var ErrorProxyInvalidProxyFormat = fmt.Errorf("%s", "InvalidProxyFormat")

var ErrorProxyNotConnected = func(data string) error {
	return fmt.Errorf("%s: %s", "NotConnected | AdditionalData:", data)
}

var ErrorProxyJarError = func(data string) error {
	return fmt.Errorf("%s: %s", "JarError | AdditionalData:", data)
}

var ErrorProxyCreateRequestError = func(data string) error {
	return fmt.Errorf("%s: %s", "CreateRequestError | AdditionalData:", data)
}
