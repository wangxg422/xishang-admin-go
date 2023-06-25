package system

import (
	"backend/common/enmu"
	"backend/global"
	sysModel "backend/model/system"

	"github.com/gin-gonic/gin"
)

type SysRoleService struct {
}

func (m *SysRoleService) ChangePassword(c *gin.Context) {
}

func (m *SysRoleService) CreateRole(Role *sysModel.SysRole) error {
	res := global.DB.Create(&Role)

	return res.Error
}

func (m *SysRoleService) UpdateRole(Role *sysModel.SysRole) error {
	res := global.DB.Model(&sysModel.SysRole{RoleId: Role.RoleId}).Updates(&Role)
	return res.Error
}

func (m *SysRoleService) DeleteRole(Roleid int64) error {
	res := global.DB.Model(&sysModel.SysRole{RoleId: Roleid}).Update("del_flag", enmu.EnmuGroupApp.DelFlagDelete.GetCode())
	return res.Error
}

func (m *SysRoleService) ListRole() ([]sysModel.SysRole, error) {
	var list []sysModel.SysRole

	res := global.DB.Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode()).Find(&list)

	return list, res.Error
}

func (m *SysRoleService) GetRoleById(Roleid int64) (sysModel.SysRole, error) {
	Role := sysModel.SysRole{
		RoleId: Roleid,
	}

	res := global.DB.Take(&Role, Roleid).Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode())
	return Role, res.Error
}
