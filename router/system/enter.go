package system

import "backend/api/v1/system"

type SysRouterGroup struct {
	SysUserRouter
	SysDeptRouter
	SysRoleRouter
	SysPostRouter
	SysMenuRouter
}

var SysRouterGroupApp = new(SysRouterGroup)

var (
	userApi = system.SysUserApi{}
	deptApi = system.SysDeptApi{}
	roleApi = system.SysRoleApi{}
	postApi = system.SysPostApi{}
	menuApi = system.SysMenuApi{}
)
