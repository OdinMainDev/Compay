package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/merchantwallet"
	"compay/src/models/user"
	"efw/components/entity"
	"efw/core/efwamount"
	"efw/core/efwgin"
	"efw/services/efwcontroller"
	"efw/services/efwdt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type requestMerchantWalletsList struct {
	Id             types.PaymentId          `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	PaymentId      string                   `json:"payment_id" efwdt:"-" binding:"omitempty,gt=0"`
	Currency       entity.CurrencyCodeAlpha `json:"currency" efwdt:"currency" binding:"omitempty,currencyExist"`
	Address        types.PaymentAddress     `json:"address" efwdt:"address" binding:"omitempty,gt=0"`
	Amount         float64                  `json:"amount" efwdt:"amount" binding:"omitempty,amount"`
	ReceivedAmount float64                  `json:"received_amount" efwdt:"received_amount" binding:"omitempty,amount"`
	SentAmount     float64                  `json:"sent_amount" efwdt:"sent_amount" binding:"omitempty,amount"`
	efwdt.WithCreatedAt
}

func MerchantWalletsList(context *gin.Context) {
	var input efwdt.Request[requestMerchantWalletsList]
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	var amountFilter = func(fieldName string, value any) (efwdt.SpecialFilters, bool) {
		v, ok := value.(float64)
		if !ok {
			return efwdt.SpecialFilters{}, false
		}
		amount := efwamount.FloatToAmountByCodeAlpha(v, input.Filters.Currency)
		return efwdt.SpecialFilters{
			Query: fieldName + " >= ?",
			Args:  []any{amount},
		}, true
	}

	userModel := context.MustGet(constants.GinUser).(*user.Model)

	efwcontroller.Datable[merchantwallet.Model, requestAdminMerchantWalletsList](
		context,
		[]string{"id", "currency", "address", "amount", "received_amount", "sent_amount", "created_at"},
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
		efwdt.SpecialFilterMap{
			"amount":          amountFilter,
			"received_amount": amountFilter,
			"sent_amount":     amountFilter,
		},
		merchantwallet.DTOMerchant,
	)
}

func MerchantWalletsListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "payment_id", Type: types.EFWDT_FilterTypeText},
		{Name: "currency", Type: types.EFWDT_FilterTypeCurrency},
		{Name: "address", Type: types.EFWDT_FilterTypeText},
		{Name: "amount", Type: types.EFWDT_FilterTypeNumber},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}

func MerchantWalletsGet(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.GetModel(
		context,
		merchantwallet.GetModel,
		merchantwallet.DTOMerchant,
		func(id int64) merchantwallet.QueryHelper {
			return merchantwallet.QueryHelper{
				Id:          types.MerchantWalletId(id),
				MerchantId:  adminUserModel.MerchantId,
				WithPayment: true,
			}
		},
	)
}
