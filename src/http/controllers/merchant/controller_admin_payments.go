package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/merchant"
	"compay/src/models/payment"
	"compay/src/models/user"
	"efw/services/efwcontroller"

	"github.com/gin-gonic/gin"
)

func AdminPaymentsGet(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.GetModel(
		context,
		payment.GetModel,
		payment.DTOAdmin,
		func(id int64) payment.QueryHelper {
			qH := payment.QueryHelper{Id: types.PaymentId(id), DontCheckMerchantId: true}
			if !isUserRootAdmin(adminUserModel.GetId()) {
				merchantModels := merchant.GetModels(merchant.QueryHelper{CreatorId: adminUserModel.GetId()})
				merchantIds := []types.MerchantId{}
				for _, merchantModel := range merchantModels {
					merchantIds = append(merchantIds, merchantModel.GetId())
				}
				if len(merchantIds) != 0 {
					merchantIds = append(merchantIds, 0)
				}
				qH.MerchantIds = merchantIds
			}
			return qH
		},
	)
}
