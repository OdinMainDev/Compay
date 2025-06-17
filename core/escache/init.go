package escache

import (
	"compay/core/config"
	"context"
	"reflect"

	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/lib/v4/cache"
	redis_store "github.com/eko/gocache/store/redis/v4"
	ristretto_store "github.com/eko/gocache/store/ristretto/v4"
	"github.com/redis/go-redis/v9"
)

var (
	ctx     = context.Background()
	esredis *ESRedisStore
	eslocal *ESLocalStore
	types   = make(map[reflect.Type]bool)
	manager *cache.ChainCache[any]
)

type CachableModel interface {
	GetId() int64
}

func CreateInstance() {
	chainCaches := []cache.SetterCacheInterface[any]{}

	if config.MemoryCacheEnabled {
		initLocalCache()
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
	}

	if config.RedisEnable {
		redisClient := redis.NewClient(&redis.Options{
			Addr: config.RedisHost,
		})
		redisStore := redis_store.NewRedis(redisClient)
		if config.RedisClearAtStartUp {
			redisStore.Clear(ctx)
		}
		esredis = newEsRedisStore(cache.New[interface{}](redisStore))
		chainCaches = append(chainCaches, esredis)
	}

	manager = cache.NewChain(chainCaches...)
}
