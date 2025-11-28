package rates

import (
	"encoding/json"
	"fmt"
)

type GarantexResponse []struct {
	Close  float64 `json:"close"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Open   float64 `json:"open"`
	Time   int     `json:"time"`
	Volume float64 `json:"volume"`
}

func getUsdToRubRateFromGrinex() (float64, error) {
	return 0, fmt.Errorf("notWorking")
	body, err := doRequest(parserCfg["grinex"])
	if err != nil {
		return 0, fmt.Errorf("grinex 1: %s", err.Error())
	}
	var data GarantexResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, fmt.Errorf("grinex 2: %s | body: %s", err.Error(), string(body))
	}
	if len(data) == 0 {
		return 0, fmt.Errorf("grinex 3: %s | body: %s", "nil array", string(body))
	}
	return data[0].Close, nil
}
