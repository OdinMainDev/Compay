package merchantwallet

import (
	entity "compay/src/entities/merchantwallet"
	"compay/src/entities/payment"

	"github.com/gin-gonic/gin"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
	Payment       *payment.Entity `json:"payment,omitempty" gorm:"-,foreignKey:PaymentId;references:Id"`
}

func DTO(model *Model) gin.H         { return dto(model, payment.DTO) }
func DTOMerchant(model *Model) gin.H { return dto(model, payment.DTOMerchant) }
func DTOAdmin(model *Model) gin.H    { return dto(model, payment.DTOAdmin) }

func dto(model *Model, fn func(*payment.Entity) gin.H) gin.H {
	paymentDTO := gin.H{}
	if model.Payment != nil {
		paymentDTO = fn(model.Payment)
	}
	return entity.DTO(&model.Entity, paymentDTO)
}
