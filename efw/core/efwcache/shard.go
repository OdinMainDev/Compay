package efwcache

import (
	"context"
	"errors"
)

func (s *Shard) worker() {
	defer s.wg.Done()

	for operation := range s.operations {
		s.processOperation(operation)
	}
}

func (s *Shard) processOperation(operation CacheOperation) {
	ctx := context.Background()
	if s.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, s.timeout)
		defer cancel()
	}

	var (
		resp CacheResponse
		err  error
	)

	switch operation.OperationType {
	case OperationGet:
		resp.Items, err = s.storage.get(ctx, operation.Key)
		if errors.Is(err, ErrorNilStorage) {
			err = nil
		}
		resp.Error = err

	case OperationSet:
		resp.Error = s.storage.set(ctx, operation.Key, operation.Items)

	case OperationUpdate:
		var items []CacheItem
		items, err = s.storage.get(ctx, operation.Key)
		if err != nil && !errors.Is(err, ErrorNilStorage) {
			resp.Error = err
			break
		}
		if errors.Is(err, ErrorNilStorage) {
			items = []CacheItem{}
		}

		updated := false
		for i, item := range items {
			if item.Id == operation.Item.Id {
				items[i] = operation.Item
				updated = true
				break
			}
		}

		if !updated {
			items = append(items, operation.Item)
		}

		sortDataIfNeeded(items)
		resp.Error = s.storage.set(ctx, operation.Key, items)

	case OperationDelete:
		var items []CacheItem
		items, err = s.storage.get(ctx, operation.Key)
		if err != nil && !errors.Is(err, ErrorNilStorage) {
			resp.Error = err
			break
		}
		if errors.Is(err, ErrorNilStorage) {
			resp.Error = nil
			break
		}

		newItems := make([]CacheItem, 0, len(items))
		for _, item := range items {
			if item.Id != operation.Item.Id {
				newItems = append(newItems, item)
			}
		}

		resp.Error = s.storage.set(ctx, operation.Key, newItems)

	case OperationClear:
		resp.Error = s.storage.clear(ctx)

	default:
		resp.Error = ErrorInvalidOperation
	}

	if operation.Response != nil {
		operation.Response <- resp
	}
}
