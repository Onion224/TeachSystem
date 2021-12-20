package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/config"
	"server/global"
	"server/initialize/internal"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	// 定义m对象指向yaml中配置的mysql数据库
	m := global.GVA_CONFIG.Mysql
	// 如果数据库名为空则返回nil
	if m.Dbname == "" {
		return nil
	}
	// 如果数据库已经被定义,则根据yaml中的配置定义具体的mysql对象
	// mysql.Config是gorm中对mysql数据库的配置
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), //网络连接数据源
		DefaultStringSize:         191,     // string类型字段的默认长度,和数据库设计中varchar的默认长度一致
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	//internal.Gorm.Config()用来gorm自定义配置
	//gorm利用Mysql的配置和gorm的配置来创建一个数据库db连接
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// 用传参的方式来传入配置初始化Mysql数据库
func GormMysqlByConfig(m config.DB) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
