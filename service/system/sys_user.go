package system

import (
	"backend/common/enmu"
	"backend/global"
	"backend/model/system"

	"github.com/gin-gonic/gin"
)

type SysUserService struct {
}

func (m *SysUserService) GetUserInfoWithDeptRoles(id int64) (system.SysUser, error) {
	user := &system.SysUser{}
	res := global.DB.Model(&system.SysUser{}).Preload("SysRoles").Preload("SysDept").Take(&user, id).Where("del_flag = ?", enmu.EnmuGroupApp.StatusNormal)

	return *user, res.Error
}

func (m *SysUserService) ChangePassword(c *gin.Context) {
}

func (m *SysUserService) CreateUser(user *system.SysUser) error {
	res := global.DB.Create(&user)

	return res.Error
}

func (m *SysUserService) UpdateUser(user *system.SysUser) error {
	res := global.DB.Model(&system.SysUser{UserId: user.UserId}).Updates(&user)
	return res.Error
}

func (m *SysUserService) DeleteUser(id int64) error {
	res := global.DB.Model(&system.SysUser{UserId: id}).Update("del_flag", enmu.EnmuGroupApp.DelFlagDelete.GetCode())
	return res.Error
}

func (m *SysUserService) ListUser() ([]system.SysUser, error) {
	var list []system.SysUser

	res := global.DB.Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode()).Find(&list)

	return list, res.Error
}

func (m *SysUserService) GetUserById(id int64) (system.SysUser, error) {
	user := system.SysUser{
		UserId: id,
	}

	res := global.DB.Take(&user, id).Where("del_flag = ?", enmu.EnmuGroupApp.StatusNormal)
	return user, res.Error
}
