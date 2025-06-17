package start

import (
	"compay/core/escache"
)

var cachableModels = []any{
	// (*wallet.Model)(nil),
}

func registerCachableTypes() {
	for _, model := range cachableModels {
		escache.RegisterType(model)
	}
}
