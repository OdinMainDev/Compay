package efwamount

import (
	"efw/components/entity"
	"maps"
	"math"

	"github.com/shopspring/decimal"
)

func GetMinorUnits(codeAlpha entity.CurrencyCodeAlpha) entity.CurrencyMinorUnits {
	mu.RLock()
	defer mu.RUnlock()
	return minorUnitsMap[codeAlpha]
}

func GetMinorUnitsMap() map[entity.CurrencyCodeAlpha]entity.CurrencyMinorUnits {
	mu.RLock()
	defer mu.RUnlock()
	result := make(map[entity.CurrencyCodeAlpha]entity.CurrencyMinorUnits, len(minorUnitsMap))
	maps.Copy(result, minorUnitsMap)
	return result
}

func ToFloat(amount entity.Amount, units entity.CurrencyMinorUnits) float64 {
	return toFloat(amount, units)
}

func ToFloatByCodeAlpha(amount entity.Amount, codeAlpha entity.CurrencyCodeAlpha) float64 {
	return toFloat(amount, GetMinorUnits(codeAlpha))
}

func toFloat(amount entity.Amount, units entity.CurrencyMinorUnits) float64 {
	if amount == 0 {
		return 0
	}
	pow := decimal.NewFromInt32(int32(units))
	dec, _ := decimal.NewFromInt(int64(amount)).Div(decimal.NewFromInt(10).Pow(pow)).Float64()
	return dec
}

func ToAmount(amount float64, units entity.CurrencyMinorUnits) entity.Amount {
	return toAmount(amount, units)
}

func ToAmountByCodeAlpha(amount float64, codeAlpha entity.CurrencyCodeAlpha) entity.Amount {
	return toAmount(amount, GetMinorUnits(codeAlpha))
}

func toAmount(amount float64, units entity.CurrencyMinorUnits) entity.Amount {
	if amount == 0 {
		return 0
	}

	if math.IsInf(amount, 0) {
		amount = toFixed(amount, int(units))
	}
	pow := decimal.NewFromInt32(int32(units))
	return entity.Amount(decimal.NewFromFloat(amount).Mul(decimal.NewFromInt(10).Pow(pow)).IntPart())
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Round(num*output) / output
}
