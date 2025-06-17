package basemodel

import (
	"compay/core/utils"
	"database/sql/driver"
	"fmt"
)

type EncryptedString string

func (str EncryptedString) Value() (driver.Value, error) {
	if str == "" {
		return "", nil
	}
	return utils.EncryptString(string(str))
}

// TODO: ошибки в константы
func (es *EncryptedString) Scan(value interface{}) error {
	str := ""
	bytesArr, ok := value.([]byte)
	if ok {
		str = string(bytesArr)
	} else {
		strData, ok := value.(string)
		if !ok {
			return fmt.Errorf("%s %T", "BadFormat of data: ", value)
		}
		str = strData
	}
	if str == "" {
		return nil
	}
	decrypted, err := utils.DecryptString(string(str))
	if err != nil {
		return fmt.Errorf("decrypt failed: %w", err)
	}
	*es = EncryptedString(decrypted)
	return nil
}
