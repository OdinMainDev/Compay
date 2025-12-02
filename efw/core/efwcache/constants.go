package efwcache

// CacheType описывает типы кешей которые есть ( memory - в памяти | redis - в редисе | both - оба одновременно)
type CacheType int

const (
	CacheTypeMemory CacheType = iota
	CacheTypeRedis
	CacheTypeAllInOne
)

// OperationType описывает типы операций ( для thread safe запросов )
type OperationType int

const (
	OperationGet OperationType = iota
	OperationSet
	OperationUpdate
	OperationDelete
	OperationClear
)
