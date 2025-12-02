package efwcache

import (
	"context"
)

type AllInOneStorage struct {
	memory Storage
	redis  Storage
}

func newAllInOneStorage(memory, redis Storage) *AllInOneStorage {
	return &AllInOneStorage{
		memory: memory,
		redis:  redis,
	}
}

func (that *AllInOneStorage) get(ctx context.Context, key string) ([]CacheItem, error) {
	items, err := that.memory.get(ctx, key)
	if err == nil && len(items) > 0 {
		return items, nil
	}

	items, err = that.redis.get(ctx, key)
	if err != nil {
		return nil, err
	}

	if len(items) > 0 {
		err = that.memory.set(ctx, key, items)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (that *AllInOneStorage) set(ctx context.Context, key string, items []CacheItem) error {
	if err := that.redis.set(ctx, key, items); err != nil {
		return err
	}

	return that.memory.set(ctx, key, items)
}

func (that *AllInOneStorage) delete(ctx context.Context, key string) error {
	that.memory.delete(ctx, key) // игнорируем ошибку ( постоянно nil )
	return that.redis.delete(ctx, key)
}

func (that *AllInOneStorage) clear(ctx context.Context) error {
	that.memory.clear(ctx) // игнорируем ошибку ( постоянно nil )
	return that.redis.clear(ctx)
}

func (that *AllInOneStorage) close() error {
	that.memory.close() // игнорируем ошибку ( постоянно nil )
	return that.redis.close()
}
