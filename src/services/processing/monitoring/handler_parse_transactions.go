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
			continue
		}

		if !that.isPaymentActive(paymentModel, now) {
			that.doSetPaymentExpired(paymentModel.GetId())
			continue
		}

		that.txParserJobChan <- txParserJob{
			PaymentModel: paymentModel,
			TxData:       &txData,
		}
	}
}
