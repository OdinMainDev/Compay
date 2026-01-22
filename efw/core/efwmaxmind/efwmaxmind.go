package efwmaxmind

import (
	"efw/components/entity"
	"efw/core/efwlog"
	_ "embed"
	"net"

	"github.com/oschwald/geoip2-golang"
)

//go:embed assets/GeoLite2-City.mmdb
var geoLiteCity []byte

var db *geoip2.Reader

func CreateInstance() {
	var err error
	db, err = geoip2.FromBytes(geoLiteCity)
	if err != nil {
		efwlog.Panic(err.Error())
	}
}

func CheckIp(ipStr entity.Ip) (*geoip2.City, error) {
	ip := net.ParseIP(string(ipStr))
	city, err := db.City(ip)
	if err != nil {
		return nil, err
	}
	return city, nil
}
