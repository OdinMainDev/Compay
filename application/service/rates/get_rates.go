package rates

import "maps"

func GetRates() map[string]map[string]float64 {
	ratesMu.RLock()
	defer ratesMu.RUnlock()
	result := make(map[string]map[string]float64)
	for k, v := range rates {
		result[k] = make(map[string]float64)
		maps.Copy(result[k], v)
	}
	return result
}
