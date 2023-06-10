package zap

import (
	"backend/global"
	"backend/utils"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
