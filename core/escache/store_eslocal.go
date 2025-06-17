package escache

import (
	"context"
	"time"

	"github.com/eko/gocache/lib/v4/codec"
	"github.com/eko/gocache/lib/v4/store"
)

type ESLocalStore struct {
	codec codec.CodecInterface
	cache store.StoreInterface
}

func (ls *ESLocalStore) Get(ctx context.Context, key any) (any, error) {
	return ls.cache.Get(ctx, key)
}

func (ls *ESLocalStore) Set(ctx context.Context, key any, object any, options ...store.Option) error {
	return ls.cache.Set(ctx, key, object, options...)
}

func (ls *ESLocalStore) Delete(ctx context.Context, key any) error {
	return ls.cache.Delete(ctx, key)
}

func (ls *ESLocalStore) Invalidate(ctx context.Context, options ...store.InvalidateOption) error {
	return ls.cache.Invalidate(ctx, options...)
}

func (ls *ESLocalStore) Clear(ctx context.Context) error {
	return ls.cache.Clear(ctx)
}

func (ls *ESLocalStore) GetType() string {
	return ls.cache.GetType()
}

func (ls *ESLocalStore) GetWithTTL(ctx context.Context, key any) (any, time.Duration, error) {
	return ls.cache.GetWithTTL(ctx, key)
}

func (ls *ESLocalStore) GetCodec() codec.CodecInterface {
	return ls.codec
}

func newESLocalStore(cache store.StoreInterface) *ESLocalStore {
	return &ESLocalStore{
		codec: codec.New(cache),
		cache: cache,
	}
}
