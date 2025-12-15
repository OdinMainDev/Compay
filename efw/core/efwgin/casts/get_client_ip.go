package casts

import (
	"efw/config"
	"fmt"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

var ErrorNilClientIp = fmt.Errorf("%s", "[EFW: GIN-CASTS] NilClientIp")

func GetClientIP(context *gin.Context) (string, error) {
	forwardHeader := context.Request.Header.Get("X-Forwarded-For")
	firstAddress := strings.TrimSpace(strings.Split(forwardHeader, ",")[0])
	if net.ParseIP(firstAddress) != nil {
		return firstAddress, nil
	}

	clientIp := context.ClientIP()
	if ip := net.ParseIP(clientIp); ip != nil && !ip.IsLoopback() && !ip.IsUnspecified() {
		return clientIp, nil
	}

	if config.AppDebug {
		return "167.107.108.136", nil
	}

	return "", ErrorNilClientIp
}
