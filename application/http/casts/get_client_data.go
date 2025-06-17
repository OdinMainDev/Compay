package casts

import (
	"compay/common/config"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetClientIP(c *gin.Context) string {
	forwardHeader := c.Request.Header.Get("X-Forwarded-For")
	firstAddress := strings.TrimSpace(strings.Split(forwardHeader, ",")[0])
	if net.ParseIP(firstAddress) != nil {
		return firstAddress
	}

	clientIp := c.ClientIP()
	if ip := net.ParseIP(clientIp); ip != nil && !ip.IsLoopback() && !ip.IsUnspecified() {
		return clientIp
	}

	if config.AppDebug {
		return "167.107.108.136"
	}
	return ""
}

func GetClientUserAgent(c *gin.Context) string {
	return c.GetHeader("User-Agent")
}
