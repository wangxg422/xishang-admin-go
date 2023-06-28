package main

import (
	"backend/global"
	"backend/initial"

	"backend/initial/logger"

	"go.uber.org/zap"
)

func main() {
	global.Viper = initial.Viper()
	global.Log = logger.InitZap()
	//zap.ReplaceGlobals()

	initial.Redis()

	db := initial.InitDb()
	global.DB = db
	defer func() {
		if sqlDb, err := db.DB(); err != nil && sqlDb != nil {
			sqlDb.Close()
		}
	}()

	initial.InitRouter()
	logger.Info("server run at port %s", zap.String("port", global.AppConfig.App.Port))
}
