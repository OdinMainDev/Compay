package monitoring

import (
	"fmt"
	"log"
)

// Start запускает мониторинг
func (that *Monitoring) Start() {
	fmt.Printf("[Monitoring] Starting invoice monitoring service with %d workers...\n", that.workerCount)

	// Подгружаем активные платежи
	that.loadPayments()

	// Запускаем воркеры
	for i := 0; i < that.workerCount; i++ {
		that.wg.Add(1)
		go that.worker(i)
	}

	// Запускаем диспетчер заданий
	that.wg.Add(1)
	go that.jobDispatcher()

	log.Println("[Monitoring] Invoice monitoring service started successfully")
}

// Stop останавливает мониторинг
func (that *Monitoring) Stop() {
	log.Println("[Monitoring] Stopping invoice monitoring service...")
	close(that.stopChan)
	that.wg.Wait()

	log.Println("[Monitoring] Invoice monitoring service stopped")
}
