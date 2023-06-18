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
	//zap.ReplaceGlobals()

	// db := initial.InitDb()
	// global.DB = db
	// defer func() {
	// 	if sqlDb, err := db.DB(); err != nil && sqlDb != nil {
	// 		sqlDb.Close()
	// 	}
	// }()

	//initial.InitRoute()
	logger.Info("server run at port %s", zap.String("port", global.AppConfig.App.Port))
}
