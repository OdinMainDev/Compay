package monitoring

import (
	"compay/common/config"
	"compay/common/types"
	"compay/src/models/payment"
	"compay/src/services/cryptocurrencies/btc"
	"sync"
	"time"
)

type ClosePaymentChan struct {
	PaymentModel *payment.Model
	Status       types.PaymentStatus
	Amount       types.Amount
}

// Monitoring мониторит инвойсы в pool worker режиме
type Monitoring struct {
	btcService    *btc.BTCService
	checkInterval time.Duration
	workerCount   int

	// Ответ в процессинг
	closePaymentChan chan ClosePaymentChan

	// Очередь заданий
	jobQueue chan *payment.Model

	// Активные платежи (для быстрой проверки существования)
	activePayments map[int64]*payment.Model
	mutex          sync.RWMutex

	// Контроль жизненного цикла
	stopChan chan struct{}
	wg       sync.WaitGroup
}

var globalMonitor *Monitoring

// newInvoiceMonitor создает новый экземпляр Monitoring
func newInvoiceMonitor(closePaymentChan chan ClosePaymentChan) *Monitoring {
	return &Monitoring{
		btcService:       btc.NewBTCService(),
		checkInterval:    config.MonitoringCheckInterval,
		workerCount:      config.MonitoringWorkerCount,
		closePaymentChan: closePaymentChan,
		jobQueue:         make(chan *payment.Model, 1000), // Буфер для 1000 платежей
		activePayments:   make(map[int64]*payment.Model),
		stopChan:         make(chan struct{}),
	}
}

// StartMonitoring запускает глобальный мониторинг
func StartMonitoring(closePaymentChan chan ClosePaymentChan) {
	globalMonitor = newInvoiceMonitor(closePaymentChan)
	globalMonitor.Start()
}

// StopMonitoring останавливает глобальный мониторинг
func StopMonitoring() {
	if globalMonitor != nil {
		globalMonitor.Stop()
	}
}

// AddPayment добавляет платеж в пул мониторинга
func AddPayment(paymentModel *payment.Model) {
	globalMonitor.addPayment(paymentModel)
}

// RemovePayment убирает платеж из пула
func RemovePayment(paymentModel *payment.Model) {
	globalMonitor.removePayment(paymentModel.GetIdInt64())
}
