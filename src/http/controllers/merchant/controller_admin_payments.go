package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/currency"
	"compay/src/models/merchant"
	"compay/src/models/payment"
	"compay/src/models/user"
	"efw/core/efwgin"
	"efw/services/efwcontroller"
	"efw/services/efwdt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type requestAdminPaymentsList struct {
	Id                types.PaymentId         `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	TrxId             string                  `json:"trx_id" efwdt:"trx_id" binding:"omitempty,gt=0"`
	MerchantId        types.MerchantId        `json:"merchant_id" efwdt:"merchant_id" binding:"omitempty,merchantExist"`
	Currency          types.CurrencyCodeAlpha `json:"currency" efwdt:"currency" binding:"omitempty,currencyExist"`
	Status            types.PaymentStatus     `json:"status" efwdt:"status" binding:"omitempty,paymentStatusExist"`
	Address           types.PaymentAddress    `json:"address" efwdt:"address" binding:"omitempty,gt=0"`
	Amount            float64                 `json:"amount" efwdt:"amount" binding:"omitempty,amount"`
	AmountUsd         float64                 `json:"amount_usd" efwdt:"amount_usd" binding:"omitempty,amount"`
	ReceivedAmount    float64                 `json:"received_amount" efwdt:"received_amount" binding:"omitempty,amount"`
	ReceivedAmountUsd float64                 `json:"received_amount_usd" efwdt:"received_amount_usd" binding:"omitempty,amount"`
	efwdt.WithCreatedAt
	efwdt.WithExpiresAt
}

func AdminPaymentsList(context *gin.Context) {
	var input efwdt.Request[requestAdminPaymentsList]
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	var amountFilter = func(fieldName string, value any) (efwdt.SpecialFilters, bool) {
		v, ok := value.(float64)
		if !ok {
			return efwdt.SpecialFilters{}, false
		}
		amount := currency.FloatToAmountByCodeAlpha(v, input.Filters.Currency)
		return efwdt.SpecialFilters{
			Query: fieldName + " >= ?",
			Args:  []any{amount},
		}, true
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

	efwcontroller.Datable[payment.Model, requestAdminPaymentsList](
		context,
		[]string{"id", "trx_id", "merchant_id", "currency", "status", "address", "amount", "amount_usd", "received_amount", "received_amount_usd", "expires_at", "created_at"},
		[]efwdt.ChainFilter[*payment.Model]{
			func(chain gorm.ChainInterface[*payment.Model]) gorm.ChainInterface[*payment.Model] {
				if len(merchantIds) != 0 {
					return chain.Where("merchant_id IN (?)", merchantIds)
				}
				return chain
			},
		},
		efwdt.SpecialFilterMap{
			"amount":              amountFilter,
			"amount_usd":          amountFilter,
			"received_amount":     amountFilter,
			"received_amount_usd": amountFilter,
		},
		payment.DTOAdmin,
	)
}

func AdminPaymentsListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "trx_id", Type: types.EFWDT_FilterTypeText},
		{Name: "merchant_id", Type: types.EFWDT_FilterTypeMerchantId},
		{Name: "currency", Type: types.EFWDT_FilterTypeCurrency},
		{Name: "status", Type: types.EFWDT_FilterTypeList, Values: constants.PaymentStatuses},
		{Name: "address", Type: types.EFWDT_FilterTypeText},
		{Name: "amount", Type: types.EFWDT_FilterTypeNumber},
		{Name: "amount_usd", Type: types.EFWDT_FilterTypeNumber},
		{Name: "received_amount", Type: types.EFWDT_FilterTypeNumber},
		{Name: "received_amount_usd", Type: types.EFWDT_FilterTypeNumber},
		{Name: "expires_at", Type: types.EFWDT_FilterTypeDate},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}

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
