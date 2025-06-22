package rates

import (
	"log"
	"sync"
	"time"
)

var (
	currencies = []string{"USD", "USDT", "RUB", "BTC", "ETH", "DAI", "ADA", "XRP", "LTC", "LINK", "BCH", "XLM", "TRX", "DOGE", "SOL", "MATIC", "XMR", "EUR", "BNB"}
	parserCfg  = map[string]string{
		"kucoin": "https://api.kucoin.com/api/v1/market/allTickers",
		"commex": "https://api.commex.com/api/v1/ticker/bookTicker?symbol=USDTRUB",
		"rapira": "https://api.rapira.net/market/symbol-thumb",
		"cbr":    "https://www.cbr.ru/scripts/XML_daily.asp",
	}

	rates   map[string]map[string]float64
	ratesMu sync.RWMutex
)

func Initialize() {
	ticker := time.NewTicker(30 * time.Second)

	for {
		<-ticker.C
		data, err := UpdateRates()
		if err != nil {
			log.Println("Failed to update rates:", err)
			return
		}
		ratesMu.Lock()
		rates = data
		ratesMu.Unlock()
	}
}
