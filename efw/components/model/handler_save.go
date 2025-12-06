package model

import (
	"context"
	"efw"
	"efw/core/efwcache"
	"fmt"
	"reflect"
)

var ErrorModelIsNilOnSave = fmt.Errorf("%s", "[EFW: MODEL] ModelIsNilOnSave")

func save[T Model](model *T, additionalCacheKeys ...int64) error {
	if model == nil || isNil(*model) {
		return ErrorModelIsNilOnSave
	}

	ensureBaseEntity(model)

	(*model).Lock()
	(*model).UpdateTimestamps()
	(*model).Unlock()

	if isDatabaseModel(*model) {
		if err := saveToDb(model); err != nil {
			return err
		}
	}

	if isCacheModel(*model) {
		if err := updateCache(model, additionalCacheKeys...); err != nil {
			return err
		}
	}

	return nil
}

func saveToDb[T Model](model *T) error {
	if model == nil {
		return ErrorModelIsNilOnSave
	}

	query := efw.DBTransaction[T]()
	if (*model).GetIdInt64() == 0 {
		return query.Create(context.Background(), model)
	}

	chain := query.Scopes()
	_, err := chain.Where("id = ?", (*model).GetIdInt64()).Updates(context.Background(), *model)
	return err
}

func updateCache[T Model](model *T, additionalCacheKeys ...int64) error {
	if model == nil {
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

func ensureBaseEntity[T any](model *T) {
	if model == nil {
		return
	}

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
