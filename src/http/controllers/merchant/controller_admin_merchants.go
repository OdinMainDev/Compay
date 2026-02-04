package merchant

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/models/currency"
	"compay/src/models/merchant"
	"compay/src/models/merchantcurrency"
	"compay/src/models/user"
	"efw/components/entity"
	"efw/components/model"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwcontroller"
	"efw/services/efwdt"
	"net/http"

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
	Name          string `json:"name" binding:"required,min=6,max=50"`
	AdminUsername string `json:"admin_username" binding:"required,email,min=6,max=255"`
	AdminPassword string `json:"admin_password" binding:"required,min=6,max=255"`
}

type requestAdminMerchantsProcess struct {
	Id                                     types.MerchantId `json:"id" binding:"required,gt=0"`
	Name                                   string           `json:"name" binding:"required,min=6,max=50"`
	NotificationTelegramChatId             int64            `json:"notification_telegram_chat_id" binding:"omitempty"`
	NotificationDisableCreateInvoice       bool             `json:"notification_disable_create_invoice"`
	NotificationDisableDetectedTransaction bool             `json:"notification_disable_detected_transaction"`
	NotificationDisableFundsRecieved       bool             `json:"notification_disable_funds_recieved"`
	NotificationDisableCompletedInvoice    bool             `json:"notification_disable_completed_invoice"`
	NotificationDisableExpireInvoice       bool             `json:"notification_disable_expire_invoice"`
	NotificationDisableInsufficientFunds   bool             `json:"notification_disable_insufficient_funds"`
}

type requestAdminMerchantsAction struct {
	Id types.MerchantId `json:"id" binding:"required,gt=0"`
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
				return chain
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
		{Name: "owner_id", Type: types.EFWDT_FilterTypeUserId},
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
				qH.CreatorId = adminUserModel.GetId()
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
	merchantModel.Token = token
	merchantModel.CreatorId = adminUserModel.GetId()
	merchantModel.OwnerId = userModel.GetId()
	merchantModel.IsEnabled = true
	merchant.Save(merchantModel, merchant.QueryHelper{})

	userModel.MerchantId = merchantModel.GetId()
	user.Save(userModel, user.QueryHelper{})

	go asyncAdminMerchantsCreate(context.Copy(), adminUserModel, merchantModel)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchantModel})
}

func asyncAdminMerchantsCreate(context *gin.Context, adminUserModel *user.Model, merchantModel *merchant.Model) {
	currenyModels := currency.GetModels(currency.QueryHelper{
		IsEnabled: true,
	})
	for _, currenyModel := range currenyModels {
		merchantCurrencyModel := merchantcurrency.GetModel(merchantcurrency.QueryHelper{
			MerchantId: merchantModel.GetId(),
			Currency:   currenyModel.CodeAlpha,
		})

		if merchantCurrencyModel == nil {
			merchantCurrencyModel = merchantcurrency.DefaultCurrency(currenyModel.CodeAlpha, merchantModel.GetId(), merchantModel.CreatorId)
			merchantcurrency.Save(merchantCurrencyModel, merchantcurrency.QueryHelper{})
		}
	}

	createLogUser(context, adminUserModel, constants.LogUserAdminMerchantCreate)
}

func AdminMerchantsProcess(context *gin.Context) {
	var input requestAdminMerchantsProcess
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	merchantModel := merchant.GetModel(merchant.QueryHelper{Id: input.Id})
	if merchantModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) && !isUserCanUpdateMerchant(adminUserModel, merchantModel) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	merchantModel.Name = input.Name
	// Public
	merchantModel.Settings.NotificationSettings.TelegramChatId = input.NotificationTelegramChatId
	merchantModel.Settings.NotificationSettings.DisableCreateInvoice = input.NotificationDisableCreateInvoice
	merchantModel.Settings.NotificationSettings.DisableDetectedTransaction = input.NotificationDisableDetectedTransaction
	merchantModel.Settings.NotificationSettings.DisableFundsRecieved = input.NotificationDisableFundsRecieved
	merchantModel.Settings.NotificationSettings.DisableCompletedInvoice = input.NotificationDisableCompletedInvoice
	merchantModel.Settings.NotificationSettings.DisableExpireInvoice = input.NotificationDisableExpireInvoice
	merchantModel.Settings.NotificationSettings.DisableInsufficientFunds = input.NotificationDisableInsufficientFunds

	merchant.Save(merchantModel, merchant.QueryHelper{})

	createLogUser(context, adminUserModel, constants.LogUserAdminMerchantProcess)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchantModel})
}

func AdminMerchantsSearch(context *gin.Context) {
	var input requestSearch
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	creatorId := entity.UserId(0)
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		creatorId = adminUserModel.GetId()
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchant.GetModels(merchant.QueryHelper{
		Search:    input.Search,
		CreatorId: creatorId,
		QueryHelper: model.QueryHelper{
			HasPagination: model.HasPagination{
				HasLimit: true,
			},
		},
	})})
}

func AdminMerchantsActionDisable(context *gin.Context) {
	adminMerchantsActionDisableOrNot(context, constants.LogUserAdminMerchantDisable)
}

func AdminMerchantsActionEnable(context *gin.Context) {
	adminMerchantsActionDisableOrNot(context, constants.LogUserAdminMerchantEnable)
}

func adminMerchantsActionDisableOrNot(context *gin.Context, action types.LogUserAction) {
	var input requestAdminMerchantsAction
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	merchantModel := merchant.GetModel(merchant.QueryHelper{Id: input.Id})
	if merchantModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) && !isUserCanUpdateMerchant(adminUserModel, merchantModel) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	actionSuccess := false

	if action == constants.LogUserAdminMerchantDisable && merchantModel.IsEnabled {
		merchantModel.IsEnabled = false
		user.BanMerchantUsers(merchantModel.GetId())
		actionSuccess = true
	}

	if action == constants.LogUserAdminMerchantEnable && !merchantModel.IsEnabled {
		merchantModel.IsEnabled = true
		user.UnbanMerchantUser(merchantModel.OwnerId, merchantModel.GetId())
		actionSuccess = true
	}

	if actionSuccess {
		merchant.Save(merchantModel, merchant.QueryHelper{})
		createLogUser(context, adminUserModel, action)
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchantModel})
}
