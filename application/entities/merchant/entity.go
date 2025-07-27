package merchant

import (
	"compay/application/entities/basemodel"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Model struct {
	basemodel.BaseModel
	Name      string         `json:"name" gorm:"column:name"`
	Token     string         `json:"token" gorm:"column:token;index:"`
	IsEnabled bool           `json:"is_enabled" gorm:"column:is_enabled"`
	Settings  Settings       `json:"settings" gorm:"column:settings;type:jsonb"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Settings struct {
	ProfitPercent   float64  `json:"profit_percent"`
	AvailCurrencies []string `json:"avail_currencies"`
}

func (Model) TableName() string {
	return "merchants"
}

func (settings *Settings) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	err := json.Unmarshal(bytes, settings)
	return err
}

func (settings Settings) Value() (driver.Value, error) {
	jsonBytes, err := json.Marshal(settings)
	return string(jsonBytes), err
}
