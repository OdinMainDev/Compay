package logpayment

import (
	entity "compay/src/entities/logpayment"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}
