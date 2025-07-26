package database

import (
	"embed"
)

//go:embed migrations_wallet/*.sql
var MigrationsWalletDir embed.FS

//go:embed migrations_merchant/*.sql
var MigrationsMerchantDir embed.FS

var isWallet bool

func SetMigrationsType(wallet bool) {
	isWallet = wallet
}

func GetMigrationsDir() (embed.FS, string) {
	if isWallet {
		return MigrationsWalletDir, "migrations_wallet"
	}
	return MigrationsMerchantDir, "migrations_merchant"
}
