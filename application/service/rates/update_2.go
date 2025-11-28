package rates

import (
	"compay/core/eslog"
	"fmt"
)

func UpdateRatesBestFirst() (map[string]map[string]float64, bool, error) {
	var usedBest bool

	kucoinRates, err := getKucoin()
	if err != nil {
		return nil, false, err
	}

	rubRate, err := getUsdToRubRateFromBestChange()
	if err != nil {
		eslog.LogError("bestchange: ", err.Error())
		rubRate, err = getUsdToRubRateFromCbr()
		if err != nil {
			eslog.LogError(err.Error())
			return nil, false, err
		}
		usedBest = false
	} else {
		usedBest = true
	}

	eslog.Log("IS BESTCHANGE PARSED ", usedBest)

	if rubRate == 0 {
		return nil, false, fmt.Errorf("%s", "badRubRate")
	}

	eurRate, err := getUsdToEurRate(kucoinRates)
	if err != nil {
		return nil, false, err
	}
	cnyRate, err := getUsdToCnyRateFromChCbr()
	if err != nil {
		return nil, false, err
	}

	response := kucoinRates
	response = addFiatRates(response, rubRate, eurRate, cnyRate)
	response = completeCurrencyPairs(response, rubRate, eurRate)
	for _, currency := range []string{"RUB", "EUR", "CNY"} {
		response = addOppositeConversion(response, currency)
	}
	response = fillMissingRates(response)

	return response, usedBest, nil
}
