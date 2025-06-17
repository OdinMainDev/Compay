package escache

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"reflect"
	"time"

	"github.com/eko/gocache/lib/v4/codec"
	"github.com/eko/gocache/lib/v4/store"
)

type ESRedisStore struct {
	codec codec.CodecInterface
	cache store.StoreInterface
}

func (rs *ESRedisStore) Get(ctx context.Context, key any) (any, error) {
	value, err := rs.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	return rs.unpack(value)
}

func (rs *ESRedisStore) Set(ctx context.Context, key any, object any, options ...store.Option) error {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(object)
	if err != nil {
		return err
	}
	return rs.cache.Set(ctx, key, buff.Bytes(), options...)
}

func (rs *ESRedisStore) Delete(ctx context.Context, key any) error {
	return rs.cache.Delete(ctx, key)
}

func (rs *ESRedisStore) Invalidate(ctx context.Context, options ...store.InvalidateOption) error {
	return rs.cache.Invalidate(ctx, options...)
}

func (rs *ESRedisStore) Clear(ctx context.Context) error {
	return rs.cache.Clear(ctx)
}

func (rs *ESRedisStore) GetType() string {
	return rs.cache.GetType()
}

func (rs *ESRedisStore) GetWithTTL(ctx context.Context, key any) (any, time.Duration, error) {
	value, ttl, err := rs.cache.GetWithTTL(ctx, key)
	if err != nil {
		return nil, 0, err
	}
	value, err = rs.unpack(value)
	if err != nil {
		return nil, 0, err
	}
	return value, ttl, nil
}

func (rs *ESRedisStore) GetCodec() codec.CodecInterface {
	return rs.codec
}

func (rs *ESRedisStore) RegisterType(t reflect.Type) {
	sliceOfModels := reflect.MakeSlice(reflect.SliceOf(t), 0, 0).Interface()
	gob.Register(sliceOfModels)
}

func (rs *ESRedisStore) unpack(value any) (any, error) {
	var valueBytes []byte
	switch v := value.(type) {
	case []byte:
		valueBytes = v
	case string:
		valueBytes = []byte(v)
	default:
		return nil, fmt.Errorf("an error: content not []bytes or string")
	}
	reader := bytes.NewReader(valueBytes)
	dec := gob.NewDecoder(reader)
	var v cacheDTO
	err := dec.Decode(&v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func newEsRedisStore(cache store.StoreInterface) *ESRedisStore {
	gob.Register(cacheDTO{})
	return &ESRedisStore{
		codec: codec.New(cache),
		cache: cache,
	}
}
