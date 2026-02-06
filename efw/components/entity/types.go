package entity

type IdInt64 interface{ ~int64 }

type SaveType int // Тип сохранения
const (
	SAVE_TYPE_DB_CACHE SaveType = iota
	SAVE_TYPE_DB
	SAVE_TYPE_CACHE
	SAVE_TYPE_LOCAL_CACHE
)

// Базовые типы
type FingerPrint string
type Url string
type Ip string
type UserAgent string

// User Types
type UserId int64

// Amount Types
type Amount int64
type CurrencyCodeAlpha string
type CurrencyMinorUnits int
