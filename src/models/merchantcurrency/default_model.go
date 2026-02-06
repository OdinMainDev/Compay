package merchantcurrency

import (
	"compay/common/constants"
	"compay/common/types"
	"efw/components/entity"
	"efw/core/efwlog"
)

func DefaultCurrency(codeAlpha entity.CurrencyCodeAlpha, merchantId types.MerchantId, cretorId entity.UserId) *Model {
	switch codeAlpha {
	case constants.BtcCodeAlpha:
		return defaultBtc(merchantId, cretorId)

	default:
		efwlog.Panic("BAD DEFAULT CURRENCY: ", codeAlpha)
	}
	return nil
}

func defaultBtc(merchantId types.MerchantId, cretorId entity.UserId) *Model {
	model := &Model{}
	model.MerchantId = merchantId
	model.Currency = constants.BtcCodeAlpha
	model.CreatorId = cretorId
	model.MinAmount = constants.BtcMinAmount
	model.MaxAmount = constants.BtcMaxAmount
	model.MinAmountWithdraw = constants.BtcMinAmountWithdraw
	model.MaxAmountWithdraw = constants.BtcMaxAmountWithdraw
	model.Confirms = constants.BtcConfirmation
	model.InvoiceTime = constants.BtcInvoiceTime
	model.FeeType = constants.BtcFeeType
	model.FeeFixed = constants.BtcFeeFixed
	model.FeePercent = constants.BtcFeePercent
	return model
}
