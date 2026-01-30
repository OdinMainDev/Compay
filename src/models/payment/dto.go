package payment

import (
	"github.com/gin-gonic/gin"
)

func DTO(model *Model) gin.H {
	return gin.H{
		"id":         model.TrxId,
		"currency":   model.Currency,
		"address":    model.Address,
		"amount":     model.Amount,
		"amount_usd": model.AmountUsd,
		"expires_at": model.ExpiresAt,
	}
}
