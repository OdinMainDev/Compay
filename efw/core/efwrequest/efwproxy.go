package efwrequest

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

type Proxy struct {
	Username string
	Password string
	Host     string
	Port     string
	Type     string
}

var ErrorProxyNotConnected = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: REQUEST] ProxyNotConnected: ", err)
}

func addProxy(transport *http.Transport, data RequestHelper) (Proxy, error) {
	if !data.IsProxyEnable {
		return Proxy{}, nil
	}

	if data.ProxyData.Host != "" {
		if data.ProxyType == "http" {
			return addProxyHttp(transport, data.ProxyData)
		}
		return addProxySocks(transport, data.ProxyData)
	}

	proxyAddr, proxyType := getAvailProxy()
	if proxyAddr == "" {
		// TODO: возможно что тут стоит сделать ошибку но это не точно ( так как тут в случае если проксей нет совсем только такое )
		return Proxy{}, nil
	}

	proxyData, err := ParseProxy(proxyAddr)
	proxyData.Type = proxyType
	if err != nil {
		return Proxy{}, err
	}
	if proxyType == "http" {
		return addProxyHttp(transport, proxyData)
	}
	return addProxySocks(transport, proxyData)
}

func addProxyHttp(transport *http.Transport, proxyData Proxy) (Proxy, error) {
	proxyURL := &url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort(proxyData.Host, proxyData.Port),
	}
	if proxyData.Username != "" && proxyData.Password != "" {
		proxyURL.User = url.UserPassword(proxyData.Username, proxyData.Password)
	}
	transport.Proxy = http.ProxyURL(proxyURL)
	return proxyData, nil
}

func addProxySocks(transport *http.Transport, proxyData Proxy) (Proxy, error) {
	host := net.JoinHostPort(proxyData.Host, proxyData.Port)
	var auth *proxy.Auth = nil
	if proxyData.Username != "" && proxyData.Password != "" {
		auth = &proxy.Auth{
			User:     proxyData.Username,
			Password: proxyData.Password,
		}
	}

	dialer, err := proxy.SOCKS5("tcp", host, auth, proxy.Direct)
	if err != nil {
		return Proxy{}, ErrorProxyNotConnected(err.Error())
	}
	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial(network, addr)
	}
	return proxyData, nil
}
