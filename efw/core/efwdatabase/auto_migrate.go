package efwdatabase

import (
	"efw/config"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var ErrorMigrationsDirNil = fmt.Errorf("%s", "[EFW: DB MIGRATE] MigrationsDirNil")

var ErrorMigrationsIOFSFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: DB MIGRATE] MigrationsIOFSFail:", err)
}

var ErrorMigrationsConnectionDBFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: DB MIGRATE] MigrationsConnectionDBFail:", err)
}

var ErrorMigrationsPostgresFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: DB MIGRATE] MigrationsPostgresFail:", err)
}

var ErrorMigrationsFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: DB MIGRATE] MigrationsFail:", err)
}

func Migrate(Conn *gorm.DB, migrationsDir func() (embed.FS, string)) error {
	if migrationsDir == nil {
		return ErrorMigrationsDirNil
	}

	srcDriver, err := iofs.New(migrationsDir())
	if err != nil {
		return ErrorMigrationsIOFSFail(err.Error())
	}

	rawConn, err := Conn.DB()
	if err != nil {
		return ErrorMigrationsConnectionDBFail(err.Error())
	}

	driver, err := postgres.WithInstance(rawConn, &postgres.Config{})
	if err != nil {
		return ErrorMigrationsPostgresFail(err.Error())
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		srcDriver,
		config.DbName, driver)
	if err != nil {
		return ErrorMigrationsFail(err.Error())
	}

	m.Up()
	return nil
}
