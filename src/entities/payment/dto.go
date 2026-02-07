package payment

import (
	"efw/core/efwamount"

	"github.com/gin-gonic/gin"
)

func DTO(model *Entity) gin.H {
	return gin.H{
		"id":         model.TrxId,
		"currency":   model.Currency,
		"status":     model.Status,
		"address":    model.Address,
		"amount":     efwamount.ToFloatByCodeAlpha(model.Amount, model.Currency),
		"amount_usd": efwamount.ToFloat(model.AmountUsd, 2),
		"expires_at": model.ExpiresAt,
	}
}

func DTOMerchant(model *Entity) gin.H {
	return gin.H{
		"id":                     model.TrxId,
		"currency":               model.Currency,
		"status":                 model.Status,
		"address":                model.Address,
		"amount":                 efwamount.ToFloatByCodeAlpha(model.Amount, model.Currency),
		"amount_usd":             efwamount.ToFloat(model.AmountUsd, 2),
		"received_amount":        efwamount.ToFloatByCodeAlpha(model.ReceivedAmount, model.Currency),
		"received_amount_usd":    efwamount.ToFloat(model.ReceivedAmountUsd, 2),
		"required_confirmations": model.RequiredConfirmations,
		"expires_at":             model.ExpiresAt,
	}
}

func DTOAdmin(model *Entity) gin.H {
	return gin.H{
		"id":                     model.Id,
		"trx_id":                 model.TrxId,
		"merchant_id":            model.MerchantId,
		"currency":               model.Currency,
		"status":                 model.Status,
		"address":                model.Address,
		"amount":                 efwamount.ToFloatByCodeAlpha(model.Amount, model.Currency),
		"amount_usd":             efwamount.ToFloat(model.AmountUsd, 2),
		"received_amount":        efwamount.ToFloatByCodeAlpha(model.ReceivedAmount, model.Currency),
		"received_amount_usd":    efwamount.ToFloat(model.ReceivedAmountUsd, 2),
		"required_confirmations": model.RequiredConfirmations,
		"expires_at":             model.ExpiresAt,
		"created_at":             model.CreatedAt,
		"updated_at":             model.UpdatedAt,
	}
}
