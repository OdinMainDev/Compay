package middleware

import (
	"compay/application/helpers"
	"compay/application/http/casts"
	"net/http"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

func keyFunc(context *gin.Context) string {
	return casts.GetClientIP(context)
}

func errorHandler(context *gin.Context, _ ratelimit.Info) {
	helpers.GinErrorResponse(context, "TooManyRequests", http.StatusTooManyRequests)
}

func CheckThrottling(count uint) gin.HandlerFunc {
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
