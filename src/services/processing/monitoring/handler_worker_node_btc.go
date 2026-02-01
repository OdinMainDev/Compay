package monitoring

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/src/helpers"
	"compay/src/models/payment"
	"encoding/json"
	"fmt"
	"maps"
	"time"
)

func (that *monitoring) startWorkerBtcNode() {
	ticker := time.NewTicker(that.checkInterval)
	// Первая проверка сразу при старте
	that.handleBtcNode()

	for {
		select {
		case <-that.stopChan:
			fmt.Println("[Monitoring] STOP worker btc node")
			return
		case <-ticker.C:
			that.handleBtcNode()
		}
	}
}

func (that *monitoring) handleBtcNode() {
	fmt.Println("[Monitoring] Start handle btc node")

	that.paymentsMutex.RLock()
	activePayments := make(map[string]*payment.Model, len(that.btc.activePayments))
	maps.Copy(activePayments, that.btc.activePayments)
	that.paymentsMutex.RUnlock()

	if len(activePayments) == 0 {
		that.updateBtcBlockCount()
		return
	}

	for {
		postData, err := json.Marshal(getTransactionsBtcRequest{that.btc.blockchainStateModel.CurrentBlock})
		if err != nil {
			helpers.NotificationError(cperrors.ErrorMonitoringErrorOnGetBtcTransactionsMarshal(err.Error()).Error())
			time.Sleep(2 * time.Second)
			continue
		}
		var response getTransactionsBtcResponse
		if err := that.doRequest(requestGetTransactions, constants.BtcCodeAlpha, postData, &response); err != nil {
			helpers.NotificationError(cperrors.ErrorMonitoringErrorOnGetBtcTransactions(err.Error()).Error())
			time.Sleep(2 * time.Second)
			continue
		}
		that.parseBtcTransactions(&response, activePayments)
		return
	}

}

func (that *monitoring) updateBtcBlockCount() {
	fmt.Println("[Monitoring] Start update btc blocks node")

	var response getBlockCountBtcResponse
	err := that.doRequest(requestGetBlockCount, constants.BtcCodeAlpha, nil, &response)
	if err != nil {
		helpers.NotificationError(cperrors.ErrorMonitoringErrorOnUpdateBtcBlock(err.Error()).Error())
		time.Sleep(2 * time.Second)
		that.updateBtcBlockCount()
		return
	}

	that.btc.blockchainStateModel.CurrentBlock = response.Data
}

func (that *monitoring) parseBtcTransactions(response *getTransactionsBtcResponse, activePayments map[string]*payment.Model) {
	fmt.Println("[Monitoring] Start parse btc transactions")

	for _, tx := range response.Data.Transactions {
		if paymentModel, ok := activePayments[tx.Address]; !ok {
			continue
		}

	}
}
