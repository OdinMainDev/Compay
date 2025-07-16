package rates

import (
	"encoding/json"
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
	body, err := doRequest(parserCfg["grinex"])
	if err != nil {
		return getUsdToRubRateFromRapira()
	}
	var data GarantexResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return getUsdToRubRateFromRapira()
	}
	if len(data) == 0 {
		return getUsdToRubRateFromRapira()
	}
	return data[0].Close, nil
}
