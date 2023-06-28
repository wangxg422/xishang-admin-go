package system

import (
	"backend/common/enmu"
	"backend/global"
	"backend/model/system"
)

type SysRoleService struct {
}

func (m *SysRoleService) GetRolesByUserId(id int64) ([]system.SysRole, error) {
	var list []system.SysRole
	res := global.DB.Model(&system.SysRole{}).Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode()).Where("user_id", id).Find(&list)

	return list, res.Error
}

func (m *SysRoleService) CreateRole(role *system.SysRole) error {
	res := global.DB.Create(&role)

	return res.Error
}

func (m *SysRoleService) UpdateRole(role *system.SysRole) error {
	res := global.DB.Model(&system.SysRole{RoleId: role.RoleId}).Updates(&role)
	return res.Error
}

func (m *SysRoleService) DeleteRole(id int64) error {
	res := global.DB.Model(&system.SysRole{RoleId: id}).Update("del_flag", enmu.EnmuGroupApp.DelFlagDelete.GetCode())
	return res.Error
}

func (m *SysRoleService) ListRole() ([]system.SysRole, error) {
	var list []system.SysRole

	res := global.DB.Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode()).Find(&list)

	return list, res.Error
}

func (m *SysRoleService) GetRoleById(id int64) (system.SysRole, error) {
	Role := system.SysRole{
		RoleId: id,
	}

	res := global.DB.Take(&Role, id).Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode())
	return Role, res.Error
}
