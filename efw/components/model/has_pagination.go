package model

import (
	"efw/config"

	"gorm.io/gorm"
)

// HasPagination это интегрируемая структура которая обозначает что у QueryHelper есть пагинация
type HasPagination struct {
	HasLimit      bool
	HasPagination bool
	PageNumber    int
	ElementsCount int
}

func getHasPagination[T Model](filter Filter[T]) *HasPagination {
	return filter.GetPagination()
}

func applyHasPagination[T Model](chain gorm.ChainInterface[T], pagination *HasPagination) gorm.ChainInterface[T] {
	if pagination == nil || (!pagination.HasLimit && !pagination.HasPagination) {
		return chain
	}

	if pagination.ElementsCount <= 0 {
		pagination.ElementsCount = config.PaginationDefaultElementsCount
	}

	if pagination.HasPagination || pagination.HasLimit {
		chain = chain.Limit(pagination.ElementsCount)
	}

	if pagination.HasPagination {
		page := max(pagination.PageNumber, 1)
		chain = chain.Offset((page - 1) * pagination.ElementsCount)
	}

	return chain
}
