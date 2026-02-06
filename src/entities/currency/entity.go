package currency

import (
	"compay/common/types"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.CurrencyId] `gorm:"embedded"`
	CodeAlpha                           entity.CurrencyCodeAlpha  `json:"code_alpha" gorm:"column:code_alpha"`
	IsEnabled                           bool                      `json:"is_enabled" gorm:"column:is_enabled"`
	Title                               string                    `json:"title" gorm:"column:title"`
	Network                             types.CurrencyNetwork     `json:"network" gorm:"column:network"`
	Image                               string                    `json:"image" gorm:"column:image"`
	IsCoin                              types.CurrencyIsCoin      `json:"is_coin" gorm:"column:is_coin"`
	MinorUnits                          entity.CurrencyMinorUnits `json:"minor_units" gorm:"column:minor_units"`
}

func (*Entity) TableName() string { return "currencies" }
