package database

import "embed"

//go:embed migrations/*.sql
var MigrationsDir embed.FS
