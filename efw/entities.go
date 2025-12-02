package efw

import "embed"

type EfwOptions struct {
	// App configuration
	AppEnvVariables []string
	AppEnvCallbacks []func()

	// Database configuration
	DatabaseDisabled               bool
	DatabaseAutoMigrationsDisabled bool
	DatabaseMigrationsDir          func() (embed.FS, string)

	// Clamav configuration
	ClamavDisabled bool

	// Cache configuration
	CacheDisabled       bool
	RedisCacheDisabled  bool
	MemoryCacheDisabled bool

	// Generic Сache Configuration
	GenericCacheDisabled bool

	// RabbitMQ configuration
	RabbitMQDisabled bool

	// Notify configuration
	NotifyDisabled         bool
	NotifyTelegramDisabled bool
}
