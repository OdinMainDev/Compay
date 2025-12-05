package model

import (
	"efw/core/efwlog"
	"fmt"
)

var ErrorModelMethod = func(method, err string) error {
	return fmt.Errorf("%s %s: %s", "[EFW: MODEL]", method, err)
}

// func GetModels[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) ([]*T, error) {
// 	return getModels(filter, additionalCacheKeys...)
// }

// func Save[T Model](model T, additionalCacheKeys ...int64) error {
// 	return save(model, additionalCacheKeys...)
// }

// func Delete[T Model](model T, additionalCacheKeys ...int64) error {
// 	return delete(model, additionalCacheKeys...)
// }

// func GetModel[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) (*T, error) {
// 	return getModel(filter, additionalCacheKeys...)
// }

// func GetCount[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) (int64, error) {
// 	return getCount(filter, additionalCacheKeys...)
// }

func GetModels[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) []*T {
	r, err := getModels(filter, additionalCacheKeys...)
	if err != nil {
		efwlog.LogError(ErrorModelMethod("GetModels", err.Error()))
	}
	return r
}

func Save[T Model](model T, additionalCacheKeys ...int64) {
	err := save(model, additionalCacheKeys...)
	if err != nil {
		efwlog.LogError(ErrorModelMethod("Save", err.Error()))
	}
}

func Delete[T Model](model T, additionalCacheKeys ...int64) {
	err := delete(model, additionalCacheKeys...)
	if err != nil {
		efwlog.LogError(ErrorModelMethod("Delete", err.Error()))
	}
}

func GetModel[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) *T {
	r, err := getModel(filter, additionalCacheKeys...)
	if err != nil {
		efwlog.LogError(ErrorModelMethod("GetModel", err.Error()))
	}
	return r
}

func GetCount[T Model, F Filter[T]](filter F, additionalCacheKeys ...int64) int64 {
	r, err := getCount(filter, additionalCacheKeys...)
	if err != nil {
		efwlog.LogError(ErrorModelMethod("GetCount", err.Error()))
	}
	return r
}
