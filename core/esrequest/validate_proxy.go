package esrequest

import (
	"compay/core/config"
	"compay/core/utils"
	"strings"
)

type Proxy struct {
	Username string
	Password string
	Host     string
	Port     string
	Type     string
}

func ParseProxy(proxyString string) (resp Proxy, err error) {
	parts := utils.Explode(proxyString, "@")
	if len(parts) == 2 {
		authParts := utils.Explode(parts[0], ":")
		if len(authParts) != 2 {
			return resp, config.ErrorProxyInvalidAuthFormat
		}
		hostParts := utils.Explode(parts[1], ":")
		if len(hostParts) != 2 {
			return resp, config.ErrorProxyInvalidHostFormat
		}
		resp = Proxy{
			Username: authParts[0],
			Password: authParts[1],
			Host:     hostParts[0],
			Port:     hostParts[1],
		}
	} else if len(parts) == 1 {
		hostParts := strings.Split(parts[0], ":")
		if len(hostParts) != 2 {
			return resp, config.ErrorProxyInvalidHostFormat
		}
		resp = Proxy{
			Host: hostParts[0],
			Port: hostParts[1],
		}
	} else {
		return resp, config.ErrorInvalidProxyFormat(proxyString)
	}
	return resp, nil
}
