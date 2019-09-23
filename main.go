package main

import (
"fmt"
"github.com/gin-gonic/gin"
"github.com/marshhu/file-store-server/conf"
"github.com/marshhu/file-store-server/handler/router"
"net/http"
)

func main() {
	gin.SetMode(conf.ServerSetting.RunMode)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.ServerSetting.HttpPort),
		Handler:        router.InitRouter(),
		ReadTimeout:    conf.ServerSetting.ReadTimeout,
		WriteTimeout:   conf.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
