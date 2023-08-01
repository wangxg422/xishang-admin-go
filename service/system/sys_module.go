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

	return global.DB.Create(mod).Error
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
	return global.DB.Where("module_id IN ?", id).Delete(&sysModel.SysModule{}).Error
}

func (m SysModuleService) UpdateModule(mod *sysModel.SysModule) error {
	var exist sysModel.SysModule
	err := global.DB.Where("module_code = ?", mod.ModuleCode).
		Find(&exist).Limit(1).Error
	if err != nil {
		return err
	}

	if exist.ModuleId != 0 && exist.ModuleId != mod.ModuleId {
		return errors.New("模块编码 " + mod.ModuleName + " 已经存在")
	}

	return global.DB.Model(&sysModel.SysModule{ModuleId: mod.ModuleId}).Updates(mod).Error
}
