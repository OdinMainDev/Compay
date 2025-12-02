package efwrequest

import (
	"efw/config"
	"fmt"
	"net/http"
)

var ErrorInvalidProxyType = func(proxyType string) error {
	return fmt.Errorf("%s %s", "[EFW: REQUEST] InvalidProxyType: ", proxyType)
}

func CheckProxy(proxyType, proxy string) error {
	if proxyType != "http" && proxyType != "socks" {
		return ErrorInvalidProxyType(proxyType)
	}

	proxyData, err := ParseProxy(proxy)
	if err != nil {
		return err
	}
	proxyData.Type = proxyType

	req, err := http.NewRequest("GET", config.ProxyCheckTarget, nil)
	if err != nil {
		return ErrorCreateNewRequest(err.Error())
	}

	if _, _, err := doRequest(req, RequestHelper{
		ProxyData:    proxyData,
		ProxyType:    proxyType,
		isProxyCheck: true,
	}); err != nil {
		return err
	}

	return nil
}
