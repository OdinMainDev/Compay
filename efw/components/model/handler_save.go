package model

import (
	"context"
	"efw"
	"efw/core/efwcache"
	"fmt"
	"reflect"
)

var ErrorModelIsNilOnSave = fmt.Errorf("%s", "[EFW: MODEL] ModelIsNilOnSave")

func save[T Model, F Filter[T]](model T, filter F) error {
	if isNil(model) {
		return ErrorModelIsNilOnSave
	}

	ensureBaseEntity(model)

	model.Lock()
	model.UpdateTimestamps()
	model.Unlock()

	if isDatabaseModel(model) {
		if err := saveToDb(model); err != nil {
			return err
		}
	}

	if isCacheModel(model) {
		if err := updateCache(model, filter.CacheAdditionalKeys()...); err != nil {
			return err
		}
	}

	return nil
}

func saveToDb[T Model](model T) error {
	if isNil(model) {
		return ErrorModelIsNilOnSave
	}

	query := efw.DBTransaction[T]()
	if model.GetIdInt64() == 0 {
		return query.Create(context.Background(), &model)
	}

	chain := query.Scopes()
	rows, err := chain.Updates(context.Background(), model)
	if err != nil {
		return err
	}
	if rows == 0 {
		return query.Create(context.Background(), &model)
	}
	return nil
}

func updateCache[T Model](model T, additionalCacheKeys ...int64) error {
	if isNil(model) {
		return ErrorModelIsNilOnSave
	}

	cacheAny := any(model)
	cacheModel, ok := cacheAny.(efwcache.CacheableModel)
	if !ok {
		return fmt.Errorf("%s", "[EFW: MODEL] ModelIsNotCacheableModelOnUpdate")
	}

	if err := efwcache.Update(cacheModel, additionalCacheKeys...); err != nil {
		return err
	}
	return nil
}

func ensureBaseEntity[T any](model T) {
	val := reflect.ValueOf(model)
	if val.Kind() != reflect.Pointer || val.IsNil() {
		return
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return
	}

	field := val.FieldByName("BaseEntity")
	if !field.IsValid() || field.Kind() != reflect.Pointer {
		return
	}

	if field.IsNil() {
		field.Set(reflect.New(field.Type().Elem()))
	}
}
