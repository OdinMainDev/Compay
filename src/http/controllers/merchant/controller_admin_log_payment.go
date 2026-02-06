package merchant

import (
	"compay/common/types"
	"compay/src/models/logpayment"
	"efw/services/efwcontroller"
	"efw/services/efwdt"

	"github.com/gin-gonic/gin"
)

type requestAdminLogPaymentList struct {
	Id types.PaymentId `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	efwdt.WithCreatedAt
}

func AdminLogPaymentList(context *gin.Context) {
	efwcontroller.Datable[logpayment.Model, requestAdminLogPaymentList](
		context,
		[]string{"id", "created_at"},
		nil,
		nil,
		nil,
	)
}

func AdminLogPaymentListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}
