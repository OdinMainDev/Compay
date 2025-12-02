package model

import (
	"context"
	"efw"
)

func fetchFromDb[T Model](filter Filter[T]) ([]T, error) {
	query := efw.DBTransaction[T]()
	chain := query.Scopes()
	chain = filter.ApplyDBFilters(chain)
	return chain.Find(context.Background())
}
