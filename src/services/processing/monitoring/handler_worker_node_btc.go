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

	that.btc.mutex.RLock()
	activePayments := make(map[string]*payment.Model, len(that.btc.activePayments))
	maps.Copy(activePayments, that.btc.activePayments)
	that.btc.mutex.RUnlock()

	if len(activePayments) == 0 {
		that.updateBtcBlockCount()
		return
	}

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
		that.btc.currentBlock.Store(response.Data.NewHeight)
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

	that.btc.currentBlock.Store(response.Data)
}
