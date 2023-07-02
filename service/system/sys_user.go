package system

import (
	"backend/common/enmu"
	"backend/global"
	"backend/model/dto"
	"backend/model/system"
	sysVo "backend/model/vo/system"

	"github.com/gin-gonic/gin"
)

type SysUserService struct {
}

func (m *SysUserService) GetUserWithDept(id int64) (system.SysUser, error) {
	user := &system.SysUser{}
	res := global.DB.Preload("SysDept").
		Where("del_flag = ?", enmu.StatusNormal.Value()).
		First(&user, "user_id = ?", id)

	return *user, res.Error
}

func (m *SysUserService) GetUserInfo(id int64) (system.SysUser, error) {
	user := &system.SysUser{}
	res := global.DB.Preload("SysRoles").
		Preload("SysPosts").Preload("SysDept").
		First(&user, "user_id = ?", id).
		Where("del_flag = ?", enmu.StatusNormal.Value())

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
	res := global.DB.Model(&system.SysUser{UserId: id}).Update("del_flag", enmu.DelFlagDeleted.Value())
	return res.Error
}

func (m *SysUserService) ListUser() ([]system.SysUser, error) {
	var list []system.SysUser

	res := global.DB.Where("del_flag = ?", enmu.DelFlagNormal.Value()).Find(&list)

	return list, res.Error
}

func (m *SysUserService) GetUserById(id int64) (system.SysUser, error) {
	user := system.SysUser{
		UserId: id,
	}

	res := global.DB.Take(&user, id).Where("del_flag = ?", enmu.StatusNormal.Value())
	return user, res.Error
}

func (m *SysUserService) GetUserByUserName(userName string) (system.SysUser, error) {
	user := system.SysUser{}

	res := global.DB.Take(&user, "user_name = ?", userName).Where("del_flag = ?", enmu.StatusNormal.Value())
	return user, res.Error
}

func (m *SysUserService) ListUserPage(page dto.PageInfo, deptId int64) (sysVo.PageResult, error) {
	pageResult := sysVo.PageResult{}
	db := global.DB.Model(&system.SysUser{})

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	var userList []system.SysUser
	if deptId == 0 {
		err = db.Limit(page.Limit).Offset(page.Offset).
			Where("del_flag = ?", enmu.DelFlagNormal.Value()).
			Find(&userList).Error
	} else {
		err = db.Limit(page.Limit).Offset(page.Offset).
			Where("dept_id = ? AND del_flag = ?", deptId, enmu.DelFlagNormal.Value()).
			Find(&userList).Error
	}

	pageResult.List = userList
	pageResult.Total = total
	return pageResult, err
}
