package esrequest

import (
	"compay/core/config"
	"context"
	"net/http"
	"time"
)

func CheckProxy(proxy, proxyType string) error {
	proxyData, err := ParseProxy(proxy)
	if err != nil {
		return err
	}
	proxyData.Type = proxyType

	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.RequestTimeout))
	defer ctxCancel()
	req, err := http.NewRequestWithContext(ctx, "GET", config.ProxyCheckTarget, nil)
	if err != nil {
		return config.ErrorProxyCreateRequestError(err.Error())
	}
	if _, _, err := doRequest(req, RequestHelper{
		ProxyData:    proxyData,
		ProxyType:    proxyType,
		IsProxyCheck: true,
	}); err != nil {
		return err
	}

	return nil
}
