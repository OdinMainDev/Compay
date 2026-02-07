package merchantcurrency

import (
	"efw/components/entity"
	"efw/core/efwamount"

	"github.com/gin-gonic/gin"
)

func DTO(model *Entity) gin.H {
	return gin.H{
		"merchant_id":               model.MerchantId,
		"currency":                  model.Currency,
		"creator_id":                model.CreatorId,
		"min_amount":                efwamount.AmountToFloatByCodeAlpha(model.MinAmount, model.Currency),
		"max_amount":                efwamount.AmountToFloatByCodeAlpha(model.MaxAmount, model.Currency),
		"min_amount_withdraw":       efwamount.AmountToFloatByCodeAlpha(model.MinAmountWithdraw, model.Currency),
		"max_amount_withdraw":       efwamount.AmountToFloatByCodeAlpha(model.MaxAmountWithdraw, model.Currency),
		"confirms":                  model.Confirms,
		"invoice_time":              model.InvoiceTime,
		"estimated_fee_block_count": model.EstimatedFeeBlockCount,
		"fee_type":                  model.FeeType,
		"fee_fixed":                 efwamount.AmountToFloatByCodeAlpha(entity.Amount(model.FeeFixed), model.Currency),
		"fee_percent":               efwamount.AmountToFloat(entity.Amount(model.FeePercent), 2),
		"created_at":                model.CreatedAt,
		"updated_at":                model.UpdatedAt,
	}
}
