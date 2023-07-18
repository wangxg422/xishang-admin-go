package system

import (
	"backend/common/enmu"
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	"backend/utils"
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
	arr := []string{enmu.MenuTypeMenu.Value(), enmu.MenuTypeDir.Value()}
	var menus []sysModel.SysMenu
	res := global.DB.
		Where("type IN ? and status = ?", arr, enmu.StatusNormal.Value()).
		Order("parent_id, sort").
		Find(&menus)
	return menus, res.Error
}

func (m *SysMenuService) GetMenu(params *sysDto.SysMenuQuery) ([]sysModel.SysMenu, error) {
	db := global.DB.Model(&sysModel.SysMenu{})

	likeArr := []string{
		"name",
		"title",
	}

	utils.ConcatLikeWhereCondition(db, likeArr, params.Name, params.Title)
	utils.ConcatOneEqualsStrWhereCondition(db, "status", params.Status)

	var menus []sysModel.SysMenu
	res := db.Order("parent_id, sort").Find(&menus)

	return menus, res.Error
}
