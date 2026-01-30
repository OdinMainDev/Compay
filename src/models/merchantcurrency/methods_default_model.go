package merchantcurrency

import (
	"compay/common/constants"
	"compay/common/types"
	"efw/core/efwlog"
)

func DefaultCurrency(codeAlpha types.CurrencyCodeAlpha) *Model {
	switch codeAlpha {
	case constants.BtcCodeAlpha:
		return defaultBtc()
	}

	efwlog.Panic("BAD DEFAULT CURRENCY: ", codeAlpha)
	return nil
}

func defaultBtc() *Model {
	model := &Model{}
	model.Currency = constants.BtcCodeAlpha
	model.MinAmount = constants.BtcMinAmount
	model.MaxAmount = constants.BtcMaxAmount
	model.MinAmountWithdraw = constants.BtcMinAmountWithdraw
	model.MaxAmountWithdraw = constants.BtcMaxAmountWithdraw
	model.MinConfirms = constants.BtcConfirmation
	model.FeeType = constants.BtcFeeType
	model.FeeFixed = constants.BtcFeeFixed
	model.FeePercent = constants.BtcFeePercent
	return model
}
