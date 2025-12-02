package efwcache

import (
	"crypto/md5"
	"encoding/binary"
)

func (c *EfwCache) getShard(key string) *Shard {
	hash := md5.Sum([]byte(key))
	shardIndex := binary.LittleEndian.Uint32(hash[:4]) & c.shardMask
	return c.shards[shardIndex]
}

func (c *EfwCache) get(key string) ([]CacheItem, error) {
	shard := c.getShard(key)

	respChan := make(chan CacheResponse, 1)
	operation := CacheOperation{
		OperationType: OperationGet,
		Key:           key,
		Response:      respChan,
	}

	shard.operations <- operation
	resp := <-respChan
	return resp.Items, resp.Error
}

func (c *EfwCache) set(key string, items []CacheItem) error {
	shard := c.getShard(key)

	respChan := make(chan CacheResponse, 1)
	operation := CacheOperation{
		OperationType: OperationSet,
		Key:           key,
		Items:         items,
		Response:      respChan,
	}

	shard.operations <- operation
	resp := <-respChan
	return resp.Error
}

func (c *EfwCache) update(key string, item CacheItem) error {
	shard := c.getShard(key)

	respChan := make(chan CacheResponse, 1)
	operation := CacheOperation{
		OperationType: OperationUpdate,
		Key:           key,
		Item:          item,
		Response:      respChan,
	}

	shard.operations <- operation
	resp := <-respChan
	return resp.Error
}

func (c *EfwCache) delete(key string, itemID int64) error {
	shard := c.getShard(key)

	respChan := make(chan CacheResponse, 1)
	operation := CacheOperation{
		OperationType: OperationDelete,
		Key:           key,
		Item:          CacheItem{Id: itemID},
		Response:      respChan,
	}

	shard.operations <- operation
	resp := <-respChan
	return resp.Error
}
