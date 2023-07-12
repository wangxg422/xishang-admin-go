package system

import (
	"backend/common/enmu"
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
	"errors"
	"github.com/gin-gonic/gin"
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

func (m *SysUserService) CreateUser(userDto sysDto.SysUserCreateDTO) error {
	// 检查用户是否已存在，userName、userNumber保证唯一
	var existUser sysModel.SysUser
	err := global.DB.Where("(user_name = ? OR user_number = ?) AND del_flag = ?", userDto.UserName, userDto.UserNumber, enmu.DelFlagNormal.Value()).Find(&existUser).Limit(1).Error
	if err != nil {
		return err
	}

	if existUser.UserId != 0 {
		return errors.New("用户 " + userDto.UserNumber + "/" + userDto.UserNumber + " 已存在")
	}

	user := &sysModel.SysUser{}
	userDto.Convert(user)

	user.DelFlag = enmu.DelFlagNormal.Value()
	user.Status = enmu.StatusNormal.Value()
	user.Password = utils.EnPassword(user.Password)

	if userDto.RoleIds != nil && len(userDto.RoleIds) > 0 {
		var arr []sysModel.SysRole
		for _, id := range userDto.RoleIds {
			arr = append(arr, sysModel.SysRole{RoleId: id})
		}

		user.SysRoles = arr
	}

	if userDto.PostIds != nil && len(userDto.PostIds) > 0 {
		var arr []sysModel.SysPost
		for _, id := range userDto.PostIds {
			arr = append(arr, sysModel.SysPost{PostId: id})
		}

		user.SysPosts = arr
	}

	return global.DB.Create(user).Error
}

func (m *SysUserService) UpdateUser(userDto *sysDto.SysUserUpdateDTO) error {
	// 检查用户是否已存在，userName、userNumber保证唯一
	var existUser sysModel.SysUser
	err := global.DB.Where("(user_name = ? OR user_number = ?) AND del_flag = ?", userDto.UserName, userDto.UserNumber, enmu.DelFlagNormal.Value()).Find(&existUser).Limit(1).Error
	if err != nil {
		return err
	}

	if existUser.UserId == 0 {
		return errors.New("用户不存在")
	}

	if existUser.UserId != userDto.UserId {
		return errors.New("用户 " + userDto.UserNumber + "/" + userDto.UserNumber + " 已存在")
	}

	//paramMap := make(map[string]any)
	//paramMap["dept_id"] = userDto.DeptId
	//paramMap["user_name"] = userDto.UserName
	//paramMap["user_number"] = userDto.UserNumber
	//paramMap["real_name"] = userDto.RealName
	//paramMap["nick_name"] = userDto.NickName
	//paramMap["user_type"] = userDto.UserType
	//paramMap["email"] = userDto.Email
	//paramMap["phone_number"] = userDto.PhoneNumber
	//paramMap["sex"] = userDto.Sex
	//paramMap["status"] = userDto.Status
	//paramMap["remark"] = userDto.Remark
	//paramMap["update_time"] = userDto.UpdateTime
	//paramMap["update_by"] = userDto.UpdateBy
	var user sysModel.SysUser
	userDto.Convert(&user)

	if userDto.Password != "" {
		user.Password = utils.EnPassword(userDto.Password)
	}

	if userDto.RoleIds != nil && len(userDto.RoleIds) > 0 {
		var arr []sysModel.SysRole
		for _, id := range userDto.RoleIds {
			arr = append(arr, sysModel.SysRole{RoleId: id})
		}

		user.SysRoles = arr
	}

	if userDto.PostIds != nil && len(userDto.PostIds) > 0 {
		var arr []sysModel.SysPost
		for _, id := range userDto.PostIds {
			arr = append(arr, sysModel.SysPost{PostId: id})
		}

		user.SysPosts = arr
	}

	return global.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Model(&sysModel.SysUser{}).Omit("create_time", "create_by", "del_flag").
		Where("user_id = ? AND del_flag = ?", userDto.UserId, enmu.DelFlagNormal.Value()).
		Save(&user).Error
}

func (m *SysUserService) DeleteUser(id int64) error {
	res := global.DB.Model(&sysModel.SysUser{UserId: id}).Update("del_flag", enmu.DelFlagDeleted.Value())
	return res.Error
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
		"user_number",
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
