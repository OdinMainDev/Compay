package model

import (
	"context"
	"efw"
	"efw/core/efwcache"
	"fmt"
)

var ErrorModelIsNilOnSave = fmt.Errorf("%s", "[EFW: MODEL] ModelIsNilOnSave")

func save[T Model](model T, additionalCacheKeys ...int64) error {
	if isNil(model) {
		return ErrorModelIsNilOnSave
	}

	model.Lock()
	model.UpdateTimestamps()
	model.Unlock()

	if isDatabaseModel(model) {
		if err := saveToDb(model); err != nil {
			return err
		}
	}

	if isCacheModel(model) {
		if err := updateCache(model, additionalCacheKeys...); err != nil {
			return err
		}
	}

	return nil
}

func saveToDb[T Model](model T) error {
	query := efw.DBTransaction[T]()
	if model.GetId() == 0 {
		return query.Create(context.Background(), &model)
	}

	chain := query.Scopes()
	_, err := chain.Where("id = ?", model.GetId()).Updates(context.Background(), model)
	return err
}

func updateCache[T Model](model T, additionalCacheKeys ...int64) error {
	if err := efwcache.Update(model, additionalCacheKeys...); err != nil {
		return err
	}
	return nil
}
