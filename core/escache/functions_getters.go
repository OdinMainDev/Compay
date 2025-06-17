package escache

import (
	"compay/core/eslog"
	"errors"

	"github.com/eko/gocache/lib/v4/store"
)

func Get[Model any](args ...int64) []Model {
	key := createKey[Model](args...)
	ret := GetByKey[Model](key)
	return ret
}

func GetByKey[Model any](key string) []Model {
	dto := getDtoByKey(key)
	if dto == nil {
		return nil
	}
	data, ok := dto.Data.([]Model)
	if ok {
		return data
	}
	// TODO: error log
	return []Model{}
}

func getDtoByKey(key string) *cacheDTO {
	ret, err := manager.Get(ctx, key)
	if err != nil {
		if !errors.Is(err, store.NotFound{}) {
			eslog.LogError(err.Error())
		}
		return nil
	}
	result, ok := ret.(cacheDTO)
	if ok {
		return &result
	}

	eslog.LogError("Wrong model type for key", key)
	return nil
}
