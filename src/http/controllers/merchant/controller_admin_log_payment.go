package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/logpayment"
	"compay/src/models/merchant"
	"compay/src/models/user"
	"efw/core/efwgin"
	"efw/services/efwcontroller"
	"efw/services/efwdt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type requestAdminLogPaymentList struct{}

func AdminLogPaymentList(context *gin.Context) {
	var input efwcontroller.RequestWithId
	if !efwgin.ShouldBindUri(context, &input) {
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	merchantIds := []types.MerchantId{}
	if !isUserRootAdmin(adminUserModel.GetId()) {
		merchantModels := merchant.GetModels(merchant.QueryHelper{CreatorId: adminUserModel.GetId()})
		for _, merchantModel := range merchantModels {
			merchantIds = append(merchantIds, merchantModel.GetId())
		}
		if len(merchantIds) != 0 {
			merchantIds = append(merchantIds, 0)
		}
	}

	efwcontroller.Datable[logpayment.Model, requestAdminLogPaymentList](
		context,
		[]string{"id"},
		[]efwdt.ChainFilter[*logpayment.Model]{
			func(chain gorm.ChainInterface[*logpayment.Model]) gorm.ChainInterface[*logpayment.Model] {
				if len(merchantIds) != 0 {
					return chain.Where("merchant_id IN (?)", merchantIds)
				}
				return chain
			},
			func(chain gorm.ChainInterface[*logpayment.Model]) gorm.ChainInterface[*logpayment.Model] {
				return chain.Where("payment_id = ?", input.Id)
			},
		},
		nil,
		nil,
	)
}
