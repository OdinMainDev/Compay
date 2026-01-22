package model

import (
	"efw/core/efwlog"
	"fmt"
)

var ErrorModelMethod = func(method, err string) error {
	return fmt.Errorf("%s %s: %s", "[EFW: MODEL]", method, err)
}

func GetModels[T Model, F Filter[T]](filter F) []T {
	r, err := getModels(filter)
	if err != nil {
		efwlog.LogError(ErrorModelMethod("GetModels", err.Error()))
	}
	return r
}

func GetModel[T Model, F Filter[T]](filter F) T {
	r, err := getModel(filter)
	if err != nil {
		efwlog.LogError(ErrorModelMethod("GetModel", err.Error()))
	}
	return r
}

func GetCount[T Model, F Filter[T]](filter F) int64 {
	r, err := getCount(filter)
	if err != nil {
		efwlog.LogError(ErrorModelMethod("GetCount", err.Error()))
	}
	return r
}

func Save[T Model, F Filter[T]](model T, filter F) {
	err := save(model, filter)
	if err != nil {
		efwlog.LogError(ErrorModelMethod("Save", err.Error()))
	}
}

func Delete[T Model, F Filter[T]](model T, filter F) {
	err := delete(model, filter)
	if err != nil {
		efwlog.LogError(ErrorModelMethod("Delete", err.Error()))
	}
}
