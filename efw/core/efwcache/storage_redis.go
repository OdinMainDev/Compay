package efwcache

import (
	"bytes"
	"context"
	"efw/config"
	"encoding/gob"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
)

type RedisStorage struct {
	client *redis.Client
}

func newRedisStorage() (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Username: config.RedisUsername,
		Password: config.RedisPassword,
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, ErrorStorageInit("[REDIS] " + err.Error())
	}

	resp := &RedisStorage{
		client: client,
	}

	if config.RedisClearAtStartUp {
		resp.clear(context.Background())
	}

	return resp, nil
}

func (that *RedisStorage) get(ctx context.Context, key string) ([]CacheItem, error) {
	data, err := that.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return []CacheItem{}, ErrorNilStorage
		}
		return nil, ErrorStorageGet("[REDIS] " + err.Error())
	}

	items, err := that.unpack(data)
	if err != nil {
		return nil, ErrorStorageGet("[REDIS] unpack error: " + err.Error())
	}

	return items, nil
}

func (that *RedisStorage) set(ctx context.Context, key string, items []CacheItem) error {
	safeCopy := make([]CacheItem, len(items))
	for i, item := range items {
		safeCopy[i] = CacheItem{
			Id:   item.Id,
			Data: that.createSafeCopy(item.Data),
		}
	}

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(safeCopy)
	if err != nil {
		return ErrorStorageSet("[REDIS] " + err.Error())
	}

	cmd := that.client.Set(ctx, key, buff.Bytes(), 0)
	if cmd.Err() != nil {
		return ErrorStorageSet("[REDIS] " + cmd.Err().Error())
	}

	return nil
}

func (that *RedisStorage) delete(ctx context.Context, key string) error {
	cmd := that.client.Del(ctx, key)
	if cmd.Err() != nil {
		return ErrorStorageDelete("[REDIS] " + cmd.Err().Error())
	}

	return nil
}

func (that *RedisStorage) clear(ctx context.Context) error {
	cmd := that.client.FlushDB(ctx)
	if cmd.Err() != nil {
		return ErrorStorageClear("[REDIS] " + cmd.Err().Error())
	}

	return nil
}

func (that *RedisStorage) close() error {
	err := that.client.Close()
	if err != nil {
		return ErrorStorageClose("[REDIS] " + err.Error())
	}

	return nil
}

func (*RedisStorage) unpack(value any) ([]CacheItem, error) {
	var valueBytes []byte
	switch v := value.(type) {
	case []byte:
		valueBytes = v
	case string:
		valueBytes = []byte(v)
	default:
		return nil, ErrorUnpackBadType
	}
	reader := bytes.NewReader(valueBytes)
	dec := gob.NewDecoder(reader)
	var items []CacheItem
	err := dec.Decode(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (that *RedisStorage) createSafeCopy(data any) any {
	if data == nil {
		return nil
	}

	return that.deepCopy(data)
}

func (that *RedisStorage) deepCopy(src any) any {
	if src == nil {
		return nil
	}

	original := reflect.ValueOf(src)
	cpy := reflect.New(original.Type()).Elem()
	that.copyRecursive(original, cpy)
	return cpy.Interface()
}

func (that *RedisStorage) copyRecursive(original, cpy reflect.Value) {
	switch original.Kind() {
	case reflect.Ptr:
		if !original.IsNil() {
			originalValue := original.Elem()
			copyValue := reflect.New(originalValue.Type())
			that.copyRecursive(originalValue, copyValue.Elem())
			cpy.Set(copyValue)
		}
	case reflect.Interface:
		if !original.IsNil() {
			copyValue := reflect.New(original.Elem().Type()).Elem()
			that.copyRecursive(original.Elem(), copyValue)
			cpy.Set(copyValue)
		}
	case reflect.Struct:
		for i := 0; i < original.NumField(); i++ {
			field := original.Field(i)
			// Do not check mutex and priv fields
			if !field.CanInterface() {
				continue
			}
			fieldType := original.Type().Field(i)
			if fieldType.Name == "mu" || fieldType.Type.String() == "sync.Mutex" {
				continue
			}
			that.copyRecursive(field, cpy.Field(i))
		}
	case reflect.Slice:
		if !original.IsNil() {
			cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
			for i := 0; i < original.Len(); i++ {
				that.copyRecursive(original.Index(i), cpy.Index(i))
			}
		}
	case reflect.Map:
		if !original.IsNil() {
			cpy.Set(reflect.MakeMap(original.Type()))
			for _, key := range original.MapKeys() {
				originalValue := original.MapIndex(key)
				copyValue := reflect.New(originalValue.Type()).Elem()
				that.copyRecursive(originalValue, copyValue)
				cpy.SetMapIndex(key, copyValue)
			}
		}
	default:
		cpy.Set(original)
	}
}
