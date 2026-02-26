package merchant

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/models/user"
	"efw/components/entity"
	"efw/components/model"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwcontroller"
	"efw/services/efwdt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type requestAdminUsersList struct {
	Id         entity.UserId    `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	MerchantId types.MerchantId `json:"merchant_id" efwdt:"merchant_id" binding:"omitempty,merchantExist"`
	CreatorId  entity.UserId    `json:"creator_id" efwdt:"creator_id" binding:"omitempty,userExist"`
	Username   string           `json:"username" efwdt:"username" binding:"omitempty,gt=0,max=255"`
	IsBanned   *bool            `json:"is_banned" efwdt:"is_banned" binding:"omitempty"`
	efwdt.WithCreatedAt
}

type requestAdminUsersCreate struct {
	Username   string           `json:"username" binding:"required,email,min=6,max=62"`
	Password   string           `json:"password" binding:"required,min=6,max=36"`
	IsAdmin    bool             `json:"is_admin"`
	MerchantId types.MerchantId `json:"merchant_id" binding:"required_if=IsAdmin false,merchantExist"`
}

type requestAdminUsersProcess struct {
	Id          entity.UserId `json:"id" binding:"required,gt=0"`
	NewUsername string        `json:"new_username" binding:"omitempty,email,min=6,max=255"`
	NewPassword string        `json:"new_password" binding:"omitempty,min=6,max=62"`
}

type requestAdminUsersAction struct {
	Id entity.UserId `json:"id" binding:"required,gt=0"`
}

func AdminUsersList(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.Datable[user.Model, requestAdminUsersList](
		context,
		[]string{"id", "username", "is_banned", "merchant_id", "creator_id", "created_at"},
		[]efwdt.ChainFilter[*user.Model]{
			func(chain gorm.ChainInterface[*user.Model]) gorm.ChainInterface[*user.Model] {
				return chain.Where("id NOT IN (?)", constants.AdminIds)
			},
			func(chain gorm.ChainInterface[*user.Model]) gorm.ChainInterface[*user.Model] {
				if !isUserRootAdmin(adminUserModel.GetId()) {
					return chain.Where("creator_id = ?", adminUserModel.GetId())
				}
				return chain
			},
		},
		efwdt.SpecialFilterMap{
			"username": func(fieldName string, value any) (efwdt.SpecialFilters, bool) {
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
		nil,
	)
}

func AdminUsersListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "merchant_id", Type: types.EFWDT_FilterTypeMerchantId},
		{Name: "creator_id", Type: types.EFWDT_FilterTypeUserId},
		{Name: "username", Type: types.EFWDT_FilterTypeText},
		{Name: "is_banned", Type: types.EFWDT_FilterTypeBoolean},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}

func AdminUsersGet(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.GetModel(
		context,
		user.GetModel,
		nil,
		func(id int64) user.QueryHelper {
			if isUserRootAdmin(entity.UserId(id)) {
				id = -1
			}

			qH := user.QueryHelper{Id: entity.UserId(id), DontCheckMerchantId: true}
			if !isUserRootAdmin(adminUserModel.GetId()) {
				qH.CreatorId = adminUserModel.GetId()
			}
			return qH
		},
	)
}

func AdminUsersCreate(context *gin.Context) {
	var input requestAdminUsersCreate
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	if input.IsAdmin && input.MerchantId != 0 {
		efwginhelpers.ResponseError(context, cperrors.ErrorUserCreateAdminWithMerchantIdIsNotAllowed)
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) && input.IsAdmin {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	if !input.IsAdmin && input.MerchantId == 0 {
		efwginhelpers.ResponseError(context, cperrors.ErrorUserCreateUserWithoutMerchantIdIsNotAllowed)
		return
	}

	checkUserExist := user.GetModel(user.QueryHelper{Username: input.Username, DontCheckMerchantId: true})
	if checkUserExist != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorUserUsernameIsTaken))
		return
	}

	userModel := &user.Model{}
	userModel.Username = input.Username
	userModel.MerchantId = input.MerchantId
	userModel.NewPassword = input.Password
	userModel.CreatorId = adminUserModel.GetId()
	user.Save(userModel, user.QueryHelper{})

	createLogUser(context, adminUserModel, constants.LogUserAdminUserCreate)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": userModel})
}

func AdminUsersProcess(context *gin.Context) {
	var input requestAdminUsersProcess
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	userModel := user.GetModel(user.QueryHelper{Id: input.Id, DontCheckMerchantId: true})
	if userModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if isUserRootAdmin(userModel.GetId()) || !isAdminCanUpdateUser(adminUserModel, userModel) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	if input.NewUsername != "" && userModel.Username != input.NewUsername {
		checkUserExist := user.GetModel(user.QueryHelper{Username: input.NewUsername, DontCheckMerchantId: true})
		if checkUserExist != nil {
			efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorUserUsernameIsTaken))
			return
		}
		userModel.Username = input.NewUsername
	}

	userModel.NewPassword = input.NewPassword
	user.Save(userModel, user.QueryHelper{})

	createLogUser(context, adminUserModel, constants.LogUserAdminUserProcess)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": userModel})
}

func AdminUsersSearch(context *gin.Context) {
	var input requestSearch
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	creatorId := entity.UserId(0)
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		creatorId = adminUserModel.GetId()
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": user.GetModels(user.QueryHelper{
		Search:              input.Search,
		CreatorId:           creatorId,
		DontCheckMerchantId: true,
		QueryHelper: model.QueryHelper{
			HasPagination: model.HasPagination{
				HasLimit: true,
			},
		},
	})})
}

func AdminUsersActionDisableTotp(context *gin.Context) {
	var input requestAdminUsersAction
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	userModel := user.GetModel(user.QueryHelper{Id: input.Id, DontCheckMerchantId: true})
	if userModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if isUserRootAdmin(userModel.GetId()) || !isAdminCanUpdateUser(adminUserModel, userModel) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	if userModel.IsTwofaEnabled {
		userModel.TotpToken = ""
		userModel.IsTwofaEnabled = false
		user.Save(userModel, user.QueryHelper{})
		createLogUser(context, adminUserModel, constants.LogUserAdminUserDisableTotp)
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": userModel})
}

func AdminUsersActionBan(context *gin.Context) {
	adminUsersActionBanOrNot(context, constants.LogUserAdminUserBan)
}

func AdminUsersActionUnban(context *gin.Context) {
	adminUsersActionBanOrNot(context, constants.LogUserAdminUserUnban)
}

func adminUsersActionBanOrNot(context *gin.Context, action types.LogUserAction) {
	var input requestAdminUsersAction
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	userModel := user.GetModel(user.QueryHelper{Id: input.Id, DontCheckMerchantId: true, WithMerchant: true})
	if userModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if isUserRootAdmin(userModel.GetId()) || !isAdminCanUpdateUser(adminUserModel, userModel) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	if userModel.MerchantId != 0 && (userModel.Merchant == nil || !userModel.Merchant.IsEnabled) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	actionSuccess := false

	if action == constants.LogUserAdminUserBan && !userModel.IsBanned {
		userModel.IsBanned = true
		actionSuccess = true
	}

	if action == constants.LogUserAdminUserUnban && userModel.IsBanned {
		userModel.IsBanned = false
		actionSuccess = true
	}

	if actionSuccess {
		user.Save(userModel, user.QueryHelper{})
		createLogUser(context, adminUserModel, action)
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": userModel})
}
