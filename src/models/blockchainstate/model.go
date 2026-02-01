package blockchainstate

import (
	entity "compay/src/entities/blockchainstate"
)

type Model struct {
	entity.Entity `gorm:"embedded"`
}
