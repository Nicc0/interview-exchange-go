package main

import (
	"exchang-go/middleware"
	"exchang-go/pkg/clients"
	"exchang-go/pkg/crypto"
	"exchang-go/pkg/setting"
	"exchang-go/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	setting.Setup()
}

func main() {
	gin.DisableConsoleColor()
	gin.SetMode(setting.ServerSetting.RunMode)

	addr := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(middleware.Timeout())

	route.InitRoutes(r)
	clients.InitClients()
	crypto.InitCryptoCurrencies()

	log.Printf("[info] start http server listening %s", addr)

	r.Run(addr)
}
