package entity

import (
	"database/sql/driver"
	"efw/utils"
	"fmt"
)

var ErrorEncryptedStringBadType = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: ENTITY] EncryptedStringBadType:", err)
}

type EncryptedString string

func (str EncryptedString) Value() (driver.Value, error) {
	if str == "" {
		return "", nil
	}
	return utils.EncryptString(string(str))
}

func (es *EncryptedString) Scan(value any) error {
	str := ""
	bytesArr, ok := value.([]byte)
	if ok {
		str = string(bytesArr)
	} else {
		strData, ok := value.(string)
		if !ok {
			return ErrorEncryptedStringBadType(fmt.Sprintf("%T", value))
		}
		str = strData
	}
	if str == "" {
		return nil
	}
	decrypted, err := utils.DecryptString(string(str))
	if err != nil {
		return err
	}
	*es = EncryptedString(decrypted)
	return nil
}
