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
		"min_amount":                efwamount.ToFloatByCodeAlpha(model.MinAmount, model.Currency),
		"max_amount":                efwamount.ToFloatByCodeAlpha(model.MaxAmount, model.Currency),
		"min_amount_withdraw":       efwamount.ToFloatByCodeAlpha(model.MinAmountWithdraw, model.Currency),
		"max_amount_withdraw":       efwamount.ToFloatByCodeAlpha(model.MaxAmountWithdraw, model.Currency),
		"confirms":                  model.Confirms,
		"invoice_time":              model.InvoiceTime,
		"estimated_fee_block_count": model.EstimatedFeeBlockCount,
		"fee_type":                  model.FeeType,
		"fee_fixed":                 efwamount.ToFloatByCodeAlpha(entity.Amount(model.FeeFixed), model.Currency),
		"fee_percent":               efwamount.ToFloat(entity.Amount(model.FeePercent), 2),
		"created_at":                model.CreatedAt,
		"updated_at":                model.UpdatedAt,
	}
}
