package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jnhu76/dwz/models"
	"github.com/jnhu76/dwz/pkg/gredis"
	"github.com/jnhu76/dwz/pkg/logging"
	"github.com/jnhu76/dwz/pkg/setting"
	"github.com/jnhu76/dwz/pkg/util"
	"github.com/jnhu76/dwz/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title dwz server
// @version 1.0
// @description 短网址服务
// @securitydefinitions.apikey Bearer
// @in header
// @name token
// @termsOfService https://github.com/jnhu76/dwz
// @license.name MIT
// @license.url https://github.com/jnhu76/dwz/blob/master/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
