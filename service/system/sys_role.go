package system

import (
	"backend/common/enmu"
	"backend/global"
	sysModel "backend/model/system"
)

type SysRoleService struct {
}

func (m *SysRoleService) CreateRole(role *sysModel.SysRole) error {
	res := global.DB.Create(&role)

	return res.Error
}

func (m *SysRoleService) UpdateRole(role *sysModel.SysRole) error {
	res := global.DB.Model(&sysModel.SysRole{RoleId: role.RoleId}).Updates(&role)
	return res.Error
}

func (m *SysRoleService) DeleteRole(id int64) error {
	res := global.DB.Model(&sysModel.SysRole{RoleId: id}).Update("del_flag", enmu.EnmuGroupApp.DelFlagDelete.GetCode())
	return res.Error
}

func (m *SysRoleService) ListRole() ([]sysModel.SysRole, error) {
	var list []sysModel.SysRole

	res := global.DB.Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode()).Find(&list)

	return list, res.Error
}

func (m *SysRoleService) GetRoleById(id int64) (sysModel.SysRole, error) {
	Role := sysModel.SysRole{
		RoleId: id,
	}

	res := global.DB.Take(&Role, id).Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode())
	return Role, res.Error
}
