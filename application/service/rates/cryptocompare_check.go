package rates

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

// ФУНКЦИЯ ИСПОЛЬЗУЕТСЯ В ТЕСТАХ
func CheckValues(response map[string]map[string]float64) {
	for currencyFrom, pairs := range response {
		for currencyTo, value := range pairs {
			if currencyTo == currencyFrom {
				continue
			}

			if currencyTo != "CNY" && currencyFrom != "CNY" {
				continue
			}
			time.Sleep(time.Millisecond * 100)
			check, _, val2, err := CompareRate(currencyFrom, currencyTo, value)
			fmt.Println("---------------------------")
			fmt.Println("CHECK PAIR: ", currencyFrom, currencyTo)
			if err != nil {
				fmt.Println("HAS ERROR: ", err.Error())
			}
			fmt.Println("CURRENT VALUE: ", value)
			fmt.Println("COMPARE VALUE: ", val2)
			fmt.Println("IS GOOD: ", check)

			if !check || val2 == 0 {
				continue
			}
		}
	}
}

func CompareRate(baseCurrency, targetCurrency string, givenRate float64) (bool, float64, float64, error) {
	apiKey := "c8a318211fea04ed9627e5e437c98a7d00d75bb2137f98229b1cb01af2258f05"

	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemulti?api_key=%s&fsyms=%s&tsyms=%s",
		apiKey, baseCurrency, targetCurrency)

	resp, err := http.Get(url)
	if err != nil {
		return false, 0, 0, fmt.Errorf("http request error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, 0, 0, fmt.Errorf("read body error: %v", err)
	}

	var result map[string]map[string]float64
	if err := json.Unmarshal(body, &result); err != nil {
		return false, 0, 0, fmt.Errorf("json unmarshal error: %v", err)
	}

	rate, ok := result[baseCurrency][targetCurrency]
	if !ok {
		return false, 0, 0, fmt.Errorf("rate for %s/%s not found in response", baseCurrency, targetCurrency)
	}

	diffPercent := math.Abs(rate-givenRate) / givenRate * 100
	return diffPercent <= 5.0, diffPercent, rate, nil
}
