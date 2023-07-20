package system

import (
	"backend/common/constant"
	"backend/common/enmu"
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
	"github.com/pkg/errors"
	"strconv"
)

type SysRoleService struct {
}

func (m *SysRoleService) GetRolesByUserId(id int64) ([]sysModel.SysRole, error) {
	var list []sysModel.SysRole
	res := global.DB.Model(&sysModel.SysRole{}).Where("del_flag = ?", enmu.DelFlagNormal.Value()).Where("user_id", id).Find(&list)

	return list, res.Error
}

func (m *SysRoleService) CreateRole(roleDto *sysDto.SysRoleCreateDTO) error {
	var existRole sysModel.SysRole
	err := global.DB.
		Where("role_code = ? AND del_flag = ?", roleDto.RoleCode, enmu.DelFlagNormal.Value()).
		Find(&existRole).Limit(1).Error
	if err != nil {
		return err
	}

	if existRole.RoleId != 0 {
		return errors.New("角色编码 " + roleDto.RoleCode + " 已经存在")
	}

	role := roleDto.Convert()
	role.DelFlag = enmu.DelFlagNormal.Value()

	// 角色关联的菜单列表
	if roleDto.MenuIds != nil && len(roleDto.MenuIds) > 0 {
		var list []sysModel.SysMenu
		for _, m := range roleDto.MenuIds {
			if m != "" {
				id, err := strconv.ParseInt(m, 10, 64)
				if err != nil {
					return err
				}
				list = append(list, sysModel.SysMenu{MenuId: id})
			}
		}
		role.SysMenus = list
	}

	// 如果自定义数据权限
	if roleDto.DataScope == "2" && roleDto.DeptIds != nil && len(roleDto.DeptIds) > 0 {
		var list []sysModel.SysDept
		for _, m := range roleDto.DeptIds {
			if m != "" {
				id, err := strconv.ParseInt(m, 10, 64)
				if err != nil {
					return err
				}
				list = append(list, sysModel.SysDept{DeptId: id})
			}
		}
		role.SysDepts = list
	}
	
	return global.DB.Create(&role).Error
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
