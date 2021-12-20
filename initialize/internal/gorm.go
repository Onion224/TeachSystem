package internal

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"server/global"
	"time"
)

var Gorm = new(_gorm)

type _gorm struct {
}

func (g *_gorm) Config() *gorm.Config {
	//DisableForeignKeyConstraintWhenMigrating: true表示禁用自动为关联创建外键约束
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	//gorm默认的logger配置
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	switch global.GVA_CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
