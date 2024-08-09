package route

import (
	"exchang-go/route/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.GET("/rates", v1.GetRates).Use()
	api.GET("/exchange", v1.DoExchange)
}
