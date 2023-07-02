package system

import "backend/service/system"

var (
	userService   = system.SysUserService{}
	deptService   = system.SysDeptService{}
	roleService   = system.SysRoleService{}
	postService   = system.SysPostService{}
	menuService   = system.SysMenuService{}
	dictService   = system.SysDictService{}
	configService = system.SysConfigService{}
)
