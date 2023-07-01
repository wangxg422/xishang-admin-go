package system

import (
	"backend/common/enmu"
	"backend/global"
	sysModel "backend/model/system"
)

type SysMenuService struct {
}

func (m *SysMenuService) CreateMenu(d *sysModel.SysMenu) error {
	res := global.DB.Create(&d)

	return res.Error
}

func (m *SysMenuService) UpdateMenu(d *sysModel.SysMenu) error {
	res := global.DB.Model(&sysModel.SysMenu{MenuId: d.MenuId}).Updates(&d)
	return res.Error
}

func (m *SysMenuService) DeleteMenu(id int64) error {
	res := global.DB.Model(&sysModel.SysMenu{MenuId: id}).Update("del_flag", enmu.DelFlagDeleted.Value())
	return res.Error
}

func (m *SysMenuService) ListMenu() ([]sysModel.SysMenu, error) {
	var list []sysModel.SysMenu

	res := global.DB.Where("del_flag = ?", enmu.DelFlagNormal.Value()).Find(&list)

	return list, res.Error
}

func (m *SysMenuService) GetMenuById(id int64) (sysModel.SysMenu, error) {
	user := sysModel.SysMenu{
		MenuId: id,
	}

	res := global.DB.Take(&user, id).Where("del_flag = ?", enmu.StatusNormal.Value())
	return user, res.Error
}

func (m *SysMenuService) GetMenuByUserId(userId int64) ([]sysModel.SysMenu, error) {
	user := sysModel.SysMenu{
		MenuId: userId,
		Status: enmu.StatusNormal.Value(),
	}

	var menus []sysModel.SysMenu
	res := global.DB.Find(&menus).Where(&user)
	return menus, res.Error
}

func (m *SysMenuService) GetAllMenu() ([]sysModel.SysMenu, error) {
	user := sysModel.SysMenu{
		Status: enmu.StatusNormal.Value(),
	}

	var menus []sysModel.SysMenu
	res := global.DB.Find(&menus).Where(&user)
	return menus, res.Error
}
