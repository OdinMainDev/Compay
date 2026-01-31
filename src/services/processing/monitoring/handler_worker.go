package monitoring

import (
	"fmt"
)

// worker обрабатывает задания из очереди
func (that *Monitoring) worker(id int) {
	defer that.wg.Done()

	fmt.Printf("[Monitoring] Worker %d started\n", id)

	for {
		select {
		case <-that.stopChan:
			fmt.Printf("[Monitoring] Worker %d stopped\n", id)
			return
		case job := <-that.jobQueue:
			that.processPayment(job)
		}
	}
}
