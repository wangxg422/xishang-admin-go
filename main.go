package main

import (
	"backend/global"
	"backend/initial"

	myzap "backend/initial/zap"

	"go.uber.org/zap"
)

func main() {
	global.Viper = initial.Viper()
	global.Log = myzap.GetZap()
	zap.ReplaceGlobals(global.Log)

	// db := initial.InitDb()
	// global.DB = db
	// defer func() {
	// 	if sqlDb, err := db.DB(); err != nil && sqlDb != nil {
	// 		sqlDb.Close()
	// 	}
	// }()

	global.Log.Debug("port is " + global.AppConfig.App.Port)
	//initial.InitRoute()

	for i := 0; i < 10; i++ {
		global.Log.Debug("i am debug")
		global.Log.Info("i am info")
		global.Log.Warn("i am warning")
	}
}
