package monitoring

import (
	"compay/common/types"
	"compay/src/models/blockchainstate"
	"compay/src/models/logpayment"
	"compay/src/models/payment"
	"sync"
	"time"
)

type txParserJob struct {
	PaymentModel *payment.Model
	TxData       *types.TxData
}

type currencyData struct {
	blockchainStateModel *blockchainstate.Model
	activePayments       map[string]*payment.Model
}

// Monitoring мониторит инвойсы по блокам
type monitoring struct {
	// интервал дефолтный когда начинать проверку
	checkInterval time.Duration
	// количество воркеров
	workerCount int

	// Данные валют (платежей)
	btc             currencyData
	paymentsMutex   sync.RWMutex
	txParserJobChan chan txParserJob

	// Данные логов платежей
	activeLogPayments  map[string]*logpayment.Model
	logPaymentsMutex   sync.RWMutex
	logPaymentsJobChan chan *logpayment.Model

	// Контроль жизненного цикла
	stopChan chan struct{}
}

var globalMonitor *monitoring

// newInvoiceMonitor создает новый экземпляр Monitoring
func newInvoiceMonitor() *monitoring {
	return &monitoring{
		checkInterval:      30 * time.Second, // Проверяем блоки каждые 30 секунд
		workerCount:        10,
		txParserJobChan:    make(chan txParserJob),
		activeLogPayments:  make(map[string]*logpayment.Model),
		paymentsMutex:      sync.RWMutex{},
		logPaymentsMutex:   sync.RWMutex{},
		logPaymentsJobChan: make(chan *logpayment.Model),
		btc: currencyData{
			blockchainStateModel: nil,
			activePayments:       make(map[string]*payment.Model),
		},
		stopChan: make(chan struct{}),
	}
}

// StartMonitoring запускает глобальный мониторинг
func StartMonitoring() {
	globalMonitor = newInvoiceMonitor()
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
	globalMonitor.addPaymentAsync(paymentModel)
}

// RemovePayment убирает платеж из пула ( для админки к примеру )
func RemovePayment(paymentModel *payment.Model) {
	globalMonitor.removePaymentAsync(paymentModel.Address, paymentModel.Currency)
}
