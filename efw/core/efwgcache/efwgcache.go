package efwgcache

import (
	"fmt"

	"github.com/dgraph-io/ristretto/v2"
)

var ErrorRistrettoInitError = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: GCACHE] RistrettoInitError:", err)
}

var genericCache *ristretto.Cache[[]byte, any]

func Init() error {
	c, err := ristretto.NewCache(&ristretto.Config[[]byte, any]{
		NumCounters: 1e7,
		MaxCost:     500000000,
		BufferItems: 64,
	})
	if err != nil {
		return ErrorRistrettoInitError(err.Error())
	}
	genericCache = c
	return nil
}

func Get[T any](key []byte) (T, bool) {
	value, found := genericCache.Get(key)
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

func Set(key []byte, data any) {
	genericCache.Set(key, data, 1)
}

func Delete(key []byte) {
	genericCache.Del(key)
}
