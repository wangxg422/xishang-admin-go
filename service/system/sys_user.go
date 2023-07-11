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

func (m *SysUserService) GetUserPage(params *sysDto.SysUserQueryDTO) (sysVo.PageResult, error) {
	pageResult := sysVo.PageResult{}

	db := global.DB.Model(&sysModel.SysUser{})

	likeArr := []string{
		"user_name",
		"phone_number",
	}

	utils.ConcatLikeWhereCondition(db, likeArr, params.UserName, params.PhoneNumber)
	utils.ConcatTimeRangeWhereCondition(db, params.BeginTime, params.EndTime)
	utils.ConcatOneEqualsStrWhereCondition(db, "status", params.Status)
	utils.ConcatOneEqualsInt8WhereCondition(db, enmu.DelFlagNormal.Name(), enmu.DelFlagNormal.Value())
	utils.ConcatOneEqualsInt64WhereCondition(db, "dept_id", params.DeptId)

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	limit, offset := params.PageInfo.Paging()
	db = db.Limit(limit).Offset(offset)

	var users []sysModel.SysUser
	res := db.Preload("SysDept").Find(&users)

	pageResult.Total = total
	pageResult.Rows = users

	return pageResult, res.Error
}
