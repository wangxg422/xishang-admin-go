package service

import "backend/service/system"

type AppServiceGroup struct {
	SysServiceGroup system.SysServiceGroup
}

var AppServiceGroupIns = new(AppServiceGroup)
