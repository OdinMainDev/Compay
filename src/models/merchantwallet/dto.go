package merchantwallet

import (
	"efw/core/efwamount"

	"github.com/gin-gonic/gin"
)

func DTO(model *Model) any {
	return gin.H{
		"id":              model.Id,
		"merchant_id":     model.MerchantId,
		"payment_id":      model.PaymentId,
		"currency":        model.Currency,
		"address":         model.Address,
		"amount":          efwamount.AmountToFloatByCodeAlpha(model.Amount, model.Currency),
		"received_amount": efwamount.AmountToFloatByCodeAlpha(model.ReceivedAmount, model.Currency),
		"sent_amount":     efwamount.AmountToFloatByCodeAlpha(model.SentAmount, model.Currency),
		"created_at":      model.CreatedAt,
		"updated_at":      model.UpdatedAt,
	}
}
