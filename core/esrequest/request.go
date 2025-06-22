package esrequest

import (
	"bytes"
	"compay/core/eslog"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RequestHelper struct {
	Url           string
	Method        string
	IsJson        bool
	IsSafeRequest bool
	IsWithoutJar  bool
	ProxyType     string
	ProxyData     Proxy
	IsProxyCheck  bool
	Body          []byte
	Headers       map[string]string
	DisableLog    bool
}

func Request(data RequestHelper) ([]byte, int, error) {
	method := strings.ToUpper(data.Method)
	if method != "POST" {
		method = "GET"
	}
	req, err := http.NewRequest(method, data.Url, bytes.NewBuffer(data.Body))
	if err != nil {
		return []byte{}, 0, err
	}
	req.Close = true

	if data.IsJson {
		req.Header.Set("Content-Type", "application/json")
	} else {
		if method == "POST" {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	}
	for name, val := range data.Headers {
		req.Header.Set(name, val)
	}
	return doRequest(req, data)
}

func doRequest(req *http.Request, data RequestHelper) ([]byte, int, error) {
	logMessage := ""
	defer func() {
		if logMessage == "" || data.DisableLog {
			return
		}
		if data.IsProxyCheck {
			eslog.LogESProxyRequest(logMessage)
		} else {
			eslog.LogESRequest(logMessage)
		}
	}()

	httpClient, proxyData, err := prepareClient(data)
	if err != nil {
		logMessage = constructLogMessage(data, req, proxyData, err)
		return []byte{}, 0, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		logMessage = constructLogMessage(data, req, -666, proxyData, err)
		return []byte{}, 0, err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		logMessage = constructLogMessage(data, req, -777, proxyData, err)
		return []byte{}, 0, err
	}

	logMessage = constructLogMessage(data, req, resp.StatusCode, body, proxyData, nil)
	return body, resp.StatusCode, nil
}

func constructLogMessage(args ...any) string {
	if len(args) == 0 || len(args) == 1 {
		return ""
	}
	var logBuilder bytes.Buffer
	logBuilder.WriteString("REQUEST:\n==============================================\n")
	if len(args) == 4 || len(args) == 5 {
		logBuilder.WriteString("!!!!!!   REQUEST WITH ERROR   !!!!!!\n")
	}

	var data RequestHelper
	var req *http.Request
	var statusCode int
	var response []byte
	var proxyData Proxy
	var err error

	data = args[0].(RequestHelper)
	req = args[1].(*http.Request)
	switch len(args) {
	case 4:
		proxyData = args[2].(Proxy)
		err = args[3].(error)
	case 5:
		statusCode = args[2].(int)
		proxyData = args[3].(Proxy)
		err = args[4].(error)
	case 6:
		statusCode = args[2].(int)
		response = args[3].([]byte)
		proxyData = args[4].(Proxy)
		err = nil
	}
	if err != nil {
		logBuilder.WriteString("ERROR: " + err.Error() + "\n")
	}
	logBuilder.WriteString("Url: " + data.Url + "\n")
	logBuilder.WriteString("StatusCode: " + fmt.Sprint(statusCode) + "\n")
	logBuilder.WriteString("Method: " + fmt.Sprint(data.Method) + "\n")
	logBuilder.WriteString("IsJson: " + fmt.Sprint(data.IsJson) + "\n")
	logBuilder.WriteString("IsSafeRequest: " + fmt.Sprint(data.IsSafeRequest) + "\n")
	logBuilder.WriteString("IsWithoutJAR: " + fmt.Sprint(data.IsWithoutJar) + "\n")

	// БЛОК ПРО ПРОКСИ
	if proxyData.Username != "" {
		logBuilder.WriteString("\nProxyType: " + proxyData.Type + "\n")
		logBuilder.WriteString("Proxy: " + fmt.Sprint(proxyData.Username, ":", proxyData.Password, "@", proxyData.Host, ":", proxyData.Port) + "\n")
	}
	if proxyData.Username == "" && proxyData.Host != "" {
		logBuilder.WriteString("\nProxyType: " + proxyData.Type + "\n")
		logBuilder.WriteString("Proxy: " + fmt.Sprint(proxyData.Host, ":", proxyData.Port) + "\n")
	}

	// БЛОК ПРО ХИДЕРЫ
	if len(req.Header) > 0 {
		logBuilder.WriteString("\nHeaders:\n")
		for name, values := range req.Header {
			for _, value := range values {
				logBuilder.WriteString(fmt.Sprintf("%s: %s\n", name, value))
			}
		}
	}

	// БЛОКИ БАДИ И РЕСПОНЗ
	if len(data.Body) > 0 {
		logBuilder.WriteString("\nBody: " + string(data.Body) + "\n")
	}

	if len(response) != 0 {
		logBuilder.WriteString("\nResponse: " + string(response) + "\n")
	}
	logBuilder.WriteString("==============================================\n")

	return logBuilder.String()
}
