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

func newTransport() *http.Transport {
	return &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   time.Duration(config.RequestTimeout) * time.Second,
			KeepAlive: time.Duration(config.RequestTimeout) * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   time.Duration(config.RequestTimeout) * time.Second,
		ResponseHeaderTimeout: time.Duration(config.RequestTimeout) * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func NewClient(data RequestHelper) (*http.Client, Proxy, error) {
	transport := newTransport()
	proxyData, err := addProxy(transport, data)
	if err != nil {
		return nil, Proxy{}, err
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(config.RequestTimeout) * time.Second,
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
