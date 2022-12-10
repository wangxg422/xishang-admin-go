package system

import (
	"backend/config"
	"backend/global"
	model "backend/model/system"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (u *UserService) ChangePassword(c *gin.Context) {
}

func (u *UserService) CreateUser(user model.SysUser) error {
	res := global.DB.Create(&user)

	return res.Error
}

func (u *UserService) UpdateUser(user model.SysUser) error {
	res := global.DB.Model(&user).Updates(&user)
	return res.Error
}

func (u *UserService) DeleteUser(userid int64) error {
	res := global.DB.Model(&model.SysUser{UserId: userid}).Update("del_flag", config.UserDelFlagDelete)
	return res.Error
}

func (u *UserService) ListUser() ([]model.SysUser, error) {
	var list []model.SysUser

	res := global.DB.Where("del_flag = ?", config.UserDelFlagNormal).Find(&list)

	return list, res.Error
}

func (u *UserService) GetUserById(userid int64) (model.SysUser, error) {
	user := model.SysUser{
		UserId: userid,
	}

	res := global.DB.Take(&user, userid).Where("del_flag = ?", config.UserStatusNormal)
	return user, res.Error
}
