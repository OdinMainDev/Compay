package efw

import (
	"efw/core/efwdatabase"

	"gorm.io/gorm"
)

func (that *eframework) initDatabase() error {
	if that.options.DatabaseDisabled {
		return nil
	}

	db, err := efwdatabase.Init()
	if err != nil {
		return err
	}

	that.efwDB = db
	if that.options.DatabaseAutoMigrationsDisabled {
		return nil
	}

	return that.DBMigrate()
}

func (that *eframework) DBMigrate() error {
	if that.options.DatabaseDisabled {
		return nil
	}

	return efwdatabase.Migrate(that.efwDB, that.options.DatabaseMigrationsDir)
}

func DBTransaction[T any]() gorm.Interface[T] {
	return gorm.G[T](eFrame.efwDB)
}
