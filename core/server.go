package core

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	"server/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	//初始化路由
	Router := initialize.Routers()

	Router.Static("/form-generator", "./resource/page")
	//拿到端口
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	//打开并保持服务
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
