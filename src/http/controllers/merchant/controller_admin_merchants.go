package merchant

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/models/merchant"
	"compay/src/models/user"
	"efw/components/entity"
	"efw/components/model"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwcontroller"
	"efw/services/efwdt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type requestAdminMerchantsList struct {
	Id        types.MerchantId `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	CreatorId entity.UserId    `json:"creator_id" efwdt:"creator_id" binding:"omitempty,userExist"`
	OwnerId   entity.UserId    `json:"owner_id" efwdt:"owner_id" binding:"omitempty,userExist"`
	Name      string           `json:"name" efwdt:"name" binding:"omitempty,max=255"`
	IsEnabled *bool            `json:"is_enabled" efwdt:"is_enabled" binding:"omitempty"`
	efwdt.WithCreatedAt
}

type requestAdminMerchantsCreate struct {
	Name          string `json:"name" binding:"required,min=6,max=255"`
	AdminUsername string `json:"admin_username" binding:"required,email,min=6,max=255"`
	AdminPassword string `json:"admin_password" binding:"required,min=6,max=255"`
}

func AdminMerchantsList(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.Datable[merchant.Model, requestAdminMerchantsList](
		context,
		[]string{"id", "name", "creator_id", "owner_id", "is_enabled", "created_at"},
		[]efwdt.ChainFilter[merchant.Model]{
			func(chain gorm.ChainInterface[merchant.Model]) gorm.ChainInterface[merchant.Model] {
				if !isUserRootAdmin(adminUserModel.GetId()) {
					return chain.Where("creator_id = ?", adminUserModel.GetId())
				}
				return nil
			},
		},
		efwdt.SpecialFilterMap{
			"name": func(fieldName string, value any) (efwdt.SpecialFilters, bool) {
				v, ok := value.(string)
				if !ok {
					return efwdt.SpecialFilters{}, false
				}
				return efwdt.SpecialFilters{
					Query: fieldName + " ILIKE ?",
					Args:  []any{"%" + v + "%"},
				}, true
			},
		},
	)
}

func AdminMerchantsListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "creator_id", Type: types.EFWDT_FilterTypeUserId},
		{Name: "admin_id", Type: types.EFWDT_FilterTypeUserId},
		{Name: "name", Type: types.EFWDT_FilterTypeText},
		{Name: "is_enabled", Type: types.EFWDT_FilterTypeBoolean},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}

func AdminMerchantsGet(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.GetModel(
		context,
		merchant.GetModel,
		nil,
		func(id int64) merchant.QueryHelper {
			qH := merchant.QueryHelper{Id: types.MerchantId(id)}
			if !isUserRootAdmin(adminUserModel.GetId()) {
				qH.CreatorId = entity.UserId(id)
			}
			return qH
		},
	)
}

func AdminMerchantsCreate(context *gin.Context) {
	var input requestAdminMerchantsCreate
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	checkUserExist := user.GetModel(user.QueryHelper{Username: input.AdminUsername, DontCheckMerchantId: true})
	if checkUserExist != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorUserUsernameIsTaken))
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	token := uuid.NewString()

	userModel := &user.Model{}
	userModel.Username = input.AdminUsername
	userModel.NewPassword = input.AdminPassword
	userModel.CreatorId = adminUserModel.GetId()
	user.Save(userModel, user.QueryHelper{})

	merchantModel := &merchant.Model{}
	merchantModel.Name = input.Name
	merchantModel.Token = entity.EncryptedString(token)
	merchantModel.CreatorId = adminUserModel.GetId()
	merchantModel.OwnerId = userModel.GetId()
	merchantModel.IsEnabled = true
	merchant.Save(merchantModel, merchant.QueryHelper{})

	userModel.MerchantId = merchantModel.GetId()
	user.Save(userModel, user.QueryHelper{})

	createLogUser(context, userModel, constants.LogUserAdminMerchantCreate)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchantModel, "token": token})
}

func AdminMerchantsSearch(context *gin.Context) {
	var input requestSearch
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchant.GetModels(merchant.QueryHelper{
		Search: input.Search,
		QueryHelper: model.QueryHelper{
			HasPagination: model.HasPagination{
				HasLimit: true,
			},
		},
	})})
}
