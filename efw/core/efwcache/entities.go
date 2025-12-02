package efwcache

import (
	"context"
	"sync"
	"time"
)

// EfwCache представляет собой фасад кэша, отвечающий за фрагментацию и делегирование хранения
type EfwCache struct {
	config           CacheConfig
	shards           []*Shard
	shardMask        uint32
	storage          Storage
	operationTimeout time.Duration
	closeOnce        sync.Once
}

// Шарды обслуживают очередь операций, назначенную слоту хеша шарда
type Shard struct {
	id         int
	operations chan CacheOperation
	storage    Storage
	timeout    time.Duration
	wg         sync.WaitGroup
}

type CacheableModel interface {
	GetId() int64
}

type CacheItem struct {
	Id   int64 `json:"id"`
	Data any   `json:"data"`
}

type CacheConfig struct {
	CacheType          CacheType
	CountOfShards      int           // Количество шардов для параллелизма
	OperationQueueSize int           // Размер очереди операций на шарде
	OperationTimeout   time.Duration // Таймаут на выполнение операции
}

var DefautConfig = CacheConfig{
	CacheType:          CacheTypeAllInOne,
	CountOfShards:      16,
	OperationQueueSize: 1024,
	OperationTimeout:   30 * time.Second,
}

type Storage interface {
	get(ctx context.Context, key string) ([]CacheItem, error)
	set(ctx context.Context, key string, items []CacheItem) error
	delete(ctx context.Context, key string) error
	clear(ctx context.Context) error
	close() error
}

// CacheOperation описывает одну команду кэша, выполняемую последовательно для каждого шарда
type CacheOperation struct {
	Key           string
	OperationType OperationType
	Item          CacheItem
	Items         []CacheItem
	Response      chan CacheResponse
}

type CacheResponse struct {
	Items []CacheItem
	Error error
}
