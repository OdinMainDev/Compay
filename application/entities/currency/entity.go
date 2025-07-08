package currency

import "compay/application/entities/basemodel"

type Model struct {
	basemodel.BaseModel
	CodeAlpha  string `json:"code_alpha" gorm:"column:code_alpha"`
	Title      string `json:"title" gorm:"column:title"`
	Network    string `json:"network" gorm:"column:network"`
	Image      string `json:"image" gorm:"column:image"`
	IsCoin     bool   `json:"is_coin" gorm:"column:is_coin"`
	MinorUnits int    `json:"minor_units" gorm:"column:minor_units"`
}

func (Model) TableName() string {
	return "currencies"
}
