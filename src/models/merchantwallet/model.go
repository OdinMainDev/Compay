package merchantwallet

import (
	entity "compay/src/entities/merchantwallet"
	"compay/src/entities/payment"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
	Payment       *payment.Entity `json:"payment,omitempty" gorm:"-,foreignKey:PaymentId;references:Id"`
}
