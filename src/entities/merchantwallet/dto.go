package merchantwallet

import (
	"efw/core/efwamount"

	"github.com/gin-gonic/gin"
)

func DTOAdmin(model *Entity, paymentDTO gin.H) gin.H {
	return gin.H{
		"id":              model.Id,
		"merchant_id":     model.MerchantId,
		"payment_id":      model.PaymentId,
		"currency":        model.Currency,
		"address":         model.Address,
		"amount":          efwamount.ToFloatByCodeAlpha(model.Amount, model.Currency),
		"received_amount": efwamount.ToFloatByCodeAlpha(model.ReceivedAmount, model.Currency),
		"sent_amount":     efwamount.ToFloatByCodeAlpha(model.SentAmount, model.Currency),
		"payment":         paymentDTO,
		"created_at":      model.CreatedAt,
		"updated_at":      model.UpdatedAt,
	}
}

func DTOMerchant(model *Entity, paymentDTO gin.H) gin.H {
	return gin.H{
		"id":              model.Id,
		"currency":        model.Currency,
		"address":         model.Address,
		"amount":          efwamount.ToFloatByCodeAlpha(model.Amount, model.Currency),
		"received_amount": efwamount.ToFloatByCodeAlpha(model.ReceivedAmount, model.Currency),
		"sent_amount":     efwamount.ToFloatByCodeAlpha(model.SentAmount, model.Currency),
		"payment":         paymentDTO,
		"created_at":      model.CreatedAt,
		"updated_at":      model.UpdatedAt,
	}
}
