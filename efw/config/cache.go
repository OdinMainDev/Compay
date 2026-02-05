package config

import "time"

var (
	RedisHost     = "127.0.0.1:6379"
	RedisUsername = ""
	RedisPassword = ""

	CacheShardCount         = 16
	CacheOperationQueueSize = 1024
	CacheOperationTimeout   = 30 * time.Second
)
