//go:build testmode

package escache

import (
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/lib/v4/cache"
	ristretto_store "github.com/eko/gocache/store/ristretto/v4"
	"time"
)

var ristrettoCache *ristretto.Cache

func InitForTests() {
	chainCaches := []cache.SetterCacheInterface[any]{}

	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000000,
		MaxCost:     100000,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}
	ristrettoStore := ristretto_store.NewRistretto(ristrettoCache)
	eslocal = newESLocalStore(cache.New[interface{}](ristrettoStore))
	chainCaches = append(chainCaches, eslocal)

	manager = cache.NewChain[interface{}](chainCaches...)
}

func Wait() {
	ristrettoCache.Wait()
	time.Sleep(100 * time.Millisecond)
}
