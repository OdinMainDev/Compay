package withdraw

import (
	"efw/core/efwamount"

	"github.com/gin-gonic/gin"
)

func DTO(model *Entity) gin.H {
	return gin.H{
		"id":         model.TrxId,
		"currency":   model.Currency,
		"is_success": model.IsComplete && model.Error == "",
		"address":    model.Address,
		"amount":     efwamount.ToFloatByCodeAlpha(model.Amount, model.Currency),
		"amount_usd": efwamount.ToFloat(model.AmountUsd, 2),
	}
}

func DTOMerchant(model *Entity) gin.H {
	return gin.H{
		"id":          model.TrxId,
		"currency":    model.Currency,
		"is_complete": model.IsComplete,
		"is_success":  model.IsComplete && model.Error == "",
		"address":     model.Address,
		"amount":      efwamount.ToFloatByCodeAlpha(model.Amount, model.Currency),
		"amount_usd":  efwamount.ToFloat(model.AmountUsd, 2),
	}
}

func DTOAdmin(model *Entity) gin.H {
	return gin.H{
		"id":                 model.Id,
		"trx_id":             model.TrxId,
		"merchant_id":        model.MerchantId,
		"merchant_wallet_id": model.MerchantWalletId,
		"currency":           model.Currency,
		"is_complete":        model.IsComplete,
		"address":            model.Address,
		"amount":             efwamount.ToFloatByCodeAlpha(model.Amount, model.Currency),
		"amount_usd":         efwamount.ToFloat(model.AmountUsd, 2),
		"error":              model.Error,
		"created_at":         model.CreatedAt,
		"updated_at":         model.UpdatedAt,
	}
}
