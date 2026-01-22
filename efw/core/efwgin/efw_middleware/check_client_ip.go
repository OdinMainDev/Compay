package efwmiddleware

import (
	"efw/core/efwgin/casts"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/core/efwmaxmind"
	"fmt"

	"github.com/gin-gonic/gin"
)

var ErrorMiddlewareInvalidIp = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: GIN-MIDDLEWARE] MiddlewareInvalidIp:", err)
}

var ErrorMiddlewareInvalidGeoliteCity = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: GIN-MIDDLEWARE] MiddlewareInvalidGeoliteCity:", err)
}

const GinIp = "_efw/gin/ip"
const GinCityData = "_efw/gin/ip_data"

func CheckClientIp() gin.HandlerFunc {
	return func(context *gin.Context) {
		ip, err := casts.GetClientIP(context)
		if err != nil {
			efwginhelpers.ResponseError(context, ErrorMiddlewareInvalidIp(err.Error()))
			context.Abort()
			return
		}

		city, err := efwmaxmind.CheckIp(ip)
		if err != nil {
			efwginhelpers.ResponseError(context, ErrorMiddlewareInvalidGeoliteCity(err.Error()))
			context.Abort()
			return
		}

		if city.Country.IsoCode == "" {
			efwginhelpers.ResponseError(context, ErrorMiddlewareInvalidGeoliteCity("InvalidCountry"))
			context.Abort()
			return
		}

		context.Set(GinCityData, city)
		context.Set(GinIp, ip)
		context.Next()
	}
}
