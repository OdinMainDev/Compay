package efwdt

import "github.com/gin-gonic/gin"

type response[T any] struct {
	Data            T     `json:"data"`
	RecordsTotal    int64 `json:"recordsTotal"`
	RecordsFiltered int64 `json:"recordsFiltered"`
	RecordsCount    int64 `json:"recordsCount"`
}

func Response[T any](data T, recordsTotal, recordsFiltered, recordsCount int64) gin.H {
	return gin.H{
		"data": response[T]{
			Data:            data,
			RecordsTotal:    recordsTotal,
			RecordsFiltered: recordsFiltered,
			RecordsCount:    recordsCount,
		},
	}
}
