package startup

import (
	"compay/src/models/currency"
	"efw/core/efwamount"
)

func FillEfwAmount() {
	currencyModels := currency.GetModels(currency.QueryHelper{})
	for _, currencyModel := range currencyModels {
		efwamount.SetMinorUnits(currencyModel.CodeAlpha, currencyModel.MinorUnits)
	}
}
