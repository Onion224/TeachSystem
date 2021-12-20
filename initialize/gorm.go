package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"server/global"
	"server/model"
)

// 	Gorm初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

//RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := global.GVA_DB.AutoMigrate(
		model.User{},
		model.Teacher{},
		model.Student{},
		model.Score{},
		model.Course{},
		model.Class{},
		model.CoursesPlan{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
