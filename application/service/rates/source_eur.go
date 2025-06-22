package rates

import (
	"errors"
)

func getUsdToEurRate(kucoinRates map[string]map[string]float64) (float64, error) {
	if v, ok := kucoinRates["USD"]["EUR"]; ok {
		return v, nil
	}

	return 0, errors.New("EUR not found")
}
