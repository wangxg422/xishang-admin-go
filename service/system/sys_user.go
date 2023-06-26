package system

import (
	"backend/common/enmu"
	"backend/global"
	sysModel "backend/model/system"

	"github.com/gin-gonic/gin"
)

type SysUserService struct {
}

func (m *SysUserService) ChangePassword(c *gin.Context) {
}

func (m *SysUserService) CreateUser(user *sysModel.SysUser) error {
	res := global.DB.Create(&user)

	return res.Error
}

func (m *SysUserService) UpdateUser(user *sysModel.SysUser) error {
	res := global.DB.Model(&sysModel.SysUser{UserId: user.UserId}).Updates(&user)
	return res.Error
}

func (m *SysUserService) DeleteUser(id int64) error {
	res := global.DB.Model(&sysModel.SysUser{UserId: id}).Update("del_flag", enmu.EnmuGroupApp.DelFlagDelete.GetCode())
	return res.Error
}

func (m *SysUserService) ListUser() ([]sysModel.SysUser, error) {
	var list []sysModel.SysUser

	res := global.DB.Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode()).Find(&list)

	return list, res.Error
}

func (m *SysUserService) GetUserById(id int64) (sysModel.SysUser, error) {
	user := sysModel.SysUser{
		UserId: id,
	}

	res := global.DB.Take(&user, id).Where("del_flag = ?", enmu.EnmuGroupApp.StatusNormal)
	return user, res.Error
}
