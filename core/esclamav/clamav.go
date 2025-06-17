package esclamav

import (
	"compay/core/config"
	"compay/core/eslog"
	"mime/multipart"

	"github.com/dutchcoders/go-clamd"
)

type ESClamav struct {
	Clamav *clamd.Clamd
}

func CreateInstance() *ESClamav {
	if config.ClamavSkipValidation {
		return &ESClamav{}
	}

	return &ESClamav{
		Clamav: clamd.NewClamd(config.ClamavUrl),
	}
}

func (esc *ESClamav) Validate(file multipart.File) bool {
	if config.ClamavSkipValidation {
		return true
	}

	response, err := esc.Clamav.ScanStream(file, make(chan bool))

	if err != nil {
		eslog.LogError(err.Error())
		return false
	}

	result := <-response
	return result.Status != "FOUND"
}
