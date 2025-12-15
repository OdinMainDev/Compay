package merchant

import (
	"compay/common/types"
	"database/sql/driver"
	"efw/components/entity"
)

type Entity struct {
	entity.BaseEntity[types.MerchantId] `gorm:"embedded"`
	Name                                string   `json:"name" gorm:"column:name"`
	Token                               string   `json:"token" gorm:"column:token"`
	IsEnabled                           bool     `json:"is_enabled" gorm:"column:is_enabled"`
	Settings                            Settings `json:"settings" gorm:"column:settings"`
}

type Settings struct {
	BlockedCountries []string `json:"blocked_countries"`
	AllowedCountries []string `json:"allowed_countries"`
}

func (*Entity) TableName() string                  { return "merchants" }
func (data *Settings) Scan(value any) error        { return entity.ScanStruct(data, value) }
func (data Settings) Value() (driver.Value, error) { return entity.ValueStruct(data) }
