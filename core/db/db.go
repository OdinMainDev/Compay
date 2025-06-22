package db

import (
	"compay/core/config"
	"compay/core/eslog"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func CreateInstance() *gorm.DB {
	if config.DbDisabled {
		return nil
	}

	godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		config.DbHost, config.DbUser, config.DbPassword,
		config.DbName, config.DbPort, config.DbSSLMode, config.DbTimezone,
	)

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: &eslog.EsGormLogger{},
	}

	var err error
	eslog.Log("connecting to database")
	Conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), gormConfig)
	if err != nil {
		eslog.PanicLog("[PANIC] Error #1", err.Error())
	}
	sqlDB, err := Conn.DB()
	if err != nil {
		eslog.PanicLog("[PANIC] Error #2", err.Error())
	}

	sqlDB.SetMaxIdleConns(15)
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	return Conn
}
