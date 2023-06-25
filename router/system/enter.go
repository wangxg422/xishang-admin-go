package system

import "backend/api/v1/system"

type SysRouterGroup struct {
	SysUserRouter
	SysDeptRouter
	SysRoleRouter
	SysPostRouter
}

var SysRouterGroupApp = new(SysRouterGroup)

var (
	userApi = system.SysUserApi{}
	deptApi = system.SysDeptApi{}
	roleApi = system.SysRoleApi{}
	postApi = system.SysPostApi{}
)
