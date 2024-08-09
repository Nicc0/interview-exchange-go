package middleware

import (
	"exchang-go/pkg/setting"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
)

func timeoutResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

func Timeout() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(setting.ServerSetting.ReadTimeout),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}
