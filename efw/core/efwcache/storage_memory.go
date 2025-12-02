package efwcache

import (
	"context"

	"github.com/dgraph-io/ristretto/v2"
)

type MemoryStorage struct {
	cache *ristretto.Cache[string, []CacheItem]
}

func newMemoryStorage() (*MemoryStorage, error) {
	cache, err := ristretto.NewCache(&ristretto.Config[string, []CacheItem]{
		NumCounters: 1000000,
		MaxCost:     100000,
		BufferItems: 64,
	})
	if err != nil {
		return nil, ErrorStorageInit("[MEMORY] " + err.Error())
	}

	return &MemoryStorage{
		cache: cache,
	}, nil
}

func (that *MemoryStorage) get(ctx context.Context, key string) ([]CacheItem, error) {
	items, found := that.cache.Get(key)
	if !found {
		return []CacheItem{}, ErrorNilStorage
	}

	result := make([]CacheItem, len(items))
	copy(result, items)
	return result, nil
}

func (that *MemoryStorage) set(ctx context.Context, key string, items []CacheItem) error {
	itemsCopy := make([]CacheItem, len(items))
	copy(itemsCopy, items)

	cost := int64(len(itemsCopy))

	that.cache.Set(key, itemsCopy, cost)
	that.cache.Wait()
	return nil
}

func (that *MemoryStorage) delete(ctx context.Context, key string) error {
	that.cache.Del(key)
	return nil
}

func (that *MemoryStorage) clear(ctx context.Context) error {
	that.cache.Clear()
	return nil
}

func (that *MemoryStorage) close() error {
	that.cache.Close()
	return nil
}
