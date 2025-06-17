package requests

import (
	"compay/application/entities/datatable"
	"compay/application/helpers"
	"compay/application/service/esvalidator"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func DatatableResponse[T any](
	context *gin.Context,
	columnsOrder []string,
	getCount func(casinoId int64) int64,
	additionalFilters map[string]string,
) {
	var req datatable.RequestDatatable
	if err := context.MustBindWith(&req, binding.Form); err != nil {
		helpers.GinErrorResponse(context, esvalidator.ValidatorResponse(err))
		return
	}
	if err := req.ValidateOrderColumn(columnsOrder); err != nil {
		helpers.GinErrorResponse(context, err)
		return
	}

	for key, value := range additionalFilters {
		req.FilterName = append(req.FilterName, key)
		req.FilterValue = append(req.FilterValue, value)
	}
	countRecords := getCount(req.CasinoId)
	if req.Start > countRecords {
		helpers.GinSuccessResponse(context, gin.H{"data": req.ResponseEmpty(countRecords)})
		return
	}

	var dataSlice []*T
	data := datatable.DataTableQuery(dataSlice, req, columnsOrder)
	helpers.GinSuccessResponse(context, gin.H{"data": req.Response(countRecords, data)})
}
