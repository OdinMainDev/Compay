package monitoring

import (
	"compay/src/models/blockchainstate"
	"log"
)

// Stop останавливает мониторинг
func (that *monitoring) Stop() {
	log.Println("[Monitoring] Stopping invoice monitoring service...")
	if that.stopChan != nil {
		close(that.stopChan)
	}

	if that.btc.blockchainStateModel != nil {
		blockchainstate.Save(that.btc.blockchainStateModel, blockchainstate.QueryHelper{})
	}

	log.Println("[Monitoring] Invoice monitoring service stopped")
}
