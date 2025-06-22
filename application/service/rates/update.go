package rates

func UpdateRates() (map[string]map[string]float64, error) {
	var err error
	kucoinRates, err := getKucoin()
	if err != nil {
		return nil, err
	}
	rubRate, err := getUsdToRubRateFromRapira()
	if err != nil {
		return nil, err
	}
	eurRate, err := getUsdToEurRate(kucoinRates)
	if err != nil {
		return nil, err
	}
	cnyRate, err := getUsdToCnyRateFromChCbr()
	if err != nil {
		return nil, err
	}

	response := kucoinRates
	response = addFiatRates(response, rubRate, eurRate, cnyRate)
	response = completeCurrencyPairs(response, rubRate, eurRate)
	for _, currency := range []string{"RUB", "EUR", "CNY"} {
		response = addOppositeConversion(response, currency)
	}
	response = fillMissingRates(response)

	// ФУНКЦИИ ДЛЯ ТЕСТОВ
	// checkValues(response) // ТЕСТ ФУНКЦИЯ ЧЕРЕЗ КРИПТОКОМПЕЙР

	// v, _ := json.Marshal(response)
	// fmt.Println(string(v))

	return response, nil
}

func addFiatRates(response map[string]map[string]float64, rubRate, eurRate, cnyRate float64) map[string]map[string]float64 {
	for base, pairs := range response {
		if v, ok := pairs["USDT"]; ok {
			if _, ok2 := response[base]["RUB"]; !ok2 {
				response[base]["RUB"] = v * rubRate
			}
			if _, ok2 := response[base]["EUR"]; !ok2 {
				response[base]["EUR"] = v * eurRate
			}
			if _, ok2 := response[base]["CNY"]; !ok2 {
				response[base]["CNY"] = v * cnyRate
			}
		}
	}

	if _, ok := response["CNY"]; !ok {
		response["CNY"] = make(map[string]float64)
	}
	response["CNY"]["CNY"] = 1
	response["CNY"]["USD"] = 1 / cnyRate
	response["USD"]["CNY"] = cnyRate
	response["USDT"]["CNY"] = cnyRate

	return response
}

func completeCurrencyPairs(response map[string]map[string]float64, rubRate, eurRate float64) map[string]map[string]float64 {
	for _, currency := range currencies {
		if _, ok := response[currency]; !ok {
			continue
		}
		if _, ok := response[currency]["USDT"]; !ok {
			continue
		}
		for key, val := range response[currency] {
			if _, ok := response[key]["USDT"]; !ok {
				continue
			}
			response[currency]["USDT"] = val * response[key]["USDT"]
			break
		}
	}

	response["USDT"]["RUB"] = rubRate
	response["USDT"]["EUR"] = eurRate
	for base, pairs := range response {
		for _, targetCurrency := range currencies {
			if _, ok := pairs[targetCurrency]; ok {
				continue
			}
			if _, ok := pairs["USDT"]; !ok {
				continue
			}
			if _, ok := response[targetCurrency]["USDT"]; !ok {
				continue
			}
			baseToUsdt := pairs["USDT"]
			targetToUsdt := response[targetCurrency]["USDT"]
			response[base][targetCurrency] = baseToUsdt / targetToUsdt
		}
	}
	return response
}

func addOppositeConversion(response map[string]map[string]float64, baseCurrency string) map[string]map[string]float64 {
	oppositeConversions := make(map[string]float64)

	for currency, pairs := range response {
		if rate, ok := pairs[baseCurrency]; ok && rate != 0 {
			oppositeConversions[currency] = 1 / rate
		}
	}

	oppositeConversions[baseCurrency] = 1

	response[baseCurrency] = oppositeConversions
	return response
}

func fillMissingRates(data map[string]map[string]float64, bridgeCurrencies ...string) map[string]map[string]float64 {
	bridgeCurrency := "USD"
	if len(bridgeCurrencies) != 0 {
		bridgeCurrency = bridgeCurrencies[0]
	}
	for from, fromRates := range data {
		for to := range data {
			if from == to {
				continue
			}

			// Пропускаем, если курс уже есть
			if _, exists := fromRates[to]; exists {
				continue
			}

			// Проверяем, есть ли путь через bridge
			fromToBridge, ok1 := data[from][bridgeCurrency]
			toToBridge, ok2 := data[to][bridgeCurrency]

			if ok1 && ok2 && toToBridge != 0 {
				// Добавляем недостающий курс
				if data[from] == nil {
					data[from] = make(map[string]float64)
				}
				data[from][to] = fromToBridge / toToBridge

				// И обратный курс
				if data[to] == nil {
					data[to] = make(map[string]float64)
				}
				data[to][from] = toToBridge / fromToBridge
			}
		}
	}

	return data
}
