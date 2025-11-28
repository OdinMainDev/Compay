package rates

import (
	"encoding/json"
	"fmt"
)

func getUsdToRubRateFromRapira() (float64, error) {
	return 0, fmt.Errorf("notWorking")
	body, err := doRequest(parserCfg["rapira"])
	if err != nil {
		return 0, fmt.Errorf("rapira 1: %s", err.Error())
	}
	var data []map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, fmt.Errorf("rapira 2: %s | body: %s", err.Error(), string(body))
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

	return 0, fmt.Errorf("rapira 3: %s | body: %s", "not avail pairs", string(body))
}
