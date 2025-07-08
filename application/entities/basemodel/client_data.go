package basemodel

import "github.com/oschwald/geoip2-golang"

type ClientData struct {
	Language    string       `json:"language"`
	Fingerprint string       `json:"fingerprint"`
	Ip          string       `json:"ip"`
	UserAgent   string       `json:"user_agent"`
	Request     string       `json:"request"`
	IpData      *geoip2.City `json:"ip_data"`
}
