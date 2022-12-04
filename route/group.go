package route

import (
	"backend/route/public"
	"backend/route/system"
)

type AppRouteGroup struct {
	PublicRouteGroup public.AppPublicRouteGroup
	SysRouteGroup    system.SysRouteGroup
}

var AppRouteGroupIns = new(AppRouteGroup)
