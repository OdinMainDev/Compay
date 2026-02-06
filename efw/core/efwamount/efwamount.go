package efwamount

import (
	"efw/components/entity"
	"sync"
)

var (
	minorUnitsMap = make(map[entity.CurrencyCodeAlpha]entity.CurrencyMinorUnits)
	mu            sync.RWMutex
)

func SetMinorUnits(codeAlpha entity.CurrencyCodeAlpha, units entity.CurrencyMinorUnits) {
	mu.Lock()
	defer mu.Unlock()
	minorUnitsMap[codeAlpha] = units
}

func SetMinorUnitsMap(m map[entity.CurrencyCodeAlpha]entity.CurrencyMinorUnits) {
	mu.Lock()
	defer mu.Unlock()
	minorUnitsMap = m
}
