package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/merchant"
	"compay/src/models/user"
	"compay/src/models/withdraw"
	"efw/services/efwcontroller"

	"github.com/gin-gonic/gin"
)

func AdminMerchantWithdrawsGetAll(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.GetModels(
		context,
		withdraw.GetModels,
		withdraw.DTOAdmin,
		func(id int64) withdraw.QueryHelper {
			qH := withdraw.QueryHelper{MerchantWalletId: types.MerchantWalletId(id), DontCheckMerchantId: true}
			if !isUserRootAdmin(adminUserModel.GetId()) {
				merchantModels := merchant.GetModels(merchant.QueryHelper{CreatorId: adminUserModel.GetId()})
				merchantIds := []types.MerchantId{}
				for _, merchantModel := range merchantModels {
					merchantIds = append(merchantIds, merchantModel.GetId())
				}
				if len(merchantIds) == 0 {
					merchantIds = append(merchantIds, 0)
				}
				qH.MerchantIds = merchantIds
			}
			return qH
		},
	)
}
