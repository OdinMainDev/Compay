package ratesparser

import (
	"efw/utils"
	"encoding/json"
	"fmt"
	"strconv"
)

type OkxResponse struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []OkxResponseData `json:"data"`
}

type OkxResponseData struct {
	InstId  string `json:"instId"`
	IdxPx   string `json:"idxPx"`
	High24H string `json:"high24h"`
	SodUtc0 string `json:"sodUtc0"`
	Open24H string `json:"open24h"`
	Low24H  string `json:"low24h"`
	SodUtc8 string `json:"sodUtc8"`
	Ts      string `json:"ts"`
}

func addOkxRates(newRates map[string]map[string]float64) error {
	body, err := doRequest(parserCfg["okx"])
	if err != nil {
		return fmt.Errorf("okx 1: %s", err.Error())
	}
	var obj OkxResponse
	if err := json.Unmarshal(body, &obj); err != nil {
		return fmt.Errorf("okx 1.1: %s", err.Error())
	}

	for _, ticker := range obj.Data {
		parts := utils.Explode("-", ticker.InstId)
		if len(parts) != 2 {
			continue
		}

		currencyFrom, currencyTo := parts[0], parts[1]
		if !utils.InArray(currencyFrom, currencies) || !utils.InArray(currencyTo, currencies) {
			continue
		}

		rate, err := strconv.ParseFloat(ticker.IdxPx, 64)
		if err != nil {
			return fmt.Errorf("okx 2: %s", err.Error())
		}

		newRates[currencyFrom][currencyTo] = rate
	}

	return nil
}
