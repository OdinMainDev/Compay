package merchantcurrency

import (
	"compay/common/types"
	"compay/src/models/currency"

	"github.com/gin-gonic/gin"
)

func DTO(model *Model) any {
	return gin.H{
		"merchant_id":         model.MerchantId,
		"currency":            model.Currency,
		"creator_id":          model.CreatorId,
		"min_amount":          currency.AmountToFloatByCodeAlpha(model.MinAmount, model.Currency),
		"max_amount":          currency.AmountToFloatByCodeAlpha(model.MaxAmount, model.Currency),
		"min_amount_withdraw": currency.AmountToFloatByCodeAlpha(model.MinAmountWithdraw, model.Currency),
		"max_amount_withdraw": currency.AmountToFloatByCodeAlpha(model.MaxAmountWithdraw, model.Currency),
		"confirms":            model.Confirms,
		"invoice_time":        model.InvoiceTime,
		"fee_type":            model.FeeType,
		"fee_fixed":           currency.AmountToFloatByCodeAlpha(types.Amount(model.FeeFixed), model.Currency),
		"fee_percent":         currency.AmountToFloat(types.Amount(model.FeePercent), 2),
		"created_at":          model.CreatedAt,
		"updated_at":          model.UpdatedAt,
	}
}
