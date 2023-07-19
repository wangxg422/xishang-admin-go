package system

import (
	"backend/common/constant"
	"backend/common/enmu"
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
)

type SysRoleService struct {
}

func (m *SysRoleService) GetRolesByUserId(id int64) ([]sysModel.SysRole, error) {
	var list []sysModel.SysRole
	res := global.DB.Model(&sysModel.SysRole{}).Where("del_flag = ?", enmu.DelFlagNormal.Value()).Where("user_id", id).Find(&list)

	return list, res.Error
}

func (m *SysRoleService) CreateRole(role *sysModel.SysRole) error {
	res := global.DB.Create(&role)

	return res.Error
}

func (m *SysRoleService) UpdateRole(role *sysModel.SysRole) error {
	res := global.DB.Model(&sysModel.SysRole{RoleId: role.RoleId}).Updates(&role)
	return res.Error
}

func (m *SysRoleService) DeleteRole(id int64) error {
	res := global.DB.Model(&sysModel.SysRole{RoleId: id}).Update("del_flag", enmu.DelFlagDeleted.Value())
	return res.Error
}

func (m *SysRoleService) GetRoleById(id int64) (sysModel.SysRole, error) {
	Role := sysModel.SysRole{
		RoleId: id,
	}

	res := global.DB.Take(&Role, id).Where("del_flag = ?", enmu.DelFlagNormal.Value())
	return Role, res.Error
}

func (m *SysRoleService) GetAllRole() ([]sysModel.SysRole, error) {
	var list []sysModel.SysRole
	res := global.DB.Where("del_flag = ?", enmu.DelFlagNormal.Value()).Find(&list)
	return list, res.Error
}

func (m *SysRoleService) GetRolePage(params *sysDto.SysRoleQueryDTO) (sysVo.PageResult, error) {
	pageResult := sysVo.PageResult{}

	db := global.DB.Model(&sysModel.SysRole{})

	likeArr := []string{
		"role_name",
		"role_code",
		"role_perms",
	}

	utils.ConcatLikeWhereCondition(db, likeArr, params.RoleName, params.RoleCode, params.RolePerms)
	utils.ConcatOneEqualsStrWhereCondition(db, "status", params.Status)
	utils.ConcatTimeRangeWhereCondition(db, params.BeginTime, params.EndTime)
	utils.ConcatOneEqualsInt8WhereCondition(db, constant.DelFlag, enmu.DelFlagNormal.Value())

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	limit, offset := params.PageInfo.Paging()
	db = db.Limit(limit).Offset(offset)

	var roles []sysModel.SysRole
	res := db.Find(&roles)

	pageResult.Total = total
	pageResult.Rows = roles

	return pageResult, res.Error
}
