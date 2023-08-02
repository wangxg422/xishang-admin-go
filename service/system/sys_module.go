package system

import (
	"backend/global"
	sysModel "backend/model/system"
	"errors"
)

type SysModuleService struct {
}

func (m SysModuleService) CreateModule(mod *sysModel.SysModule) error {
	var exist sysModel.SysModule
	err := global.DB.Where("module_code = ?", mod.ModuleCode).
		Find(&exist).Limit(1).Error
	if err != nil {
		return err
	}

	if exist.ModuleId != 0 {
		return errors.New("模块编码 " + mod.ModuleName + " 已经存在")
	}

	err = global.DB.Create(mod).Error
	if err != nil {
		return err
	}

	// 创建模块成功后即创建该模块根权限 MOD:*:*
	p := &sysModel.SysPermission{
		ModuleCode: mod.ModuleCode,
		PermName:   mod.ModuleName + " " + mod.ModuleCode + " 所有权限",
		PermCode:   mod.ModuleCode + ":*:*",
		CreateBy:   mod.CreateBy,
		Status:     mod.Status,
	}
	err = permissionService.CreatePermission(p)

	return err
}

func (m SysModuleService) GetModule() ([]sysModel.SysModule, error) {
	var list []sysModel.SysModule
	err := global.DB.Model(&sysModel.SysModule{}).Find(&list).Error

	return list, err
}

func (m SysModuleService) GetModuleById(id int64) (sysModel.SysModule, error) {
	var mod sysModel.SysModule
	err := global.DB.Model(&sysModel.SysModule{}).Where("module_id = ?", id).Find(&mod).Error

	return mod, err
}

func (m SysModuleService) DeleteModule(id int64) error {
	// TODO 满足条件才能删除

	return global.DB.Where("module_id = ?", id).Delete(&sysModel.SysModule{}).Error
}

func (m SysModuleService) UpdateModule(mod *sysModel.SysModule) error {
	var exist sysModel.SysModule
	err := global.DB.Where("module_code = ?", mod.ModuleCode).
		Find(&exist).Limit(1).Error
	if err != nil {
		return err
	}

	if exist.ModuleId != 0 && exist.ModuleId != mod.ModuleId {
		return errors.New("模块编码 " + mod.ModuleCode + " 已经存在")
	}

	return global.DB.Model(&sysModel.SysModule{ModuleId: mod.ModuleId}).Updates(mod).Error
}
