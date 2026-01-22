package model

import (
	"context"
	"efw"
	"efw/core/efwcache"
	"fmt"
)

var ErrorModelIsNilOnDelete = fmt.Errorf("%s", "[EFW: MODEL] ModelIsNilOnDelete")
var ErrorModelIsNotCacheableModelOnDelete = fmt.Errorf("%s", "[EFW: MODEL] ModelIsNotCacheableModelOnDelete")

func delete[T Model, F Filter[T]](model T, filter F) error {
	if isNil(model) {
		return ErrorModelIsNilOnDelete
	}

	if isDatabaseModel(model) {
		if err := deleteFromDb(model); err != nil {
			return err
		}
	}

	if isCacheModel(model) {
		if err := deleteFromCache(model, filter.CacheAdditionalKeys()...); err != nil {
			return err
		}
	}

	return nil
}

func deleteFromDb[T Model](model T) error {
	if model.GetIdInt64() == 0 {
		return nil
	}

	query := efw.DBTransaction[T]()
	chain := query.Scopes()
	_, err := chain.Where("id = ?", model.GetIdInt64()).Delete(context.Background())
	return err
}

func deleteFromCache[T Model](model T, additionalCacheKeys ...int64) error {
	if isNil(model) {
		return ErrorModelIsNilOnSave
	}

	cacheAny := any(model)
	cacheModel, ok := cacheAny.(efwcache.CacheableModel)
	if !ok {
		return ErrorModelIsNotCacheableModelOnDelete
	}

	if err := efwcache.Delete(cacheModel, additionalCacheKeys...); err != nil {
		return err
	}
	return nil
}
