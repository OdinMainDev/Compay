package btc

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/payment"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	btcClient     *BTCNodeClient
	btcClientOnce sync.Once
)

// GetBTCClient возвращает singleton экземпляр BTC клиента
func GetBTCClient() *BTCNodeClient {
	btcClientOnce.Do(func() {
		nodeURL := os.Getenv("BTC_NODE_URL")
		username := os.Getenv("BTC_NODE_USERNAME")
		password := os.Getenv("BTC_NODE_PASSWORD")

		if nodeURL == "" {
			nodeURL = constants.BtcNodeUrl
		}

		btcClient = NewBTCNodeClient(nodeURL, username, password)
	})
	return btcClient
}

// BTCService сервис для работы с Bitcoin платежами
type BTCService struct {
	client *BTCNodeClient
}

// NewBTCService создает новый экземпляр сервиса
func NewBTCService() *BTCService {
	return &BTCService{
		client: GetBTCClient(),
	}
}

// ImportPaymentAddress импортирует адрес платежа в ноду для отслеживания
func (s *BTCService) ImportPaymentAddress(address string, paymentID string) error {
	label := fmt.Sprintf("payment_%s", paymentID)
	return s.client.ImportAddress(address, label, false)
}

// CheckAddressBalance проверяет баланс адреса
func (s *BTCService) CheckAddressBalance(address string, minConfirmations int) (float64, error) {
	return s.client.GetReceivedByAddress(address, minConfirmations)
}

// GetAddressTransactions получает транзакции для адреса
func (s *BTCService) GetAddressTransactions(address string) ([]AddressTransaction, error) {
	received, err := s.client.ListReceivedByAddress(0, true, false)
	if err != nil {
		return nil, fmt.Errorf("failed to list received by address: %w", err)
	}

	var result []AddressTransaction
	for _, r := range received {
		if r.Address == address {
			for _, txid := range r.TxIDs {
				txDetail, err := s.client.GetTransaction(txid)
				if err != nil {
					continue
				}

				for _, detail := range txDetail.Details {
					if detail.Address == address && detail.Category == "receive" {
						result = append(result, AddressTransaction{
							TxID:          txid,
							Amount:        detail.Amount,
							Confirmations: txDetail.Confirmations,
							Time:          txDetail.Time,
						})
					}
				}
			}
			break
		}
	}

	return result, nil
}

// AddressTransaction представляет транзакцию для адреса
type AddressTransaction struct {
	TxID          string
	Amount        float64
	Confirmations int
	Time          int64
}

// BTCToSatoshi конвертирует BTC в сатоши
func BTCToSatoshi(btc float64) types.Amount {
	return types.Amount(btc * 100000000)
}

// SatoshiToBTC конвертирует сатоши в BTC
func SatoshiToBTC(satoshi types.Amount) float64 {
	return float64(satoshi) / 100000000
}

// CheckPaymentStatus проверяет статус платежа
func (s *BTCService) CheckPaymentStatus(paymentModel *payment.Model) (*PaymentStatusInfo, error) {
	transactions, err := s.GetAddressTransactions(paymentModel.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to get address transactions: %w", err)
	}

	var totalReceived types.Amount
	var confirmedAmount types.Amount
	var hasPendingTxs bool

	for _, tx := range transactions {
		amount := BTCToSatoshi(tx.Amount)
		totalReceived += amount

		if tx.Confirmations >= 1 {
			confirmedAmount += amount
		} else {
			hasPendingTxs = true
		}
	}

	info := &PaymentStatusInfo{
		TotalReceived:    totalReceived,
		ConfirmedAmount:  confirmedAmount,
		RequiredAmount:   paymentModel.Amount,
		HasPendingTxs:    hasPendingTxs,
		Transactions:     transactions,
		TransactionCount: len(transactions),
	}

	// Определяем статус платежа
	now := time.Now().Unix()
	isExpired := now > paymentModel.ExpiresAt

	if paymentModel.Amount == 0 {
		// Режим Amount == 0: записываем все транзакции в течении 24ч
		if len(transactions) > 0 {
			info.SuggestedStatus = constants.PaymentStatusFundsReceived
		} else if isExpired {
			info.SuggestedStatus = constants.PaymentStatusExpired
		} else {
			info.SuggestedStatus = constants.PaymentStatusCreated
		}
	} else {
		// Режим Amount != 0: проверяем достижение необходимой суммы
		if confirmedAmount >= paymentModel.Amount {
			info.SuggestedStatus = constants.PaymentStatusCompleted
		} else if totalReceived >= paymentModel.Amount {
			info.SuggestedStatus = constants.PaymentStatusFundsReceived
		} else if totalReceived > 0 {
			info.SuggestedStatus = constants.PaymentStatusInsufficientFunds
		} else if isExpired {
			info.SuggestedStatus = constants.PaymentStatusExpired
		} else {
			info.SuggestedStatus = constants.PaymentStatusCreated
		}
	}

	info.ShouldContinueMonitoring = !isExpired && (paymentModel.Amount == 0 || confirmedAmount < paymentModel.Amount)

	return info, nil
}

// PaymentStatusInfo информация о статусе платежа
type PaymentStatusInfo struct {
	TotalReceived            types.Amount
	ConfirmedAmount          types.Amount
	RequiredAmount           types.Amount
	HasPendingTxs            bool
	Transactions             []AddressTransaction
	TransactionCount         int
	SuggestedStatus          types.PaymentStatus
	ShouldContinueMonitoring bool
}
