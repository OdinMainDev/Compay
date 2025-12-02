package ratesparser

func addEurRate(newRates map[string]map[string]float64) error {
	if v, ok := newRates["USDT"]["EUR"]; ok {
		newRates["USD"]["EUR"] = v
		return nil
	}

	if v, ok := newRates["USD"]["EUR"]; ok {
		newRates["USDT"]["EUR"] = v
		return nil
	}

	var err error
	for range 3 {
		v, err2 := getExchangeRateRate("EUR")
		if err2 == nil {
			newRates["USD"]["EUR"] = v
			newRates["USDT"]["EUR"] = v
			return nil
		}
		err = err2
	}

	return err
}
