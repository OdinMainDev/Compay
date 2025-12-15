package efwcache

import (
	"efw/utils"
	"encoding/gob"
	"reflect"
	"sync"
)

var cacheableTypes = []reflect.Type{}
var cacheableTypesMutex = sync.Mutex{}

func RegisterType(m any) {
	modelType := reflect.TypeOf(m)

	if utils.InArray(modelType, cacheableTypes) {
		return
	}

	cacheableTypes = append(cacheableTypes, modelType)
	gob.Register(m)
}

func IsCacheableModel(m any) bool {
	cacheableTypesMutex.Lock()
	defer cacheableTypesMutex.Unlock()
	return utils.InArray(reflect.TypeOf(m), cacheableTypes)
}
