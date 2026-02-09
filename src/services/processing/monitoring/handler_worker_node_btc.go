package monitoring

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/helpers"
	"compay/src/models/payment"
	"encoding/json"
	"fmt"
	"maps"
	"time"
)

func (that *monitoring) startWorkerBtcNode() {
	ticker := time.NewTicker(that.checkNodeInterval)
	// Первая проверка сразу при старте
	that.handleBtcNode()

	for {
		select {
		case <-that.stopChan:
			that.log(fmt.Sprint("[Monitoring] STOP worker btc node"))
			return
		case <-ticker.C:
			that.handleBtcNode()
		}
	}
}

func (that *monitoring) handleBtcNode() {
	that.log(fmt.Sprint("[Monitoring] Start handle btc node"))

	that.btc.mutex.RLock()
	lenActivePayments := len(that.btc.activePayments)
	hasActivePayments := lenActivePayments > 0
	that.btc.mutex.RUnlock()

	if !hasActivePayments {
		that.updateBtcBlockCount()
		return
	}

	activePayments := make(map[types.PaymentAddress]*payment.Model, lenActivePayments)
	that.btc.mutex.RLock()
	maps.Copy(activePayments, that.btc.activePayments)
	that.btc.mutex.RUnlock()

	for {
		currentBlock := that.btc.currentBlock.Load()
		postData, err := json.Marshal(getTransactionsBtcRequest{currentBlock})
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
		that.parseTransactions(response.Data.Transactions, activePayments)
		if response.Data.NewHeight > currentBlock {
			that.btc.currentBlock.Store(currentBlock + 1)
		}
		return
	}

}

func (that *monitoring) updateBtcBlockCount() {
	that.log(fmt.Sprint("[Monitoring] Start update btc blocks node"))

	var response getBlockCountBtcResponse
	err := that.doRequest(requestGetBlockCount, constants.BtcCodeAlpha, nil, &response)
	if err != nil {
		helpers.NotificationError(cperrors.ErrorMonitoringErrorOnUpdateBtcBlock(err.Error()).Error())
		time.Sleep(2 * time.Second)
		that.updateBtcBlockCount()
		return
	}

	currentBlock := that.btc.currentBlock.Load()
	if response.Data > currentBlock {
		that.btc.currentBlock.Store(currentBlock + 1)
	}
}
