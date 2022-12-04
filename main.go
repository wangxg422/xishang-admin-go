package main

import (
	"backend/global"
	"backend/initial"
	"go.uber.org/zap"
)

func main() {
	global.Viper = initial.Viper()
	global.Log = initial.Zap()
	zap.ReplaceGlobals(global.Log)

	db := initial.InitDb()
	global.DB = db
	defer func() {
		if sqlDb, err := db.DB(); err != nil && sqlDb != nil {
			sqlDb.Close()
		}
	}()

	initial.InitRoute()
}
