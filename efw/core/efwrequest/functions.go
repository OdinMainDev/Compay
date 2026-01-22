package efwrequest

import (
	"efw/config"
	"efw/utils"
	"fmt"
	"net/http"
	"strings"
)

var ErrorInvalidProxyAuthFormat = fmt.Errorf("%s", "[EFW: REQUEST] InvalidProxyAuthFormat")
var ErrorInvalidProxyHostFormat = fmt.Errorf("%s", "[EFW: REQUEST] InvalidProxyHostFormat")
var ErrorInvalidProxyFormat = fmt.Errorf("%s", "[EFW: REQUEST] InvalidProxyFormat")

type logMessageContext struct {
	Stage        string
	RequestData  RequestHelper
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte
	Proxy        Proxy
	Err          error
}

const (
	logStagePrepareClient = "prepare_client"
	logStageDoRequest     = "do_request"
	logStageReadBody      = "read_body"
	logStageSuccess       = "success"
)

func applyHeaders(req *http.Request, data RequestHelper) {
	hasContentType := false
	for name, val := range data.Headers {
		if utils.ToLower(name) == "Content-Type" {
			hasContentType = true
		}
		req.Header.Set(name, val)
	}

	if !hasContentType {
		if data.IsJsonRequest {
			req.Header.Set("Content-Type", "application/json")
		} else {
			if data.Method == "POST" {
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			}
		}
	}
}

func getAvailProxy() (string, string) {
	if len(config.Proxies) == 0 {
		return "", ""
	}
	proxies := utils.ShuffleMap(config.Proxies)

	toIndex := len(proxies)
	if toIndex > config.ProxyMaxCheckCount {
		toIndex = config.ProxyMaxCheckCount
	}

	for i := 0; i < toIndex; i++ {
		if err := CheckProxy(proxies[i].Key, proxies[i].Value); err != nil {
			continue
		}

		return proxies[i].Key, proxies[i].Value
	}

	return "", ""
}

func ParseProxy(proxyString string) (resp Proxy, err error) {
	parts := utils.Explode(proxyString, "@")
	if l := len(parts); l == 0 || l > 2 {
		return resp, ErrorInvalidProxyFormat
	}

	// обработка прокси с авторизацией
	if len(parts) == 2 {
		authParts := utils.Explode(parts[0], ":")
		if len(authParts) != 2 {
			return resp, ErrorInvalidProxyAuthFormat
		}
		hostParts := utils.Explode(parts[1], ":")
		if len(hostParts) != 2 {
			return resp, ErrorInvalidProxyHostFormat
		}
		resp = Proxy{
			Username: authParts[0],
			Password: authParts[1],
			Host:     hostParts[0],
			Port:     hostParts[1],
		}
		return resp, nil
	}

	// обработка прокси без авторизации
	hostParts := strings.Split(parts[0], ":")
	if len(hostParts) != 2 {
		return resp, ErrorInvalidProxyHostFormat
	}
	resp = Proxy{
		Host: hostParts[0],
		Port: hostParts[1],
	}
	return resp, nil
}

func constructLogMessage(ctx logMessageContext) string {
	var logBuilder strings.Builder

	logBuilder.WriteString("REQUEST LOG\n")
	logBuilder.WriteString("==============================================\n")

	if ctx.Stage != "" {
		fmt.Fprintf(&logBuilder, "Stage: %s\n", ctx.Stage)
	}
	if ctx.Err != nil {
		fmt.Fprintf(&logBuilder, "Error: %s\n", ctx.Err)
	}

	writeRequestLog(&logBuilder, ctx)
	writeResponseLog(&logBuilder, ctx)

	logBuilder.WriteString("==============================================\n")

	return logBuilder.String()
}

func writeRequestLog(logBuilder *strings.Builder, ctx logMessageContext) {
	logBuilder.WriteString("Request:\n")

	req := ctx.Request
	if req != nil && req.URL != nil {
		fmt.Fprintf(logBuilder, "  Method: %s\n", req.Method)
		fmt.Fprintf(logBuilder, "  URL: %s\n", req.URL.String())

		if scheme := req.URL.Scheme; scheme != "" {
			fmt.Fprintf(logBuilder, "  Scheme: %s\n", scheme)
		}
		if host := req.URL.Host; host != "" {
			fmt.Fprintf(logBuilder, "  Host: %s\n", host)
		}
		if path := req.URL.Path; path != "" {
			fmt.Fprintf(logBuilder, "  Path: %s\n", path)
		}
		if query := req.URL.RawQuery; query != "" {
			fmt.Fprintf(logBuilder, "  Query: %s\n", query)
		}
		if req.ContentLength > 0 {
			fmt.Fprintf(logBuilder, "  ContentLength: %d\n", req.ContentLength)
		}
		if req.Proto != "" {
			fmt.Fprintf(logBuilder, "  Protocol: %s\n", req.Proto)
		}
	} else {
		fmt.Fprintf(logBuilder, "  URL: %s\n", ctx.RequestData.Url)
		fmt.Fprintf(logBuilder, "  Method: %s\n", ctx.RequestData.getMethod())
	}

	fmt.Fprintf(
		logBuilder,
		"  Flags: json=%t, cookies=%t, proxy_enabled=%t\n",
		ctx.RequestData.IsJsonRequest,
		ctx.RequestData.HasCookies,
		ctx.RequestData.IsProxyEnable,
	)

	if ctx.RequestData.ProxyType != "" {
		fmt.Fprintf(logBuilder, "  Proxy type requested: %s\n", ctx.RequestData.ProxyType)
	}

	proxyToLog := ctx.Proxy
	if proxyToLog == (Proxy{}) {
		proxyToLog = ctx.RequestData.ProxyData
	}

	if proxyInfo := formatProxy(proxyToLog); proxyInfo != "" {
		fmt.Fprintf(logBuilder, "  Proxy: %s\n", proxyInfo)
		if proxyToLog.Type != "" {
			fmt.Fprintf(logBuilder, "  ProxyType: %s\n", proxyToLog.Type)
		}
	}

	if req != nil && len(req.Header) > 0 {
		logBuilder.WriteString("  Headers:\n")
		for name, values := range req.Header {
			for _, value := range values {
				fmt.Fprintf(logBuilder, "    %s: %s\n", name, value)
			}
		}
	} else if len(ctx.RequestData.Headers) > 0 {
		logBuilder.WriteString("  Headers:\n")
		for name, value := range ctx.RequestData.Headers {
			fmt.Fprintf(logBuilder, "    %s: %s\n", name, value)
		}
	}

	if len(ctx.RequestData.Body) > 0 {
		fmt.Fprintf(logBuilder, "  Body (%d bytes):\n%s\n", len(ctx.RequestData.Body), string(ctx.RequestData.Body))
	}

	logBuilder.WriteString("\n")
}

func writeResponseLog(logBuilder *strings.Builder, ctx logMessageContext) {
	logBuilder.WriteString("Response:\n")

	if ctx.Response != nil {
		statusText := ctx.Response.Status
		if statusText == "" {
			statusText = http.StatusText(ctx.Response.StatusCode)
		}

		fmt.Fprintf(logBuilder, "  Status: %d %s\n", ctx.Response.StatusCode, statusText)
		if ctx.Response.ContentLength > 0 {
			fmt.Fprintf(logBuilder, "  ContentLength: %d\n", ctx.Response.ContentLength)
		}

		if len(ctx.Response.Header) > 0 {
			logBuilder.WriteString("  Headers:\n")
			for name, values := range ctx.Response.Header {
				for _, value := range values {
					fmt.Fprintf(logBuilder, "    %s: %s\n", name, value)
				}
			}
		}
	} else if ctx.Err != nil {
		logBuilder.WriteString("  Status: not available (request failed before response)\n")
	}

	if len(ctx.ResponseBody) > 0 {
		fmt.Fprintf(logBuilder, "  Body (%d bytes):\n%s\n", len(ctx.ResponseBody), string(ctx.ResponseBody))
	}

	logBuilder.WriteString("\n")
}

func formatProxy(proxyData Proxy) string {
	if proxyData.Host == "" && proxyData.Port == "" {
		return ""
	}

	if proxyData.Username != "" || proxyData.Password != "" {
		return fmt.Sprintf("%s:%s@%s:%s", proxyData.Username, proxyData.Password, proxyData.Host, proxyData.Port)
	}

	return fmt.Sprintf("%s:%s", proxyData.Host, proxyData.Port)
}
