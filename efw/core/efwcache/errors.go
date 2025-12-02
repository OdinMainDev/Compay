package efwcache

import "fmt"

var ErrorInvalidCountOfShards = fmt.Errorf("%s", "[EFW: CACHE] InvalidCountOfShards")
var ErrorCreateStorage = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: CACHE] CreateStorage: ", err)
}
var ErrorInvalidCacheType = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: CACHE] InvalidCacheType: ", err)
}
var ErrorStorageInit = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: CACHE] StorageInit: ", err)
}
var ErrorStorageGet = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: CACHE] StorageGet: ", err)
}
var ErrorStorageSet = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: CACHE] StorageSet: ", err)
}
var ErrorStorageDelete = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: CACHE] StorageDelete: ", err)
}
var ErrorStorageClear = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: CACHE] StorageClear: ", err)
}
var ErrorStorageClose = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: CACHE] StorageClose: ", err)
}
var ErrorNilStorage = fmt.Errorf("%s", "[EFW: CACHE] NilStorage")
var ErrorUnpackBadType = fmt.Errorf("%s", "[EFW: CACHE] UnpackBadType")
var ErrorCacheNotInitialized = fmt.Errorf("%s", "[EFW: CACHE] NotInitialized")
var ErrorInvalidOperation = fmt.Errorf("%s", "[EFW: CACHE] InvalidOperation")
