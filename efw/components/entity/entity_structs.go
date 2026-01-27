package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var ErrorStructScanFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: ENTITY] StructScanFail:", err)
}

var ErrorStructValueFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: ENTITY] StructValueFail:", err)
}

func ScanStruct[T any](data T, value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return ErrorStructScanFail(fmt.Sprint("ValueNotByteArray: ", value))
	}

	err := json.Unmarshal(bytes, data)
	if err != nil {
		return ErrorStructScanFail(err.Error())
	}
	return nil
}

func ValueStruct[T any](data T) (driver.Value, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", ErrorStructValueFail(err.Error())
	}
	return string(jsonBytes), nil
}
