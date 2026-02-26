package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/loguser"
	"compay/src/models/merchant"
	"compay/src/models/user"
	"efw/components/entity"
	"efw/services/efwcontroller"
	"efw/services/efwdt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type requestAdminLogUserList struct {
	Id         types.LogUserId     `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	MerchantId types.MerchantId    `json:"merchant_id" efwdt:"merchant_id" binding:"omitempty,merchantExist"`
	UserId     entity.UserId       `json:"user_id" efwdt:"user_id" binding:"omitempty,userExist"`
	Action     types.LogUserAction `json:"action" efwdt:"action" binding:"omitempty,logUserActionExist"`
	Url        entity.Url          `json:"url" efwdt:"url" binding:"omitempty,gt=0"`
	efwdt.WithCreatedAt
}

func AdminLogUserList(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	merchantIds := []types.MerchantId{}
	if !isUserRootAdmin(adminUserModel.GetId()) {
		merchantModels := merchant.GetModels(merchant.QueryHelper{CreatorId: adminUserModel.GetId()})
		for _, merchantModel := range merchantModels {
			merchantIds = append(merchantIds, merchantModel.GetId())
		}
		if len(merchantIds) == 0 {
			merchantIds = append(merchantIds, 0)
		}
	}

	efwcontroller.Datable[loguser.Model, requestAdminLogUserList](
		context,
		[]string{"id", "user_id", "merchant_id", "action", "url", "created_at"},
		[]efwdt.ChainFilter[*loguser.Model]{
			func(chain gorm.ChainInterface[*loguser.Model]) gorm.ChainInterface[*loguser.Model] {
				if len(merchantIds) != 0 {
					return chain.Where("merchant_id IN (?)", merchantIds)
				}
				return chain
			},
		},
		nil,
		nil,
	)
}

func AdminLogUserListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "user_id", Type: types.EFWDT_FilterTypeUserId},
		{Name: "merchant_id", Type: types.EFWDT_FilterTypeMerchantId},
		{Name: "action", Type: types.EFWDT_FilterTypeList, Values: constants.LogUserActions},
		{Name: "url", Type: types.EFWDT_FilterTypeText},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}
