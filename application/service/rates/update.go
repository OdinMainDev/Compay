package rates

func UpdateRates() (map[string]map[string]float64, error) {
	var err error
	kucoinRates, err := getKucoin()
	if err != nil {
		return nil, err
	}
	rubRate, err := getRapira()
	if err != nil {
		return nil, err
	}
	eurRate, err := getCbrEUR()
	if err != nil {
		return nil, err
	}

	for base, pairs := range kucoinRates {
		if v, ok := pairs["USDT"]; ok {
			kucoinRates[base]["RUB"] = v * rubRate
			kucoinRates[base]["EUR"] = v * eurRate
		}
	}

	rubMap := map[string]float64{}
	eurMap := map[string]float64{}
	for base, pairs := range kucoinRates {
		if val, ok := pairs["RUB"]; ok {
			rubMap[base] = 1 / val
		}
		if val, ok := pairs["EUR"]; ok {
			eurMap[base] = 1 / val
		}
	}
	rubMap["RUB"] = 1
	eurMap["EUR"] = 1
	kucoinRates["RUB"] = rubMap
	kucoinRates["EUR"] = eurMap

	return kucoinRates, nil
}
