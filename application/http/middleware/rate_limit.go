package middleware

import (
	"bytes"
	"compay/application/helpers"
	"compay/application/http/casts"
	"compay/core/eslog"
	"io"
	"net/http"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

func keyFunc(context *gin.Context) string {
	var bodyBytes []byte
	var err error
	if context.Request.Body != nil {
		bodyBytes, err = io.ReadAll(context.Request.Body)
		if err != nil {
			eslog.LogError("RATELIMIT MIDDLEWARE ERROR: ", err.Error())
			helpers.GinErrorResponse(context, "InternalError")
			context.Abort()
			return ""
		}
	}
	context.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	defer func() {
		context.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}()

	clientIp := casts.GetClientIP(context)
	clientUserAgent := casts.GetClientUserAgent(context)

	key := clientIp
	if clientUserAgent == "" {
		helpers.GinErrorResponse(context, "BadUserAgent")
		context.Abort()
		return ""
	}
	key += "-" + clientUserAgent

	return key
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
