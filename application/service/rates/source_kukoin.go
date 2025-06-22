package rates

import (
	"compay/core/utils"
	"encoding/json"
	"strconv"
)

type KucoinResp struct {
	Code string `json:"code"`
	Data struct {
		Ticker []struct {
			Symbol string `json:"symbol"`
			Last   string `json:"last"`
		} `json:"ticker"`
	} `json:"data"`
}

func getKucoin() (map[string]map[string]float64, error) {
	body, err := doRequest(parserCfg["kucoin"])
	if err != nil {
		return nil, err
	}
	var obj KucoinResp
	if err := json.Unmarshal(body, &obj); err != nil {
		return nil, err
	}
	response := map[string]map[string]float64{}
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
			return nil, err
		}
		if _, ok := response[currencyFrom]; !ok {
			response[currencyFrom] = map[string]float64{}
		}
		response[currencyFrom][currencyTo] = rate
	}

	if _, ok := response["USDT"]["USDT"]; !ok {
		response["USDT"]["USDT"] = 1.0
	}

	if _, ok := response["DAI"]; !ok {
		response["DAI"] = map[string]float64{}
	}
	response["DAI"]["USDT"] = 1
	response["DAI"]["DAI"] = 1

	if _, ok := response["USD"]; !ok {
		response["USD"] = map[string]float64{}
	}
	response["USD"]["USDT"] = 1
	response["USD"]["USD"] = 1

	if _, ok := response["EUR"]; !ok {
		response["EUR"] = map[string]float64{}
	}
	response["EUR"]["EUR"] = 1

	if v, ok := response["USDT"]["EUR"]; ok {
		response["USD"]["EUR"] = v
	}

	return response, nil
}
