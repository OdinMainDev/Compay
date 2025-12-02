package efwcache

import "sync"

var (
	instanceMu sync.RWMutex
	instance   *EfwCache
)

func Init(config CacheConfig) error {
	cache, err := initCache(config)
	if err != nil {
		return err
	}

	instanceMu.Lock()
	defer instanceMu.Unlock()

	if instance != nil {
		if err := instance.close(); err != nil {
			cache.close()
			return err
		}
	}

	instance = cache
	return nil
}

func Close() error {
	instanceMu.Lock()
	defer instanceMu.Unlock()

	if instance == nil {
		return nil
	}

	err := instance.close()
	instance = nil
	return err
}

func (c *EfwCache) close() error {
	var closeErr error
	c.closeOnce.Do(func() {
		for _, shard := range c.shards {
			close(shard.operations)
			shard.wg.Wait()
		}
		closeErr = c.storage.close()
	})
	return closeErr
}
