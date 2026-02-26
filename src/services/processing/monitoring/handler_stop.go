package monitoring

import (
	"compay/src/helpers"
	"compay/src/models/blockchainstate"
)

// Stop останавливает мониторинг
func (that *monitoring) Stop() {
	helpers.NotificationDebug("[Monitoring] Останавливаем...")
	if that.stopChan != nil {
		close(that.stopChan)
		that.stopChan = nil
	}

	if that.btc.blockchainStateModel != nil {
		helpers.NotificationDebug("[Monitoring] Выключаем btc ноду и сохраняем данные...")
		that.btc.blockchainStateModel.CurrentBlock = that.btc.currentBlock.Load()
		blockchainstate.Save(that.btc.blockchainStateModel, blockchainstate.QueryHelper{})
		that.btc.blockchainStateModel = nil
	}

	helpers.NotificationDebug("[Monitoring] Остановлен успешно")
}
