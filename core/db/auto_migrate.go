package db

import (
	"compay/common/database"
	"compay/core/config"
	"compay/core/eslog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func Migrate(Conn *gorm.DB) {
	if config.DbDisabled {
		return
	}

	srcDriver, err := iofs.New(database.GetMigrationsDir())
	if err != nil {
		eslog.PanicLog("Error: " + err.Error())
	}
	rawConn, err := Conn.DB()
	if err != nil {
		eslog.PanicLog("Error: " + err.Error())
	}
	driver, err := postgres.WithInstance(rawConn, &postgres.Config{})
	if err != nil {
		eslog.PanicLog("Error: " + err.Error())
	}
	m, err := migrate.NewWithInstance(
		"iofs",
		srcDriver,
		config.DbName, driver)
	if err != nil {
		eslog.PanicLog("Error: " + err.Error())
	}
	m.Up()
}
