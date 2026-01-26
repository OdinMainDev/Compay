package efw

import (
	"os"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type eframework struct {
	options     EfwOptions
	initMethods EfwInitMethods

	// esFile    *esfile.ESFile
	efwDB *gorm.DB
	// esClamav  *esclamav.ESClamav
	// esMaxmind *esmaxmind.ESMaxmind
}

type EfwInitMethods struct {
	InitLoggers    func()
	InitValidation func() [](func() (string, validator.Func))
}

var eFrame *eframework

func Init(options EfwOptions, initMethods EfwInitMethods) error {
	os.Setenv("TZ", "UTC")

	eFrame = &eframework{
		options:     options,
		initMethods: initMethods,
	}

	if err := eFrame.initLogger(); err != nil {
		return err
	}

	if err := eFrame.loadEnvironmentVariables(); err != nil {
		return err
	}

	if err := eFrame.initDatabase(); err != nil {
		return err
	}

	if err := eFrame.initCache(); err != nil {
		return err
	}

	if err := eFrame.initGenericCache(); err != nil {
		return err
	}

	if err := eFrame.initNotify(); err != nil {
		return err
	}

	if err := eFrame.initGin(); err != nil {
		return err
	}

	eFrame.printBanner()

	return nil
}
