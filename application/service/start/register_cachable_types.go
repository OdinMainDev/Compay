package start

import (
	"compay/application/entities/currency"
	"compay/application/entities/user"
	"compay/application/entities/wallet"
	"compay/core/escache"
)

var cachableModels = []any{
	(*currency.Model)(nil),
	(*wallet.Model)(nil),
	(*user.Model)(nil),
}

func registerCachableTypes() {
	for _, model := range cachableModels {
		escache.RegisterType(model)
	}
}
