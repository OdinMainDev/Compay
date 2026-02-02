package monitoring

import (
	"compay/src/helpers"
)

// Start запускает мониторинг
func (that *monitoring) Start() {
	helpers.NotificationDebug("[Monitoring] Запускается...")

	// Подгружаем состояния блокчейнов
	that.loadBlockchainStates()

	helpers.NotificationDebug("[Monitoring] Загружаем платежи...")
	// Подгружаем платежи
	that.loadPayments()
	helpers.NotificationDebug("[Monitoring] Загружаем платежи 2...")
	// Запускаем воркеров нод
	that.startWorkersNodes()

	// Запускаем воркеров обработчиков
	that.startWorkersTxParser()

	// Запускаем воркеров обработчиков логов платежей ( проверки конфирмов )
	that.startWorkersLogChecker()

	helpers.NotificationDebug("[Monitoring] Запущен успешно")
}

func (that *monitoring) startWorkersNodes() {
	go that.startWorkerBtcNode()
}

func (that *monitoring) startWorkersTxParser() {
	for i := range that.workerCount {
		go that.startWorkerTxParser(i + 1)
	}
}

func (that *monitoring) startWorkersLogChecker() {
	for i := range that.workerCount {
		go that.startWorkerLogChecker(i + 1)
	}
	go that.startLoopLogChecker()
}
