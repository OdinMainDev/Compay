package merchantwallet

import (
	entity "compay/src/entities/merchantwallet"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}
