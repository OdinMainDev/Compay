package esmaxmind

import (
	"compay/core/eslog"
	"net"

	"github.com/oschwald/geoip2-golang"
)

type ESMaxmind struct {
	db *geoip2.Reader
}

func CreateInstance() *ESMaxmind {
	var err error
	db, err := geoip2.Open("files/GeoLite2-City.mmdb")
	if err != nil {
		eslog.PanicLog(err.Error())
	}

	return &ESMaxmind{
		db: db,
	}
}

func (that *ESMaxmind) CheckIp(ipStr string) *geoip2.City {
	ip := net.ParseIP(ipStr)
	record, err := that.db.City(ip)
	if err != nil {
		eslog.LogError(err.Error())
	}
	return record
}
