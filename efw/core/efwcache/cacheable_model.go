package efwcache

import (
	"efw/utils"
	"encoding/gob"
	"reflect"
)

// TODO: добавить мутекс
var cacheableTypes = []reflect.Type{}

func RegisterType(m any) {
	modelType := reflect.TypeOf(m)

	if utils.InArray(modelType, cacheableTypes) {
		return
	}

	cacheableTypes = append(cacheableTypes, modelType)
	gob.Register(m)
}

func IsCacheableModel(m any) bool {
	return utils.InArray(reflect.TypeOf(m), cacheableTypes)
}
