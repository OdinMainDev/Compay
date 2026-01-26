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
	Username   string           `json:"username" efwdt:"username" binding:"omitempty,max=255"`
	IsBanned   *bool            `json:"is_banned" efwdt:"is_banned" binding:"omitempty"`
	efwdt.WithCreatedAt
}

type requestAdminUsersCreate struct {
	Username   string           `json:"username" binding:"required,email,min=6,max=255"`
	Password   string           `json:"password" binding:"required,min=6,max=255"`
	IsAdmin    bool             `json:"is_admin"`
	MerchantId types.MerchantId `json:"merchant_id" binding:"required_if=IsAdmin false,merchantExist"`
}

type requestAdminUsersProcess struct {
	Id          entity.UserId `json:"id" binding:"required,gt=0"`
	NewUsername string        `json:"new_username" binding:"omitempty,email,min=6,max=255"`
	NewPassword string        `json:"new_password" binding:"omitempty,min=6,max=255"`
}

type requestAdminUsersAction struct {
	Id entity.UserId `json:"id" binding:"required,gt=0"`
}

func AdminUsersList(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.Datable[user.Model, requestAdminUsersList](
		context,
		[]string{"id", "username", "is_banned", "merchant_id", "created_at"},
		[]efwdt.ChainFilter[user.Model]{
			func(chain gorm.ChainInterface[user.Model]) gorm.ChainInterface[user.Model] {
				return chain.Where("id NOT IN (?)", constants.AdminIds)
			},
			func(chain gorm.ChainInterface[user.Model]) gorm.ChainInterface[user.Model] {
				if !isUserRootAdmin(adminUserModel.GetId()) {
					return chain.Where("creator_id = ?", adminUserModel.GetId())
				}
				return nil
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
	)
}

func AdminUsersListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "merchant_id", Type: types.EFWDT_FilterTypeMerchantId},
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
				qH.CreatorId = entity.UserId(id)
			}
			return qH
		},
	)
}

func AdminUsersDelete(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	_, ok := efwcontroller.DeleteModel(
		context,
		user.GetModel,
		user.Delete,
		func(id int64) user.QueryHelper {
			qH := user.QueryHelper{Id: entity.UserId(id), DontCheckMerchantId: true}
			if !isUserRootAdmin(adminUserModel.GetId()) {
				qH.CreatorId = entity.UserId(id)
			}
			return qH
		},
	)
	if ok {
		createLogUser(context, adminUserModel, constants.LogUserAdminUserDelete)
	}
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
	if !input.IsAdmin && input.MerchantId == 0 {
		efwginhelpers.ResponseError(context, cperrors.ErrorUserCreateUserWithoutMerchantIdIsNotAllowed)
		return
	}

	checkUserExist := user.GetModel(user.QueryHelper{Username: input.Username, DontCheckMerchantId: true})
	if checkUserExist != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorUserUsernameIsTaken))
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)

	userModel := &user.Model{}
	userModel.Username = input.Username
	userModel.MerchantId = input.MerchantId
	userModel.NewPassword = input.Password
	userModel.CreatorId = adminUserModel.GetId()
	user.Save(userModel, user.QueryHelper{})

	createLogUser(context, userModel, constants.LogUserAdminUserCreate)
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

	createLogUser(context, userModel, constants.LogUserAdminUserProcess)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": userModel})
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

	if isUserRootAdmin(userModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	if userModel.IsTwofaEnabled {
		userModel.TotpToken = ""
		userModel.IsTwofaEnabled = false
		user.Save(userModel, user.QueryHelper{})
		createLogUser(context, userModel, constants.LogUserAdminUserDisableTotp)
	}

	efwginhelpers.ResponseSuccess(context, nil)
}

func AdminUsersSearch(context *gin.Context) {
	var input requestSearch
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": user.GetModels(user.QueryHelper{
		Search: input.Search,
		QueryHelper: model.QueryHelper{
			HasPagination: model.HasPagination{
				HasLimit: true,
			},
		},
	})})
}
