package monitoring

import (
	"compay/common/types"
	"compay/src/models/payment"
	"fmt"
	"time"
)

func (that *monitoring) parseTransactions(transactions []types.TxData, activePayments map[string]*payment.Model) {
	fmt.Println("[Monitoring] Start parse transactions")
	now := time.Now().Unix()

	for _, txData := range transactions {
		paymentModel, ok := activePayments[txData.Address]
		if !ok {
			if that.isPaymentActive(paymentModel, now) {
				continue
			}
			that.doSetPaymentExpired(paymentModel)
			continue
		}

		that.txParserJobChan <- txParserJob{
			PaymentModel: paymentModel,
			TxData:       &txData,
		}
	}
}
