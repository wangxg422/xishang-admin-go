package system

import (
	"backend/common/constant"
	"backend/common/enmu"
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	"backend/utils"
	"errors"
)

type SysMenuService struct {
}

func (m *SysMenuService) CreateMenu(menu *sysModel.SysMenu) error {
	// 检查menuCode是否已经存在
	existMenu, err := m.GetMenuByCode(menu.MenuCode, menu.ParentId)
	if err != nil {
		return err
	}
	if existMenu.MenuId != 0 {
		return errors.New("menu " + menu.MenuCode + " exist")
	}

	return global.DB.Create(&menu).Error
}

func (m *SysMenuService) GetMenuByCode(code string, parentId int64) (sysModel.SysMenu, error) {
	var menu sysModel.SysMenu
	err := global.DB.Where("parent_id = ? AND menu_code = ?", parentId, code).Find(&menu).Limit(1).Error
	return menu, err
}

func (m *SysMenuService) UpdateMenu(menu *sysModel.SysMenu) error {
	if menu.MenuId == 0 {
		return errors.New("menuId is null")
	}

	// 检查更新的menu_code是否已经存在
	existMenu, err := m.GetMenuByCode(menu.MenuCode, menu.ParentId)
	if err != nil {
		return err
	}
	if existMenu.MenuId != 0 && existMenu.MenuId != menu.MenuId {
		return errors.New("menu " + menu.MenuCode + " exist")
	}

	return global.DB.Omit(constant.UpdateOmit...).Save(menu).Error
}

func (m *SysMenuService) DeleteMenu(id int64) error {
	// 含有子菜单禁止删除
	count, err := m.GetChildCountById(id)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该菜单含有子菜单，禁止删除")
	}

	// TODO 有角色绑定禁止删除

	return global.DB.Delete(&sysModel.SysMenu{MenuId: id}).Error
}

func (m *SysMenuService) GetChildCountById(id int64) (int64, error) {
	var count int64 = 0
	err := global.DB.Model(&sysModel.SysMenu{}).Where("parent_id = ?", id).Count(&count).Error
	return count, err
}

func (m *SysMenuService) GetMenuById(id int64) (sysModel.SysMenu, error) {
	user := sysModel.SysMenu{
		MenuId: id,
	}

	res := global.DB.Take(&user, id).Where("del_flag = ?", enmu.StatusNormal.Value())
	return user, res.Error
}

func (m *SysMenuService) GetMenuByUserId(userId int64) ([]sysModel.SysMenu, error) {
	// TODO 根据用户角色获取有权限的菜单
	if userId == 1 {
		return m.GetAllMenu()
	}

	var menus []sysModel.SysMenu
	err := global.DB.Find(&menus).Error
	return menus, err
}

func (m *SysMenuService) GetAllMenu() ([]sysModel.SysMenu, error) {
	var menus []sysModel.SysMenu
	res := global.DB.
		Order("parent_id, sort").
		Find(&menus)
	return menus, res.Error
}

func (m *SysMenuService) GetAllDirMenuNormal() ([]sysModel.SysMenu, error) {
	arr := []string{enmu.MenuTypeMenu.Value(), enmu.MenuTypeDir.Value()}
	var menus []sysModel.SysMenu
	res := global.DB.
		Where("type IN ? and status = ?", arr, enmu.StatusNormal.Value()).
		Order("parent_id, sort").
		Find(&menus)
	return menus, res.Error
}

func (m *SysMenuService) GetDirMenuNormalByUserId(userId int64) ([]sysModel.SysMenu, error) {
	if userId == 1 {
		return m.GetAllDirMenuNormal()
	}

	var menus []sysModel.SysMenu
	err := global.DB.
		Where("status = ?", enmu.StatusNormal.Value()).
		Order("parent_id, sort").
		Find(&menus).Error

	return menus, err
}

func (m *SysMenuService) GetMenu(params *sysDto.SysMenuQuery) ([]sysModel.SysMenu, error) {
	db := global.DB.Model(&sysModel.SysMenu{})

	likeArr := []string{
		"menu_name",
		"menu_code",
	}

	utils.ConcatLikeWhereCondition(db, likeArr, params.MenuName, params.MenuCode)
	utils.ConcatOneEqualsStrWhereCondition(db, "status", params.Status)

	var menus []sysModel.SysMenu
	res := db.Order("parent_id, sort").Find(&menus)

	return menus, res.Error
}
