package escache

import (
	"compay/core/eslog"

	"github.com/dgraph-io/ristretto/v2"
)

var localCache *ristretto.Cache[[]byte, any]

func initLocalCache() {
	c, err := ristretto.NewCache(&ristretto.Config[[]byte, any]{
		NumCounters: 1e7,
		MaxCost:     500000000,
		BufferItems: 64,
	})
	if err != nil {
		eslog.PanicLog("BAD STARTUP CREATE LOCAL CACHE: ", err.Error())
	}
	localCache = c
}

func GetFromLocalCache[T any](key []byte) (T, bool) {
	value, found := localCache.Get(key)
	if !found {
		var zero T
		return zero, false
	}

	if result, ok := value.(T); ok {
		return result, true
	}

	var zero T
	return zero, false
}

func SetToLocalCache(key []byte, data any) {
	localCache.Set(key, data, 1)
}

func DeleteFromLocalCache(key []byte) {
	localCache.Del(key)
}
