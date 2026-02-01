package currency

import (
	"compay/common/types"
	"efw/utils"
	"math"

	"github.com/shopspring/decimal"
)

func GetMinorUnits(codeAlpha types.CurrencyCodeAlpha) types.CurrencyMinorUnits {
	currency := GetModel(QueryHelper{CodeAlpha: codeAlpha})
	if currency == nil || currency.GetId() == 0 {
		return 0
	}
	return currency.MinorUnits
}

func AmountToFloat(amount types.Amount, units types.CurrencyMinorUnits) float64 {
	return amountToFloat(amount, units)
}

func AmountToFloatByCodeAlpha(amount types.Amount, codeAlpha types.CurrencyCodeAlpha) float64 {
	return amountToFloat(amount, GetMinorUnits(codeAlpha))
}

func amountToFloat(amount types.Amount, units types.CurrencyMinorUnits) float64 {
	pow := decimal.NewFromInt32(int32(units))
	dec, _ := decimal.NewFromInt(int64(amount)).Div(decimal.NewFromInt(10).Pow(pow)).Float64()
	return dec
}

func FloatToAmount(amount float64, units types.CurrencyMinorUnits) types.Amount {
	return floatToAmount(amount, units)
}

func floatToAmount(amount float64, units types.CurrencyMinorUnits) types.Amount {
	if math.IsInf(amount, 0) {
		amount = utils.ToFixed(amount, int(units))
	}
	pow := decimal.NewFromInt32(int32(units))
	return types.Amount(decimal.NewFromFloat(amount).Mul(decimal.NewFromInt(10).Pow(pow)).IntPart())
}

// func AmountToInt(amount float64, currency types.CurrencyCodeAlpha) types.Amount {
// 	units := GetMinorUnits(currency)
// 	return types.Amount(amountToInt(amount, units))
// }

// func CalculatePercentageOfAmount(amount int64, _ string, percent int) int64 {
// 	percentDec := decimal.NewFromInt(int64(percent)).Div(decimal.NewFromInt(100))
// 	result := decimal.NewFromInt(amount).Mul(percentDec)
// 	result = result.Round(0)
// 	return result.IntPart()
// }

// func AmountToInt2Digit(amount float64) int64 {
// 	return amountToInt(amount, 2)
// }
