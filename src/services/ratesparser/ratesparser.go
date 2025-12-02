package ratesparser

import (
	"efw/core/efwnotify"
	"sync"
	"time"
)

var (
	currencies = []string{"USDT", "RUB", "BTC", "ETH", "DAI", "ADA", "XRP", "LTC", "LINK", "BCH", "XLM", "TRX", "DOGE", "SOL", "XMR", "EUR", "BNB", "TON"}
	parserCfg  = map[string]string{
		"okx":          "https://www.okx.com/api/v5/market/index-tickers?quoteCcy=USDT",
		"kucoin":       "https://api.kucoin.com/api/v1/market/allTickers",
		"rapira":       "https://api.rapira.net/open/market/rates",
		"exchangerate": "https://api.exchangerate.fun/latest",
	}

	rates   map[string]map[string]float64
	ratesMu sync.RWMutex
)

const sleepTime = time.Duration(30)

func Initialize() {
	efwnotify.Notify("RATES INIT", efwnotify.Options{})
	for {
		data, err := UpdateRates()

		if err != nil {
			// отправка в телегу всякой хуйни
			efwnotify.Notify("RATES ERROR: "+err.Error(), efwnotify.Options{})
			time.Sleep(sleepTime * time.Second)
			continue
		}

		ratesMu.Lock()
		rates = data
		ratesMu.Unlock()

		efwnotify.Notify("RATES PARSE SUCCESS", efwnotify.Options{})
		time.Sleep(sleepTime * time.Second)
	}
}

func UpdateRates() (map[string]map[string]float64, error) {
	newRates := constructRates()

	// получаем от okx первую часть курсов
	var err error
	for range 3 {
		err = addOkxRates(newRates)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}

	// return newRates, nil

	// получаем от kukoin вторую часть курсов
	for range 3 {
		err = addKucoinRates(newRates)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}

	// получаем от rapira рублевое значение
	for range 3 {
		err = addRapiraRubRate(newRates)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	for range 3 {
		err = addEurRate(newRates)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}

	for range 3 {
		err = addCnyRate(newRates)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}

	addFiatRates(newRates)
	completeCurrencyPairs(newRates)
	for _, currency := range []string{"RUB", "EUR", "CNY"} {
		addOppositeConversion(newRates, currency)
	}
	fillMissingRates(newRates)

	newRates["USD"] = newRates["USDT"]
	return newRates, nil
}
