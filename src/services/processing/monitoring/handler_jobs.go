package monitoring

import (
	"compay/src/models/payment"
	"fmt"
	"time"
)

// jobDispatcher периодически добавляет задания в очередь
func (that *Monitoring) jobDispatcher() {
	defer that.wg.Done()

	ticker := time.NewTicker(that.checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-that.stopChan:
			return
		case <-ticker.C:
			that.dispatchJobs()
		}
	}
}

// dispatchJobs отправляет задания воркерам
func (that *Monitoring) dispatchJobs() {
	that.mutex.RLock()
	jobs := make([]*payment.Model, 0, len(that.activePayments))
	for _, job := range that.activePayments {
		jobs = append(jobs, job)
	}
	that.mutex.RUnlock()

	if len(jobs) == 0 {
		return
	}

	fmt.Printf("[Monitoring] Dispatching %d payment jobs to workers\n", len(jobs))

	for _, job := range jobs {
		select {
		case <-that.stopChan:
			return
		case that.jobQueue <- job:
		}
	}
}
