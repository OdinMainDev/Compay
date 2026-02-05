package payment

import (
	"compay/src/models/currency"

	"github.com/gin-gonic/gin"
)

func DTO(model *Model) any {
	return gin.H{
		"id":         model.TrxId,
		"currency":   model.Currency,
		"status":     model.Status,
		"address":    model.Address,
		"amount":     currency.AmountToFloatByCodeAlpha(model.Amount, model.Currency),
		"amount_usd": currency.AmountToFloat(model.AmountUsd, 2),
		"expires_at": model.ExpiresAt,
	}
}

func DTOMerchant(model *Model) any {
	return gin.H{
		"id":                     model.TrxId,
		"currency":               model.Currency,
		"status":                 model.Status,
		"address":                model.Address,
		"amount":                 currency.AmountToFloatByCodeAlpha(model.Amount, model.Currency),
		"amount_usd":             currency.AmountToFloat(model.AmountUsd, 2),
		"received_amount":        currency.AmountToFloatByCodeAlpha(model.ReceivedAmount, model.Currency),
		"received_amount_usd":    currency.AmountToFloat(model.ReceivedAmountUsd, 2),
		"required_confirmations": model.RequiredConfirmations,
		"expires_at":             model.ExpiresAt,
	}
}

func DTOAdmin(model *Model) any {
	return gin.H{
		"id":                     model.Id,
		"trx_id":                 model.TrxId,
		"merchant_id":            model.MerchantId,
		"currency":               model.Currency,
		"status":                 model.Status,
		"address":                model.Address,
		"amount":                 currency.AmountToFloatByCodeAlpha(model.Amount, model.Currency),
		"amount_usd":             currency.AmountToFloat(model.AmountUsd, 2),
		"received_amount":        currency.AmountToFloatByCodeAlpha(model.ReceivedAmount, model.Currency),
		"received_amount_usd":    currency.AmountToFloat(model.ReceivedAmountUsd, 2),
		"required_confirmations": model.RequiredConfirmations,
		"expires_at":             model.ExpiresAt,
		"created_at":             model.CreatedAt,
		"updated_at":             model.UpdatedAt,
	}
}
