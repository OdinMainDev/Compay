package helpers

import (
	"compay/core/utils"
	"math"

	"github.com/shopspring/decimal"
)

func AmountToFloat(amount int64, units int) float64 {
	pow := decimal.NewFromInt32(int32(units))
	dec, _ := decimal.NewFromInt(amount).Div(decimal.NewFromInt(10).Pow(pow)).Float64()
	return dec
}

func AmountToInt(amount float64, units int) int64 {
	return amountToInt(amount, units)
}

func CalculatePercentageOfAmount(amount int64, units int, percent float64) int64 {
	return AmountToInt((AmountToFloat(amount, units)/100)*percent, units)
}

func AmountToInt2Digit(amount float64) int64 {
	return amountToInt(amount, 2)
}

func amountToInt(amount float64, units int) int64 {
	if math.IsInf(amount, 0) {
		amount = utils.ToFixed(amount, units)
	}
	pow := decimal.NewFromInt32(int32(units))
	return decimal.NewFromFloat(amount).Mul(decimal.NewFromInt(10).Pow(pow)).IntPart()
}
