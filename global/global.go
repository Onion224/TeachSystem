package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/config"
	"server/utils/timer"
)

//定义全局组件对象

var (
	//数据库对象
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	//配置文件对象
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Server
	GVA_REDIS  *redis.Client
	// GVA_LOG    *oplogging.Logger
	GVA_LOG   *zap.Logger
	GVA_Timer timer.Timer = timer.NewTimerTask()
	// 用来做本地缓存的包
	BlackCache local_cache.Cache
)
