package processing

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/src/models/logtransaction"
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

func (that *Processing) constructWithdrawRequestBody(wif, address string, amount entity.Amount, logTransactionModels []*logtransaction.Model) ([]byte, error) {
	resp := []byte{}
	storedUTXOs := []storedUTXO{}
	for _, logTransactionModel := range logTransactionModels {
		if logTransactionModel.Data.TxData.TxId != "" {
			storedUTXOs = append(storedUTXOs, storedUTXO{
				TxID:   string(logTransactionModel.Data.TxData.TxId),
				Vout:   uint32(logTransactionModel.Data.TxData.VoutPosition),
				Amount: int64(logTransactionModel.Data.TxData.Amount),
			})
		} else {
			storedUTXOs = append(storedUTXOs, storedUTXO{
				TxID:   string(logTransactionModel.Data.TxWithdrawData.TxId),
				Vout:   uint32(logTransactionModel.Data.TxWithdrawData.VoutPosition),
				Amount: int64(logTransactionModel.Data.TxWithdrawData.Amount),
			})
		}
	}

	switch that.getCurrency().CodeAlpha {
	case constants.BtcCodeAlpha:
		body, err := json.Marshal(baseWithdrawRequest{
			Wif:                    wif,
			Address:                address,
			EstimatedFeeBlockCount: that.getMerchantCurrency().EstimatedFeeBlockCount,
			Amount:                 amount,
			StoredUTXOs:            storedUTXOs,
		})
		if err != nil {
			return nil, cperrors.ErrorProcessingErrorOnConstructWithdrawRequest(err.Error())
		}
		resp = body

	default:
		efwlog.Panic("BAD CONFIGURATION constructWithdrawRequestBody: ", that.getCurrency().CodeAlpha)
	}

	return resp, nil
}
