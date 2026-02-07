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

func DTOMerchant(model *Model) gin.H { return dto(model, payment.DTOMerchant, entity.DTOMerchant) }
func DTOAdmin(model *Model) gin.H    { return dto(model, payment.DTOAdmin, entity.DTOAdmin) }

func dto(model *Model, paymentDtofn func(*payment.Entity) gin.H, dtoFn func(*entity.Entity, gin.H) gin.H) gin.H {
	paymentDTO := gin.H{}
	if model.Payment != nil {
		paymentDTO = paymentDtofn(model.Payment)
	}
	return dtoFn(&model.Entity, paymentDTO)
}
