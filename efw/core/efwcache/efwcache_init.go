package efwcache

import (
	"fmt"
	"math/bits"
)

func initCache(config CacheConfig) (*EfwCache, error) {
	err := checkConfig(&config)
	if err != nil {
		return nil, err
	}

	storage, err := createStorage(config)
	if err != nil {
		return nil, ErrorCreateStorage(err.Error())
	}

	cache := &EfwCache{
		config:           config,
		shards:           make([]*Shard, config.CountOfShards),
		shardMask:        uint32(config.CountOfShards - 1),
		storage:          storage,
		operationTimeout: config.OperationTimeout,
	}

	for i := 0; i < config.CountOfShards; i++ {
		shard := &Shard{
			id:         i,
			operations: make(chan CacheOperation, config.OperationQueueSize),
			storage:    storage,
			timeout:    config.OperationTimeout,
		}
		shard.wg.Add(1)
		go shard.worker()
		cache.shards[i] = shard
	}

	return cache, nil
}

func checkConfig(config *CacheConfig) error {
	if config.CountOfShards <= 0 {
		config.CountOfShards = DefautConfig.CountOfShards
	}

	if bits.OnesCount(uint(config.CountOfShards)) != 1 {
		return ErrorInvalidCountOfShards
	}

	if config.OperationQueueSize <= 0 {
		config.OperationQueueSize = DefautConfig.OperationQueueSize
	}

	if config.OperationTimeout <= 0 {
		config.OperationTimeout = DefautConfig.OperationTimeout
	}

	switch config.CacheType {
	case CacheTypeMemory, CacheTypeRedis, CacheTypeAllInOne:
	default:
		return ErrorInvalidCacheType(fmt.Sprint(config.CacheType))
	}

	return nil
}

func createStorage(config CacheConfig) (Storage, error) {
	switch config.CacheType {
	case CacheTypeMemory:
		return newMemoryStorage()

	case CacheTypeRedis:
		return newRedisStorage()

	case CacheTypeAllInOne:
		memory, err := newMemoryStorage()
		if err != nil {
			return nil, err
		}
		redis, err := newRedisStorage()
		if err != nil {
			return nil, err
		}
		return newAllInOneStorage(memory, redis), nil

	default:
		return nil, ErrorInvalidCacheType(fmt.Sprint(config.CacheType))
	}
}
