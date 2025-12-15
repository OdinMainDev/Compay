package efwmiddleware

import (
	"efw/core/efwgin/casts"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"fmt"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

var ErrorMiddlewareBadUserAgent = fmt.Errorf("%s", "[EFW: GIN-MIDDLEWARE] BadUserAgent")
var ErrorMiddlewareTooManyRequests = fmt.Errorf("%s", "[EFW: GIN-MIDDLEWARE] TooManyRequests")

type RateLimitRequest struct {
	Email string `json:"email" binding:"omitempty,email,max=255"`
	Phone string `json:"phone" binding:"omitempty,phone,max=255"`
}

func keyFunc(context *gin.Context) string {
	clientIp, err := casts.GetClientIP(context)
	if err != nil {
		efwginhelpers.ResponseError(context, err)
		context.Abort()
		return ""
	}

	clientUserAgent := casts.GetClientUserAgent(context)
	if clientUserAgent == "" {
		efwginhelpers.ResponseError(context, ErrorMiddlewareBadUserAgent)
		context.Abort()
		return ""
	}

	key := clientIp + "-" + clientUserAgent

	var input RateLimitRequest
	if err := context.ShouldBindBodyWithJSON(&input); err != nil {
		return key
	}

	if input.Email != "" {
		key += "-" + input.Email
	}

	if input.Phone != "" {
		key += "-" + input.Phone
	}

	return key
}

func errorHandler(context *gin.Context, _ ratelimit.Info) {
	efwginhelpers.ResponseError(context, ErrorMiddlewareTooManyRequests)
	context.Abort()
}

func RateLimit(count uint) gin.HandlerFunc {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute,
		Limit: count,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
	return mw
}
