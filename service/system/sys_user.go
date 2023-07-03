package system

import (
	"backend/common/enmu"
	"backend/global"
	sysDto "backend/model/dto/system"
	"backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
	"errors"
	"strconv"

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

func (m *SysUserService) ListUserPage(c *gin.Context) (sysVo.PageResult, error) {
	pageResult := sysVo.PageResult{}

	page, err := sysDto.NewPageInfo(c)
	if err != nil {
		return pageResult, errors.New("分页信息获取失败")
	}

	condition := make(map[string]any)
	condition["del_flag"] = enmu.DelFlagNormal.Value()

	if c.Query("deptId") != "" {
		deptId, err := strconv.ParseInt(c.Query("deptId"), 10, 64)
		if err != nil {
			return pageResult, errors.New("dept id获取失败")
		}
		condition["dept_id"] = deptId
	}

	if c.Query("status") != "" {
		status, err := strconv.ParseInt(c.Query("status"), 10, 8)
		if err != nil {
			return pageResult, errors.New("status 解析失败")
		}
		condition["status"] = status
	}

	db := global.DB.Model(&system.SysUser{})

	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	db = db.Limit(page.Limit).Offset(page.Offset).Where(condition)
	var userList []system.SysUser
	if c.Query("userName") != "" {
		db.Where("user_name like ?", utils.AddPercentSign(c.Query("userName")))
	}

	if c.Query("phonenumber") != "" {
		db.Where("phonenumber like ?", utils.AddPercentSign(c.Query("phonenumber")))
	}

	if c.Query("beginTime") != "" && c.Query("endTime") != "" {
		db.Where("create_time >= ? AND create_time <= ?", c.Query("beginTime")+" 00:00:00", c.Query("endTime")+" 23:59:59")
	}

	err = db.Find(&userList).Error
	pageResult.List = userList
	pageResult.Total = total
	return pageResult, err
}
