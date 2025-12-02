package ratesparser

import (
	"encoding/json"
	"fmt"
)

type RapiraResponse struct {
	Data      []RapiraResponseData `json:"data"`
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	IsWorking int                  `json:"isWorking"`
}
type RapiraResponseData struct {
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
}

func addRapiraRubRate(newRates map[string]map[string]float64) error {
	body, err := doRequest(parserCfg["rapira"])
	if err != nil {
		return fmt.Errorf("rapira 1: %s", err.Error())
	}

	var obj RapiraResponse
	if err := json.Unmarshal(body, &obj); err != nil {
		return fmt.Errorf("rapira 1.1: %s", err.Error())
	}

	for _, item := range obj.Data {
		if item.Symbol != "USDT/RUB" {
			continue
		}

		if _, ok := newRates["USD"]; !ok {
			newRates["USD"] = make(map[string]float64)
		}
		if _, ok := newRates["USDT"]; !ok {
			newRates["USDT"] = make(map[string]float64)
		}
		newRates["USD"]["RUB"] = item.Low
		newRates["USDT"]["RUB"] = item.Low
		return nil
	}

	return fmt.Errorf("rapira 2: %s", "rate not exist")
}
