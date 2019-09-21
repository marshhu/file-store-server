package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/file-store-server/handler/handlers"
	"github.com/marshhu/file-store-server/handler/middleware"
	"io"
	"os"
)

func InitRouter() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.New()

	//静态文件
	r.Static("/upload", "./upload")

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// 允许使用跨域请求  全局中间件
	r.Use(middleware.Cors())

	api := r.Group("/api")
	{
		api.GET("/ping", handlers.PingHandler)
		api.POST("/uploadSingle", handlers.UploadSingleHandler)
		api.POST("/uploadMulti", handlers.UploadMultiHandler)
	}
	return r
}
