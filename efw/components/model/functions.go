package model

import (
	"efw/components/entity"
	"reflect"
)

func isCacheModel[T Model](model T) bool {
	return model.GetSaveType() != entity.SAVE_TYPE_DB
}

func isDatabaseModel[T Model](model T) bool {
	if model.GetSaveType() == entity.SAVE_TYPE_CACHE || model.GetSaveType() == entity.SAVE_TYPE_LOCAL_CACHE {
		return false
	}
	return true
}

func isNil(value any) bool {
	if value == nil {
		return true
	}
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return val.IsNil()
	default:
		return false
	}
}
