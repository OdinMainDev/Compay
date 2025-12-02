package ratesparser

func addCnyRate(newRates map[string]map[string]float64) error {
	var err error
	for range 3 {
		v, err2 := getExchangeRateRate("CNY")
		if err2 == nil {
			newRates["USD"]["CNY"] = v
			newRates["USDT"]["CNY"] = v
			return nil
		}
		err = err2
	}

	return err
}
