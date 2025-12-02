package efw

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

type eframework struct {
	options EfwOptions

	// esFile    *esfile.ESFile
	efwDB *gorm.DB
	// esClamav  *esclamav.ESClamav
	// esMaxmind *esmaxmind.ESMaxmind
}

type EfwInitMethods struct {
	InitLoggers func()
}

var eFrame *eframework

func Init(options EfwOptions, initMethods EfwInitMethods) error {
	fmt.Println("START EFW v1.0.1")

	os.Setenv("TZ", "UTC")

	if initMethods.InitLoggers != nil {
		initMethods.InitLoggers()
	}

	eFrame = &eframework{
		options: options,
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

	eFrame.printBanner()

	return nil
}
