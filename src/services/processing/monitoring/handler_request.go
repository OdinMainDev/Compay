package monitoring

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/helpers"
	"efw/core/efwlog"
	"efw/core/efwrequest"
	"encoding/json"
	"fmt"
	"net/http"
)

type requestUrl string

const (
	requestGetBlockCount   requestUrl = "/get-block-count"
	requestGetTransactions requestUrl = "/get-transactions"
)

func (that *monitoring) doRequest(path requestUrl, codeAlpha types.CurrencyCodeAlpha, body []byte, obj any) error {
	url := that.getRequestUrl(path, codeAlpha)
	body, resp, err := efwrequest.Request(efwrequest.RequestHelper{
		Url:           url,
		Method:        "POST",
		Body:          body,
		IsLogDisabled: false,

		RequestTimeout: 180,
	})

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		helpers.NotificationErrorProcesssingResponse(url, string(body))
		return cperrors.ErrorMonitoringBadResponseStatusCode(resp.StatusCode)
	}
	if err := json.Unmarshal(body, obj); err != nil {
		helpers.NotificationErrorProcesssingResponse(url, string(body))
		return cperrors.ErrorMonitoringBadResponseBody(err.Error())
	}

	return nil
}

func (that *monitoring) getRequestUrl(path requestUrl, codeAlpha types.CurrencyCodeAlpha) string {
	switch codeAlpha {
	case constants.BtcCodeAlpha:
		return constants.BtcApiUrl + string(path)

	default:
		helpers.NotificationError(fmt.Sprint("BAD CONFIGURATION getRequestUrl: ", codeAlpha))
		efwlog.Panic("BAD CONFIGURATION getRequestUrl: ", codeAlpha)
		return ""
	}
}
