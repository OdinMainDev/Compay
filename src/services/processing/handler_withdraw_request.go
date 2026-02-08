package processing

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/src/models/logpayment"
	"efw/components/entity"
	"efw/core/efwlog"
	"encoding/json"
)

type baseWithdrawRequest struct {
	Wif                    string        `json:"wif" binding:"required"`
	Address                string        `json:"address" binding:"required"`
	EstimatedFeeBlockCount int64         `json:"estimated_fee_block_count" binding:"required"`
	Amount                 entity.Amount `json:"amount" binding:"required,gt=0"`
	StoredUTXOs            []storedUTXO  `json:"stored_utxo" binding:"required"`
}

type storedUTXO struct {
	TxID   string `json:"tx_id"`
	Vout   uint32 `json:"vout_position"`
	Amount int64  `json:"amount"`
}

func (that *Processing) constructWithdrawRequestBody(wif, address string, amount entity.Amount, logPaymentModels []*logpayment.Model) ([]byte, error) {
	resp := []byte{}
	switch that.getCurrency().CodeAlpha {
	case constants.BtcCodeAlpha:
		storedUTXOs := []storedUTXO{}
		for _, logPaymentModel := range logPaymentModels {
			storedUTXOs = append(storedUTXOs, storedUTXO{
				TxID:   string(logPaymentModel.Data.TxData.TxId),
				Vout:   uint32(logPaymentModel.Data.TxData.VoutPosition),
				Amount: int64(logPaymentModel.Data.TxData.Amount),
			})
		}
		body, err := json.Marshal(baseWithdrawRequest{
			Wif:                    wif,
			Address:                address,
			EstimatedFeeBlockCount: that.getMerchantCurrency().EstimatedFeeBlockCount,
			Amount:                 amount,
			StoredUTXOs:            storedUTXOs,
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
