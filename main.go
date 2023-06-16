package main

import (
	"backend/global"
	"backend/initial"

	"backend/initial/logger"

	"go.uber.org/zap"
)

func main() {
	global.Viper = initial.Viper()
	logger.InitZap()
	zap.ReplaceGlobals(global.Log)

	// db := initial.InitDb()
	// global.DB = db
	// defer func() {
	// 	if sqlDb, err := db.DB(); err != nil && sqlDb != nil {
	// 		sqlDb.Close()
	// 	}
	// }()

	global.Log.Debug("port is %s", zap.String("port", global.AppConfig.App.Port))
	logger.Info("port is %s", zap.String("port", global.AppConfig.App.Port))
	//initial.InitRoute()
}
