package merchant

import (
	"compay/common/types"
	"compay/src/models/logmerchantwallet"
	"efw/services/efwcontroller"
	"efw/services/efwdt"

	"github.com/gin-gonic/gin"
)

type requestAdminLogMerchantWalletList struct {
	Id types.PaymentId `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	efwdt.WithCreatedAt
}

func AdminLogMerchantWalletList(context *gin.Context) {
	efwcontroller.Datable[logmerchantwallet.Model, requestAdminLogMerchantWalletList](
		context,
		[]string{"id", "created_at"},
		nil,
		nil,
		nil,
	)
}

func AdminLogMerchantWalletListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}
