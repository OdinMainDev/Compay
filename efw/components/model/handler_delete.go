package model

import (
	"context"
	"efw"
	"efw/core/efwcache"
	"fmt"
)

var ErrorModelIsNilOnDelete = fmt.Errorf("%s", "[EFW: MODEL] ModelIsNilOnDelete")

func delete[T Model](model T, additionalCacheKeys ...int64) error {
	if isNil(model) {
		return ErrorModelIsNilOnDelete
	}

	if isDatabaseModel(model) {
		if err := deleteFromDb(model); err != nil {
			return err
		}
	}

	if isCacheModel(model) {
		if err := deleteFromCache(model, additionalCacheKeys...); err != nil {
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
	if err := efwcache.Delete(model, additionalCacheKeys...); err != nil {
		return err
	}
	return nil
}
