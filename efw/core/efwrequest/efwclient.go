package efwrequest

import (
	"efw/config"
	"fmt"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var ErrorCreateJar = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: REQUEST] ErrorCreateJar: ", err)
}

func newTransport(requestTimeout time.Duration) *http.Transport {
	if requestTimeout == 0 {
		requestTimeout = time.Duration(config.RequestTimeout)
	}
	return &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   requestTimeout * time.Second,
			KeepAlive: requestTimeout * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   requestTimeout * time.Second,
		ResponseHeaderTimeout: requestTimeout * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func NewClient(data RequestHelper) (*http.Client, Proxy, error) {
	requestTimeout := data.RequestTimeout
	if requestTimeout == 0 {
		requestTimeout = time.Duration(config.RequestTimeout)
	}
	transport := newTransport(requestTimeout)
	proxyData, err := addProxy(transport, data)
	if err != nil {
		return nil, Proxy{}, err
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   requestTimeout * time.Second,
	}

	if data.HasCookies {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return nil, proxyData, ErrorCreateJar(err.Error())
		}
		client.Jar = jar
	}

	return client, proxyData, nil
}
