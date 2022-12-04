package v1

import (
	"backend/api/v1/system"
)

type AppApiGroup struct {
	SysApiGroup system.SysApiGroup
}

var AppApiGroupIns = new(AppApiGroup)
