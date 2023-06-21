package system

import "backend/api/v1/system"

type SysRouterGroup struct {
	SysUserRouter
	SysDeptRouter
}

var SysRouterGroupApp = new(SysRouterGroup)

var (
	userApi = system.SysUserApi{}
)
