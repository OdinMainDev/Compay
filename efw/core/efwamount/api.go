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

func AmountToFloat(amount entity.Amount, units entity.CurrencyMinorUnits) float64 {
	return amountToFloat(amount, units)
}

func AmountToFloatByCodeAlpha(amount entity.Amount, codeAlpha entity.CurrencyCodeAlpha) float64 {
	return amountToFloat(amount, GetMinorUnits(codeAlpha))
}

func amountToFloat(amount entity.Amount, units entity.CurrencyMinorUnits) float64 {
	if amount == 0 {
		return 0
	}
	pow := decimal.NewFromInt32(int32(units))
	dec, _ := decimal.NewFromInt(int64(amount)).Div(decimal.NewFromInt(10).Pow(pow)).Float64()
	return dec
}

func FloatToAmount(amount float64, units entity.CurrencyMinorUnits) entity.Amount {
	return floatToAmount(amount, units)
}

func FloatToAmountByCodeAlpha(amount float64, codeAlpha entity.CurrencyCodeAlpha) entity.Amount {
	return floatToAmount(amount, GetMinorUnits(codeAlpha))
}

func floatToAmount(amount float64, units entity.CurrencyMinorUnits) entity.Amount {
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
