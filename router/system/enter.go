package system

import "backend/api/v1/system"

type SysRouterGroup struct {
	SysUserRouter
	SysDeptRouter
	SysRoleRouter
	SysPostRouter
	SysMenuRouter
	SysDictRouter
	SysConfigRouter
	SysModuleRouter
	SysPermissionRouter
}

var SysRouterGroupApp = new(SysRouterGroup)

var (
	userApi       = system.SysUserApi{}
	deptApi       = system.SysDeptApi{}
	roleApi       = system.SysRoleApi{}
	postApi       = system.SysPostApi{}
	menuApi       = system.SysMenuApi{}
	dictApi       = system.SysDictApi{}
	configApi     = system.SysConfigApi{}
	moduleApi     = system.SysModuleApi{}
	permissionApi = system.SysPermissionApi{}
)
