package ratesparser

import (
	"efw/utils"
	"encoding/json"
	"fmt"
	"strconv"
)

type KucoinResponse struct {
	Code string             `json:"code"`
	Data KucoinResponseData `json:"data"`
}

type KucoinResponseData struct {
	Ticker []KucoinResponseTicker `json:"ticker"`
}

type KucoinResponseTicker struct {
	Symbol string `json:"symbol"`
	Last   string `json:"last"`
}

func addKucoinRates(newRates map[string]map[string]float64) error {
	body, err := doRequest(parserCfg["kucoin"])
	if err != nil {
		return fmt.Errorf("kucoin 1: %s", err.Error())
	}
	var obj KucoinResponse
	if err := json.Unmarshal(body, &obj); err != nil {
		return fmt.Errorf("kucoin 1.1: %s", err.Error())
	}

	for _, ticker := range obj.Data.Ticker {
		parts := utils.Explode("-", ticker.Symbol)
		if len(parts) != 2 {
			continue
		}

		currencyFrom, currencyTo := parts[0], parts[1]
		if !utils.InArray(currencyFrom, currencies) || !utils.InArray(currencyTo, currencies) {
			continue
		}

		rate, err := strconv.ParseFloat(ticker.Last, 64)
		if err != nil {
			return fmt.Errorf("kucoin 2: %s", err.Error())
		}
		if _, ok := newRates[currencyFrom][currencyTo]; !ok {
			newRates[currencyFrom][currencyTo] = rate
		}
	}

	return nil
}
