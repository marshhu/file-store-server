package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/file-store-server/handler"
	"github.com/marshhu/file-store-server/middleware"
	"net/http"
)
func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// 允许使用跨域请求  全局中间件
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

    api := r.Group("/api")
    {
		api.POST("/uploadSingle",handler.UploadSingleHandler)
		api.POST("/uploadMulti",handler.UploadMultiHandler)
	}

	r.Run(":8080")
}
