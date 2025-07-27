package start

import (
	"compay/application/entities/currency"
	"compay/application/entities/merchant"

	"compay/core/escache"
)

var cachableModels = []any{
	(*currency.Model)(nil),
	(*merchant.Model)(nil),
}

func registerCachableTypes() {
	for _, model := range cachableModels {
		escache.RegisterType(model)
	}
}
