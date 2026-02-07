package processing

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"efw/components/entity"
	"efw/core/efwlog"
	"encoding/json"
)

type baseWithdrawRequest struct {
	Wif                    string        `json:"wif" binding:"required"`
	Address                string        `json:"address" binding:"required"`
	EstimatedFeeBlockCount int64         `json:"estimated_fee_block_count" binding:"required"`
	Amount                 entity.Amount `json:"amount" binding:"required,gt=0"`
}

func (that *Processing) constructWithdrawRequestBody(wif, address string, amount entity.Amount) ([]byte, error) {
	resp := []byte{}
	switch that.getCurrency().CodeAlpha {
	case constants.BtcCodeAlpha:
		body, err := json.Marshal(baseWithdrawRequest{
			Wif:                    wif,
			Address:                address,
			EstimatedFeeBlockCount: that.getMerchantCurrency().EstimatedFeeBlockCount,
			Amount:                 amount,
		})
		if err != nil {
			return nil, cperrors.ErrorApiV1ErrorOnConstructWithdrawRequest(err.Error())
		}
		resp = body

	default:
		efwlog.Panic("BAD CONFIGURATION constructWithdrawRequestBody: ", that.getCurrency().CodeAlpha)
	}

	return resp, nil
}
