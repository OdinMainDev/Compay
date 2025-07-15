package rates

import (
	"encoding/json"
)

func getUsdToRubRateFromRapira() (float64, error) {
	body, err := doRequest(parserCfg["rapira"])
	if err != nil {
		return getUsdToRubRateFromCbr()
	}
	var data []map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return getUsdToRubRateFromCbr()
	}
	for _, d := range data {
		if d["symbol"] == "USDT/RUB" {
			if val, ok := d["high"].(float64); ok {
				return val, nil
			}
			if val, ok := d["open"].(float64); ok {
				return val, nil
			}
			if val, ok := d["close"].(float64); ok {
				return val, nil
			}
			if val, ok := d["low"].(float64); ok {
				return val, nil
			}
		}
	}
	return getUsdToRubRateFromCbr()
}
