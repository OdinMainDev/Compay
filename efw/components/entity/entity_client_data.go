package entity

import (
	"database/sql/driver"

	"github.com/oschwald/geoip2-golang"
)

type ClientData struct {
	Fingerprint FingerPrint  `json:"fingerprint"`
	Ip          Ip           `json:"ip"`
	UserAgent   UserAgent    `json:"user_agent"`
	IpData      *geoip2.City `json:"ip_data"`
}

func (data *ClientData) Scan(value any) error        { return ScanStruct(data, value) }
func (data ClientData) Value() (driver.Value, error) { return ValueStruct(data) }
