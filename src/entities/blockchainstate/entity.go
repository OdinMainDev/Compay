package blockchainstate

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.BlockchainStateId] `gorm:"embedded"`
	Currency                                   entity.CurrencyCodeAlpha `json:"currency" gorm:"column:currency"`
	CurrentBlock                               int64                    `json:"current_block" gorm:"column:current_block"`
}

func (*Entity) TableName() string { return "blockchain_states" }
