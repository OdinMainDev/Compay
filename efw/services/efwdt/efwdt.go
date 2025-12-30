package efwdt

import (
	"gorm.io/gorm"
)

type ChainFilter[T any] func(gorm.ChainInterface[T]) gorm.ChainInterface[T]

type FilterType string
type Filter struct {
	Name   string     `json:"name"`
	Type   FilterType `json:"type"`
	Values any        `json:"values"`
}

type SpecialFilters struct {
	Query string
	Args  []any
}

type SpecialFilterHandler func(fieldName string, value any) (SpecialFilters, bool)
type SpecialFilterMap map[string]SpecialFilterHandler
