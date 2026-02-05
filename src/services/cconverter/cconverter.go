package cconverter

import (
	"compay/src/helpers"
	"sync"
	"time"
)

type CConverter struct {
	rates      map[string]map[string]float64
	lastUpdate time.Time
	mutex      sync.RWMutex
}

const (
	updateInterval = 30 * time.Second
	apiURL         = "https://compay.one/get-rates?apiKey=w1ZvhUJI89XvIeCMF3RXWmeBxArrSDWY"
)

var service *CConverter

func Start() {
	service = &CConverter{
		rates:      make(map[string]map[string]float64),
		lastUpdate: time.Time{},
	}
	service.Start()
}

// Start запускает фоновый процесс обновления курсов каждые 30 секунд
func (that *CConverter) Start() {
	if err := that.fetchRates(); err != nil {
		helpers.NotificationDebug("[CConverter] Запущен с ошибкой: " + err.Error())
		time.Sleep(1 * time.Second)
		that.Start()
		return
	}

	go that.work()
	helpers.NotificationDebug("[CConverter] Запущен успешно")
}

func (that *CConverter) work() {
	ticker := time.NewTicker(updateInterval)
	defer ticker.Stop()

	for range ticker.C {
		if err := that.fetchRates(); err != nil {
			helpers.NotificationError("[CConverter] Ошибка при обновлении курсов: " + err.Error())
		}
	}
}
