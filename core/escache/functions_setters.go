package escache

import (
	"compay/core/eslog"
)

func Set[Model any](data []Model, args ...int64) {
	key := createKey[Model](args...)
	SetByKey(key, data)
}

func SetNull[Model any](args ...int64) {
	key := createKey[Model](args...)
	SetNullByKey(key)
}

func SetByKey[Model any](key string, data []Model) {
	dto := modelsToCacheDTO(data)
	setDtoByKey(key, dto)
}

func SetNullByKey(key string) {
	err := manager.Delete(ctx, key)
	if err != nil {
		eslog.LogError(err.Error())
		return
	}
}

func setDtoByKey(key string, dto cacheDTO) {
	err := manager.Set(ctx, key, dto)
	if err != nil {
		eslog.LogError(err.Error())
		return
	}
}
