package internal

import (
	"fmt"
	"gorm.io/gorm/logger"
	"server/global"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	// 这一步实际上实在赋值
	return &writer{Writer: w}
}

func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		logZap = global.GVA_CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = global.GVA_CONFIG.Pgsql.LogZap
	}
	if logZap {
		global.GVA_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
