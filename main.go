package main

import (
	"backend/global"
	"backend/initial"
	"time"

	myzap "backend/initial/zap"

	"go.uber.org/zap"
)

func main() {
	global.Viper = initial.Viper()
	global.Log = myzap.InitZap()
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

	for {
		for i := 0; i < 1000; i++ {
			global.Log.Debug("i am debug")
			global.Log.Info("i am info")
			global.Log.Warn("i am warning")
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}
