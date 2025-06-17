package core

import (
	"compay/core/db"
	"compay/core/escache"
	"compay/core/esclamav"
	"compay/core/esfile"
	"compay/core/eslog"
	"compay/core/esmaxmind"
	"compay/core/esprofiler"
	"compay/core/rabbitmq"
	"compay/core/telegram"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type ESCore struct {
	esFile    *esfile.ESFile
	esDb      *gorm.DB
	esClamav  *esclamav.ESClamav
	esMaxmind *esmaxmind.ESMaxmind
}

type InitData struct {
	EnvCheckKeys []string
	EnvCallbacks []func()
}

var instance = &ESCore{}

func CreateInstance(initData InitData) {
	InitAndCheckEnv(initData.EnvCheckKeys, initData.EnvCallbacks)
	eslog.Initialize()
	fmt.Println("Configs loaded success")
	// возможно что стоит сделать чуть по другому
	// ( GET/SET методы в которых он будет подгружать класс уже на рантайме, чтобы не тратить время на запуск )
	instance.esFile = esfile.CreateInstance()
	fmt.Println("File uploader configured")
	instance.esDb = db.CreateInstance()
	fmt.Println("Database connection configured")
	instance.esClamav = esclamav.CreateInstance()
	fmt.Println("Clamav configured")
	escache.CreateInstance()
	fmt.Println("Cache configured")
	rabbitmq.Init()
	fmt.Println("RabbitMQ configured")
	instance.esMaxmind = esmaxmind.CreateInstance()
	fmt.Println("Maxmind database configured")
	telegram.PlatformBotInit()
	fmt.Println("Telegram bot configured")
	db.Migrate(instance.esDb)
	fmt.Println("Database migration success")

	esprofiler.InitInstance(context.Background())
	fmt.Println("Profiler instance created")
}

func GetInstance() *ESCore {
	return instance
}

func (that *ESCore) GetDB() *gorm.DB {
	return that.esDb
}

func (that *ESCore) SetDB(db *gorm.DB) {
	that.esDb = db
}

func (that *ESCore) GetFileUploader() *esfile.ESFile {
	return that.esFile
}

func (that *ESCore) GetClamav() *esclamav.ESClamav {
	return that.esClamav
}

func (that *ESCore) GetMaxMind() *esmaxmind.ESMaxmind {
	return that.esMaxmind
}
