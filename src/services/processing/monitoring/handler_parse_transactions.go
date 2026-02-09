package monitoring

import (
	"compay/common/types"
	"compay/src/models/payment"
	"fmt"
	"time"
)

func (that *monitoring) parseTransactions(transactions []types.TxData, activePayments map[types.PaymentAddress]*payment.Model) {
	that.log(fmt.Sprint("[Monitoring] Start parse transactions"))
	now := time.Now().Unix()

	for _, paymentModel := range activePayments {
		address := paymentModel.Address
		hasTx := false
		for _, txData := range transactions {
			if txData.Address != address {
				continue
			}

			hasTx = true
			if !that.isPaymentActive(paymentModel, now) {
				that.doSetPaymentExpired(paymentModel.GetId())
				continue
			}

			that.txParserJobChan <- txParserJob{
				PaymentModel: paymentModel,
				TxData:       &txData,
			}
		}

		if hasTx {
			continue
		}

		if !that.isPaymentActive(paymentModel, now) {
			that.doSetPaymentExpired(paymentModel.GetId())
			continue
		}
	}
}
