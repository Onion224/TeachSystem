package main

import (
	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	//初始化Viper
	global.GVA_VP = core.Viper()
	//初始化zap日志库
	global.GVA_LOG = core.Zap()
	//初始化Gorm连接数据库
	global.GVA_DB = initialize.Gorm()
	//初始化定时器
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		//初始化表
		initialize.RegisterTables(global.GVA_DB)
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
