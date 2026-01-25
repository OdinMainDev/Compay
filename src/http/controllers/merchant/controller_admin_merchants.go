package merchant

import (
	"compay/src/models/merchant"
	"efw/components/model"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

func AdminMerchantsSearch(context *gin.Context) {
	var input requestSearch
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchant.GetModels(merchant.QueryHelper{
		Search: input.Search,
		HasPagination: model.HasPagination{
			HasLimit: true,
		},
	})})
}
