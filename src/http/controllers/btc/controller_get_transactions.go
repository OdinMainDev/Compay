package btc

import (
	"compay/common/cperrors"
	"compay/src/services/cryptocurrencies/btc"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

type requestGetTransactions struct {
	Address string `json:"address" binding:"required"`
}

type transactionResponse struct {
	TxID          string  `json:"txid"`
	Amount        float64 `json:"amount"`
	AmountSatoshi int64   `json:"amount_satoshi"`
	Confirmations int     `json:"confirmations"`
	Time          int64   `json:"time"`
}

type responseGetTransactions struct {
	Address      string                `json:"address"`
	Transactions []transactionResponse `json:"transactions"`
	Count        int                   `json:"count"`
}

// GetTransactions получает список транзакций для адреса
func GetTransactions(context *gin.Context) {
	var input requestGetTransactions
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	service := btc.NewBTCService()
	transactions, err := service.GetAddressTransactions(input.Address)
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorBtcFailedGetTransactions(err.Error())))
		return
	}

	var txResponses []transactionResponse
	for _, tx := range transactions {
		txResponses = append(txResponses, transactionResponse{
			TxID:          tx.TxID,
			Amount:        tx.Amount,
			AmountSatoshi: int64(btc.BTCToSatoshi(tx.Amount)),
			Confirmations: tx.Confirmations,
			Time:          tx.Time,
		})
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": responseGetTransactions{
		Address:      input.Address,
		Transactions: txResponses,
		Count:        len(txResponses),
	}})
}
