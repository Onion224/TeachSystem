package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//初始化服务器的服务，让服务器提供指定端口和路由的服务
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
