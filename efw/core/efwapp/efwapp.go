package efwapp

import (
	"efw/core/efwlog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var AppShutdownTimeout = 5 * time.Second

func CreateApp(appName string, startApps []func(), shutdownApps []func()) {
	for _, start := range startApps {
		go start()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	efwlog.Log("Preparing to stop: " + appName)
	var wg sync.WaitGroup

	for _, shutdown := range shutdownApps {
		wg.Add(1)
		shutdown := shutdown
		go func() {
			defer wg.Done()
			shutdown()
		}()
	}
	wg.Wait()
	efwlog.Log("Successfully stopped: " + appName)
}
