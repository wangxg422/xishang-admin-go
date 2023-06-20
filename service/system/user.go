package system

import (
	"backend/config"
	"backend/global"
	sysModel "backend/model/system"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (u *UserService) ChangePassword(c *gin.Context) {
}

func (u *UserService) CreateUser(user *sysModel.SysUser) error {
	res := global.DB.Create(&user)

	return res.Error
}

func (u *UserService) UpdateUser(user *sysModel.SysUser) error {
	res := global.DB.Model(&sysModel.SysUser{UserId: user.UserId}).Updates(&user)
	return res.Error
}

func (u *UserService) DeleteUser(userid int64) error {
	res := global.DB.Model(&sysModel.SysUser{UserId: userid}).Update("del_flag", config.UserDelFlagDelete)
	return res.Error
}

func (u *UserService) ListUser() ([]sysModel.SysUser, error) {
	var list []sysModel.SysUser

	res := global.DB.Where("del_flag = ?", config.UserDelFlagNormal).Find(&list)

	return list, res.Error
}

func (u *UserService) GetUserById(userid int64) (sysModel.SysUser, error) {
	user := sysModel.SysUser{
		UserId: userid,
	}

	res := global.DB.Take(&user, userid).Where("del_flag = ?", config.UserStatusNormal)
	return user, res.Error
}
