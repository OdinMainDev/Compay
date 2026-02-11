package model

import (
	"efw/core/efwlog"
	"fmt"
	"reflect"
)

func modelInfo[T any]() (pkg string, name string) {
	var t T
	typ := reflect.TypeOf(t)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.PkgPath(), typ.Name()
}

func ErrorModelMethod[T any](method, err string) error {
	pkg, name := modelInfo[T]()
	return fmt.Errorf(
		"[EFW: MODEL] method=%s model=%s package=%s error=%s",
		method,
		name,
		pkg,
		err,
	)
}

// var ErrorModelMethod = func(method, err, model, filter string) error {
// 	return fmt.Errorf("%s %s: %s (%s || %s)", "[EFW: MODEL]", method, err, model, filter)
// }

func GetModels[T Model, F Filter[T]](filter F) []T {
	r, err := getModels(filter)
	if err != nil {
		efwlog.LogError(ErrorModelMethod[T]("GetModels", err.Error()))
	}
	return r
}

func GetModel[T Model, F Filter[T]](filter F) T {
	r, err := getModel(filter)
	if err != nil {
		efwlog.LogError(ErrorModelMethod[T]("GetModel", err.Error()))
	}
	return r
}

func GetCount[T Model, F Filter[T]](filter F) int64 {
	r, err := getCount(filter)
	if err != nil {
		efwlog.LogError(ErrorModelMethod[T]("GetCount", err.Error()))
	}
	return r
}

func Save[T Model, F Filter[T]](model T, filter F) {
	err := save(model, filter)
	if err != nil {
		efwlog.LogError(ErrorModelMethod[T]("Save", err.Error()))
	}
}

func Delete[T Model, F Filter[T]](model T, filter F) {
	err := delete(model, filter)
	if err != nil {
		efwlog.LogError(ErrorModelMethod[T]("Delete", err.Error()))
	}
}
