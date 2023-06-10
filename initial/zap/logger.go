package zap

import (
	"backend/global"
	"fmt"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

func (w *writer) Printf(message string, data ...interface{}) {
	logZap := global.AppConfig.Mysql.LogZap

	if logZap {
		global.Log.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
