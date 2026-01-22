package user

import (
	"compay/src/entities/merchant"
	entity "compay/src/entities/user"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
	NewPassword   string           `json:"-" gorm:"-"`
	Merchant      *merchant.Entity `json:"merchant,omitempty" gorm:"foreignKey:MerchantId;references:Id"`
}
