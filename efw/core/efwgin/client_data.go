package efwgin

import (
	"efw/components/entity"
	"efw/core/efwgin/casts"
	efwmiddleware "efw/core/efwgin/efw_middleware"
	"efw/core/efwlog"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

var ErrorNotImplementedMiddlewareCheckClientIp error = fmt.Errorf("%s", "[EFW: GIN] NotImplementedMiddlewareCheckClientIp")

func ConstructClientData(context *gin.Context) entity.ClientData {
	ip, ok := context.Get(efwmiddleware.GinIp)
	if !ok {
		efwlog.Panic(ErrorNotImplementedMiddlewareCheckClientIp.Error())
		return entity.ClientData{}
	}

	city, ok := context.Get(efwmiddleware.GinCityData)
	if !ok {
		efwlog.Panic(ErrorNotImplementedMiddlewareCheckClientIp.Error())
		return entity.ClientData{}
	}

	return entity.ClientData{
		Fingerprint: casts.GetClientFingerPrint(context),
		Ip:          ip.(entity.Ip),
		UserAgent:   casts.GetClientUserAgent(context),
		IpData:      city.(*geoip2.City),
	}
}
