package cconverter

import (
	"compay/common/cperrors"
	"efw/core/efwrequest"
	"encoding/json"
	"time"
)

type response struct {
	Data    map[string]map[string]float64 `json:"data"`
	Success bool                          `json:"success"`
}

// fetchRates загружает курсы валют из API
func (c *CConverter) fetchRates() error {
	body, _, err := efwrequest.Request(efwrequest.RequestHelper{
		Url:           apiURL,
		Method:        "GET",
		Headers:       map[string]string{},
		IsLogDisabled: true,
	})
	if err != nil {
		return cperrors.ErrorCConverterRequestFail(err.Error())
	}

	var apiResponse response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return cperrors.ErrorCConverterInvalidResponse(err.Error())
	}

	if !apiResponse.Success {
		return cperrors.ErrorCConverterBadResponse(string(body))
	}

	c.mutex.Lock()
	c.rates = apiResponse.Data
	c.lastUpdate = time.Now()
	c.mutex.Unlock()

	return nil
}
