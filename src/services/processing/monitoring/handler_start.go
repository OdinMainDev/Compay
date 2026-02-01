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
	that.startWorkersJobs()

	helpers.NotificationDebug("[Monitoring] Запущен успешно")
}
