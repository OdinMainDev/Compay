package efwrequest

import (
	"bytes"
	"efw/core/efwlog"
	"fmt"
	"net/http"
)

var ErrorCreateNewRequest = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: REQUEST] CreateNewRequest: ", err)
}

func Request(data RequestHelper) ([]byte, *http.Response, error) {
	req, err := http.NewRequest(data.getMethod(), data.Url, bytes.NewBuffer(data.Body))
	if err != nil {
		return []byte{}, nil, ErrorCreateNewRequest(err.Error())
	}
	req.Close = true
	applyHeaders(req, data)
	return doRequest(req, data)
}

func doRequest(req *http.Request, data RequestHelper) ([]byte, *http.Response, error) {
	logMessage := ""
	defer func() {
		if logMessage == "" || data.IsLogDisabled {
			return
		}
		if data.isProxyCheck {
			efwlog.LogEFWProxyRequest(logMessage)
		} else {
			efwlog.LogEFWRequest(logMessage)
		}
	}()

	httpClient, proxyData, err := NewClient(data)
	if err != nil {
		logMessage = constructLogMessage(logMessageContext{
			Stage:       logStagePrepareClient,
			RequestData: data,
			Request:     req,
			Proxy:       proxyData,
			Err:         err,
		})
		return []byte{}, nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		logMessage = constructLogMessage(logMessageContext{
			Stage:       logStageDoRequest,
			RequestData: data,
			Request:     req,
			Proxy:       proxyData,
			Err:         err,
		})
		return []byte{}, nil, err
	}

	body, err := DecodeRequest(resp)
	if err != nil {
		logMessage = constructLogMessage(logMessageContext{
			Stage:       logStageReadBody,
			RequestData: data,
			Request:     req,
			Response:    resp,
			Proxy:       proxyData,
			Err:         err,
		})
		return []byte{}, nil, err
	}

	logMessage = constructLogMessage(logMessageContext{
		Stage:        logStageSuccess,
		RequestData:  data,
		Request:      req,
		Response:     resp,
		ResponseBody: body,
		Proxy:        proxyData,
	})
	return body, resp, nil
}
