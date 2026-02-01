package monitoring

func (that *monitoring) startWorkersNodes() {
	go that.startWorkerBtcNode()
}

func (that *monitoring) startWorkersJobs() {
	for i := range that.workerCount {
		go that.startWorkerJob(i + 1)
	}
}
