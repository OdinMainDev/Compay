package ratesparser

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

var (
	muExchangeRate       sync.Mutex
	exchangeRateResponse ExchangeRateResponse
)

type ExchangeRateResponse struct {
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}

func getExchangeRateRate(curreny string) (float64, error) {
	muExchangeRate.Lock()
	defer muExchangeRate.Unlock()
	if exchangeRateResponse.Timestamp != 0 {
		if time.Now().Unix()-exchangeRateResponse.Timestamp > 3600 {
			return exchangeRateResponse.Rates[curreny], nil
		}
	}

	body, err := doRequest(parserCfg["exchangerate"])
	if err != nil {
		return 0, fmt.Errorf("exchangerate 1: %s", err.Error())
	}
	var obj ExchangeRateResponse
	if err := json.Unmarshal(body, &obj); err != nil {
		return 0, fmt.Errorf("exchangerate 1.1: %s", err.Error())
	}

	exchangeRateResponse = obj
	return exchangeRateResponse.Rates[curreny], nil
}
