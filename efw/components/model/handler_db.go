package model

import (
	"context"
	"efw"
)

func fetchFromDb[T Model](filter Filter[T]) ([]T, error) {
	query := efw.DBTransaction[T]()
	chain := query.Scopes()
	chain = filter.ApplyDBFilters(chain)

	// Контракты
	chain = applyHasPagination(chain, getHasPagination(filter))

	return chain.Find(context.Background())
}

func fetchAllFromDbForCache[T Model](filter Filter[T]) ([]T, error) {
	query := efw.DBTransaction[T]()
	chain := query.Scopes()
	chain = filter.ApplyCacheDBFilters(chain)

	return chain.Find(context.Background())
}
