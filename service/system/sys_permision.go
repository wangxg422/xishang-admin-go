package system

import (
	"backend/global"
	sysModel "backend/model/system"
	"errors"
)

type SysPermissionService struct {
}

func (m SysPermissionService) DeletePermission(id int64) error {
	// TODO 满足条件才能删除

	return global.DB.Where("perm_id = ?", id).Delete(&sysModel.SysPermission{}).Error
}

func (m SysPermissionService) CreatePermission(p *sysModel.SysPermission) error {
	return nil
}

func (m SysPermissionService) GetPermissionById(id int64) (sysModel.SysPermission, error) {
	var p sysModel.SysPermission
	err := global.DB.
		Where("perm_id = ?", id).
		Find(&p).Limit(1).Error

	return p, err
}

func (m SysPermissionService) UpdatePermission(perm *sysModel.SysPermission) error {
	var exist sysModel.SysPermission
	err := global.DB.Where("perm_code = ?", perm.ModuleCode).
		Find(&exist).Limit(1).Error
	if err != nil {
		return err
	}

	if exist.PermId != 0 && exist.PermId != perm.PermId {
		return errors.New("权限编码 " + perm.PermCode + " 已经存在")
	}

	return global.DB.Model(&sysModel.SysPermission{PermId: perm.PermId}).Updates(perm).Error
}
