package ratesparser

import "efw/core/efwrequest"

func doRequest(url string) ([]byte, error) {
	resp, _, err := efwrequest.Request(efwrequest.RequestHelper{
		Url:    url,
		Method: "GET",
		Headers: map[string]string{
			"Accept-Language": "ru-RU,ru;q=0.8,en-US;q=0.6,en;q=0.4",
			"Cache-Control":   "no-cache",
			"Pragma":          "no-cache",
			"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.3",

			"sec-ch-ua":                 `"Chromium";v="137", "Not/A)Brand";v="24"`,
			"sec-ch-ua-mobile":          "?0",
			"sec-ch-ua-platform":        `"macOS"`,
			"sec-fetch-dest":            "document",
			"sec-fetch-mode":            "navigate",
			"sec-fetch-site":            "none",
			"sec-fetch-user":            "?1",
			"upgrade-insecure-requests": "1",
		},
		IsLogDisabled: true,
		// IsProxyEnable: true,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func constructRates() map[string]map[string]float64 {
	resp := make(map[string]map[string]float64)
	for _, currency := range currencies {
		resp[currency] = map[string]float64{}
		resp[currency][currency] = 1
	}
	return resp
}

func addFiatRates(response map[string]map[string]float64) {
	rubRate := response["USD"]["RUB"]
	eurRate := response["USD"]["EUR"]
	cnyRate := response["USD"]["CNY"]
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
}

func completeCurrencyPairs(response map[string]map[string]float64) {
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
}

func addOppositeConversion(response map[string]map[string]float64, baseCurrency string) {
	oppositeConversions := make(map[string]float64)

	for currency, pairs := range response {
		if rate, ok := pairs[baseCurrency]; ok && rate != 0 {
			oppositeConversions[currency] = 1 / rate
		}
	}

	oppositeConversions[baseCurrency] = 1

	response[baseCurrency] = oppositeConversions
}

func fillMissingRates(data map[string]map[string]float64, bridgeCurrencies ...string) {
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
}
