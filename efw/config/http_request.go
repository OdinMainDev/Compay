package config

import "sync"

var Proxies = map[string]string{}
var ProxiesMutex = &sync.Mutex{}

var ProxyMaxCheckCount = 5
var ProxyCheckTarget = "https://api.ipify.org/"

var RequestTimeout = 30
