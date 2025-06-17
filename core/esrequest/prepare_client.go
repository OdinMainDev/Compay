package esrequest

import (
	"compay/core/config"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func prepareClient(data RequestHelper) (*http.Client, Proxy, error) {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   time.Duration(config.RequestTimeout) * time.Second,
			KeepAlive: time.Duration(config.RequestTimeout) * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   time.Duration(config.RequestTimeout) * time.Second,
		ResponseHeaderTimeout: time.Duration(config.RequestTimeout) * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	proxyData, err := addProxy(transport, data)
	if err != nil {
		return nil, Proxy{}, err
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(config.RequestTimeout) * time.Second,
	}

	if !data.IsWithoutJar {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return nil, proxyData, config.ErrorProxyJarError(err.Error())
		}
		client.Jar = jar
	}

	return client, proxyData, nil
}
