package efwdatabase

import (
	"efw/config"
	"efw/core/efwlog"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ErrorDatabaseGormOpenFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: DB] GormOpenFail:", err)
}

var ErrorDatabaseConnenctionDBFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: DB] ConnenctionDBFail:", err)
}

func Init() (*gorm.DB, error) {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
			config.DbHost, config.DbUser, config.DbPassword,
			config.DbName, config.DbPort, config.DbSSLMode, config.DbTimezone,
		),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: &efwlog.EsGormLogger{},
	})

	if err != nil {
		return nil, ErrorDatabaseGormOpenFail(err.Error())
	}

	sqlDB, err := conn.DB()
	if err != nil {
		return nil, ErrorDatabaseConnenctionDBFail(err.Error())
	}

	sqlDB.SetMaxIdleConns(15)
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	return conn, nil
}
