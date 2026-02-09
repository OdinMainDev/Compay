package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/currency"
	"compay/src/models/merchant"
	"compay/src/models/merchantwallet"
	"compay/src/models/user"
	"compay/src/services/processing"
	"efw/components/entity"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwcontroller"
	"efw/services/efwdt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type requestMerchantWalletsList struct {
	Id        types.PaymentId          `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	PaymentId string                   `json:"payment_id" efwdt:"-" binding:"omitempty,gt=0"`
	Currency  entity.CurrencyCodeAlpha `json:"currency" efwdt:"currency" binding:"omitempty,currencyExist"`
	Address   types.PaymentAddress     `json:"address" efwdt:"address" binding:"omitempty,gt=0"`
	efwdt.WithCreatedAt
}

type requestMerchantWalletsActionWithdraw struct {
	Id      types.MerchantWalletId `json:"id" binding:"required,gt=0"`
	Address types.PaymentAddress   `json:"address" binding:"required,max=128"`
	Amount  float64                `json:"amount" binding:"required,gt=0"`
}

func MerchantWalletsList(context *gin.Context) {
	var input efwdt.Request[requestMerchantWalletsList]
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	userModel := context.MustGet(constants.GinUser).(*user.Model)

	efwcontroller.Datable[merchantwallet.Model, requestMerchantWalletsList](
		context,
		[]string{"id", "currency", "amount", "received_amount", "sent_amount", "created_at"},
		[]efwdt.ChainFilter[*merchantwallet.Model]{
			func(chain gorm.ChainInterface[*merchantwallet.Model]) gorm.ChainInterface[*merchantwallet.Model] {
				return chain.Where("merchant_id = ?", userModel.MerchantId)
			},
			func(chain gorm.ChainInterface[*merchantwallet.Model]) gorm.ChainInterface[*merchantwallet.Model] {
				if input.Filters.PaymentId == "" {
					return chain.Preload("Payment", func(db gorm.PreloadBuilder) error { return nil })
				}

				return chain.Joins(
					clause.JoinTarget{
						Association: "Payment",
						Table:       "payments",
					},
					func(db gorm.JoinBuilder, joinTable clause.Table, curTable clause.Table) error {
						db.Where(
							clause.Eq{
								Column: clause.Column{Table: joinTable.Name, Name: "trx_id"},
								Value:  input.Filters.PaymentId,
							},
						)
						return nil
					},
				).Preload("Payment", func(db gorm.PreloadBuilder) error { return nil })
			},
		},
		nil,
		merchantwallet.DTOMerchant,
	)
}

func MerchantWalletsListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "payment_id", Type: types.EFWDT_FilterTypeText},
		{Name: "currency", Type: types.EFWDT_FilterTypeCurrency},
		{Name: "address", Type: types.EFWDT_FilterTypeText},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}

func MerchantWalletsGet(context *gin.Context) {
	userModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.GetModel(
		context,
		merchantwallet.GetModel,
		merchantwallet.DTOMerchant,
		func(id int64) merchantwallet.QueryHelper {
			return merchantwallet.QueryHelper{
				Id:          types.MerchantWalletId(id),
				MerchantId:  userModel.MerchantId,
				WithPayment: true,
			}
		},
	)
}

func MerchantWalletsActionWithdraw(context *gin.Context) {
	var input requestMerchantWalletsActionWithdraw
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	userModel := context.MustGet(constants.GinUser).(*user.Model)

	merchantWalletModel := merchantwallet.GetModel(merchantwallet.QueryHelper{
		Id:         types.MerchantWalletId(input.Id),
		MerchantId: userModel.MerchantId,
	})
	if merchantWalletModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}
	merchantModel := context.MustGet(constants.GinMerchant).(*merchant.Model)
	currencyModel := currency.GetModel(currency.QueryHelper{
		CodeAlpha: merchantWalletModel.Currency,
	})

	handler := processing.NewService(merchantModel, currencyModel)
	err := handler.StartWithdraw(processing.WithdrawData{
		Address:        input.Address,
		Amount:         input.Amount,
		MerchantWallet: merchantWalletModel,
	})
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(err))
		return
	}

	createLogUser(context, userModel, constants.LogUserMerchantWithdraw)
	efwginhelpers.ResponseSuccess(context, nil)
}
