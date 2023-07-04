package system

import (
	"backend/common/enmu"
	"backend/global"
	"backend/initial/logger"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

type SysUserService struct {
}

func (m *SysUserService) GetUserWithDept(id int64) (sysModel.SysUser, error) {
	user := &sysModel.SysUser{}
	res := global.DB.Preload("SysDept").
		Where("del_flag = ?", enmu.StatusNormal.Value()).
		First(&user, "user_id = ?", id)

	return *user, res.Error
}

func (m *SysUserService) GetUserInfo(id int64) (sysModel.SysUser, error) {
	user := &sysModel.SysUser{}
	res := global.DB.Preload("SysRoles").
		Preload("SysPosts").Preload("SysDept").
		First(&user, "user_id = ?", id).
		Where("del_flag = ?", enmu.StatusNormal.Value())

	return *user, res.Error
}

func (m *SysUserService) ChangePassword(c *gin.Context) {
}

func (m *SysUserService) CreateUser(userDto sysDto.SysCreateUserDTO) error {
	user := &sysModel.SysUser{}
	userDto.Convert(user)

	user.DelFlag = enmu.DelFlagNormal.Value()
	user.Status = enmu.StatusNormal.Value()
	user.Password = utils.EnPassword(user.Password)

	err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			logger.Error("create user failed", zap.Error(err))
			return err
		}

		if user.UserId == 0 {
			logger.Error("userId is null")
			return errors.New("userId is null")
		}

		if userDto.RoleIds != nil && len(userDto.RoleIds) > 0 {
			for _, id := range userDto.RoleIds {
				if err := tx.Create(&sysModel.SysUserRole{UserId: user.UserId, RoleId: id}).Error; err != nil {
					return err
				}
			}
		}

		if userDto.PostIds != nil && len(userDto.PostIds) > 0 {
			for _, id := range userDto.PostIds {
				if err := tx.Create(&sysModel.SysUserPost{UserId: user.UserId, PostId: id}).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	return err
}

func (m *SysUserService) UpdateUser(user *sysModel.SysUser) error {
	res := global.DB.Model(&sysModel.SysUser{UserId: user.UserId}).Updates(&user)
	return res.Error
}

func (m *SysUserService) DeleteUser(id int64) error {
	res := global.DB.Model(&sysModel.SysUser{UserId: id}).Update("del_flag", enmu.DelFlagDeleted.Value())
	return res.Error
}

func (m *SysUserService) ListUser() ([]sysModel.SysUser, error) {
	var list []sysModel.SysUser

	res := global.DB.Where("del_flag = ?", enmu.DelFlagNormal.Value()).Find(&list)

	return list, res.Error
}

func (m *SysUserService) GetUserById(id int64) (sysModel.SysUser, error) {
	user := sysModel.SysUser{
		UserId: id,
	}

	res := global.DB.Take(&user, id).Where("del_flag = ?", enmu.StatusNormal.Value())
	return user, res.Error
}

func (m *SysUserService) GetUserByUserName(userName string) (sysModel.SysUser, error) {
	user := sysModel.SysUser{}

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

	db := global.DB.Model(&sysModel.SysUser{})

	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	db = db.Limit(page.Limit).Offset(page.Offset).Where(condition)
	var userList []sysModel.SysUser
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
