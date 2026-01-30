package merchantcurrency

import (
	"compay/common/constants"
	"compay/common/types"
	"efw/core/efwlog"
)

func DefaultCurrency(codeAlpha types.CurrencyCodeAlpha) *Model {
	switch codeAlpha {
	case constants.CurrencyCodeAlphaBtc:
		return defaultBtc()
	}

	efwlog.Panic("BAD DEFAULT CURRENCY: ", codeAlpha)
	return nil
}

func defaultBtc() *Model {
	model := &Model{}
	model.Currency = constants.CurrencyCodeAlphaBtc
	model.MinAmount = constants.CurrencyMinAmountBtc
	model.MaxAmount = constants.CurrencyMaxAmountBtc
	model.MinConfirms = constants.CurrencyConfirmationBtc
	model.FeeType = constants.CurrencyFeeTypeBtc
	model.FeeFixed = constants.CurrencyFeeFixedBtc
	model.FeePercent = constants.CurrencyFeePercentBtc
	return model
}
