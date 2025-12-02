package efw

import (
	"efw/config"
	"efw/core/efwcache"
	"efw/core/efwgcache"
	"fmt"
)

var ErrorCacheInvalidConfiguration = fmt.Errorf("%s", "[EFW: CACHE INIT] CacheInvalidConfiguration:")

func (that *eframework) initCache() error {
	if that.options.CacheDisabled {
		return nil
	}

	useMemory := !that.options.MemoryCacheDisabled
	useRedis := !that.options.RedisCacheDisabled

	if !useMemory && !useRedis {
		return ErrorCacheInvalidConfiguration
	}

	cacheType := efwcache.CacheTypeMemory
	if useMemory && useRedis {
		cacheType = efwcache.CacheTypeAllInOne
	} else if useRedis {
		cacheType = efwcache.CacheTypeRedis
	}

	cacheConfig := efwcache.CacheConfig{
		CacheType:          cacheType,
		CountOfShards:      config.CacheShardCount,
		OperationQueueSize: config.CacheOperationQueueSize,
		OperationTimeout:   config.CacheOperationTimeout,
	}

	return efwcache.Init(cacheConfig)
}

func (that *eframework) initGenericCache() error {
	if that.options.GenericCacheDisabled {
		return nil
	}

	return efwgcache.Init()
}
