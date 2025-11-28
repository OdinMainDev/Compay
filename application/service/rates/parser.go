package rates

import (
	"compay/core/eslog"
	"fmt"
	"sync"
	"time"
)

var (
	currencies = []string{"USD", "USDT", "RUB", "BTC", "ETH", "DAI", "ADA", "XRP", "LTC", "LINK", "BCH", "XLM", "TRX", "DOGE", "SOL", "MATIC", "XMR", "EUR", "BNB"}
	parserCfg  = map[string]string{
		"kucoin":     "https://api.kucoin.com/api/v1/market/allTickers",
		"commex":     "https://api.commex.com/api/v1/ticker/bookTicker?symbol=USDTRUB",
		"rapira":     "https://api.rapira.net/market/symbol-thumb",
		"cbr":        "https://www.cbr.ru/scripts/XML_daily.asp",
		"grinex":     "https://grinex.io/api/v2/k?market=usdta7a5",
		"bestchange": "http://api.bestchange.ru/info.zip",
	}

	rates   map[string]map[string]float64
	ratesMu sync.RWMutex
)

func Initialize() {
	isFirstTime := true
	for {
		data, usedBest, err := UpdateRatesBestFirst()
		if err != nil {
			eslog.LogError(fmt.Sprintf("Failed to update rates: %s", err.Error()))
			if isFirstTime {
				time.Sleep(15 * time.Second)
			} else {
				time.Sleep(120 * time.Second)
			}
			continue
		}
		isFirstTime = false
		ratesMu.Lock()
		rates = data
		ratesMu.Unlock()

		if usedBest {
			time.Sleep(240 * time.Second)
		} else {
			time.Sleep(60 * time.Second)
		}
	}
}
