package logger

import (
	"backend/global"

	"go.uber.org/zap"
)

var Log = 

func InitZap() (logger *zap.Logger) {
	// 如果不存在日志目录，则新建
	if ok, _ := utils.PathExists(global.AppConfig.Zap.Path); !ok {
		fmt.Printf("create %v directory\n", global.AppConfig.Zap.Path)
		_ = os.Mkdir(global.AppConfig.Zap.Path, os.ModePerm)
	}

	cores := GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller())

	return logger
}


func Info(msg string, fields zap.Field) {
	nfo(msg, fields)
}
