package system

import (
	"backend/common/constant"
	"backend/common/enmu"
	"backend/common/response"
	"backend/initial/logger"
	"backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
	"backend/utils/jwt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysMenuApi struct{}

func (m *SysMenuApi) CreateMenu(c *gin.Context) {
	menuDto := system.SysCreateMenuDTO{}
	if err := c.ShouldBindJSON(&menuDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	menu := &sysModel.SysMenu{}
	menuDto.Convert(menu)

	menu.Status = enmu.StatusNormal.Value()

	if err := menuService.CreateMenu(menu); err != nil {
		logger.Error("create menu failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysMenuApi) GetMenuById(c *gin.Context) {
	id := c.Param("menuId")

	if id == "" {
		response.FailWithMessage("menu id is null", c)
		return
	}

	menuId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logger.Error("menu id convert failed", zap.Error(err))
		response.FailWithMessage("menu id convert failed", c)
		return
	}

	user, err := menuService.GetMenuById(menuId)
	if err != nil {
		if utils.NoRecord(err) {
			response.OkWithData([]string{}, c)
			return
		}
		logger.Error("search menu failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(user, c)
}

func (m *SysMenuApi) ListMenu(c *gin.Context) {

}

func (m *SysMenuApi) UpdateMenu(c *gin.Context) {
	menuDto := system.SysUpdateMenuDTO{}

	if err := c.ShouldBindJSON(&menuDto); err != nil {
		logger.Error("parse param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if menuDto.MenuId == 0 {
		response.FailWithMessage("menu id can not be null", c)
		return
	}

	menu := &sysModel.SysMenu{}
	menuDto.Convert(menu)
	if err := menuService.UpdateMenu(menu); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysMenuApi) DeleteMenu(c *gin.Context) {
	id := c.Param("menuId")

	menuId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("menu id convert failed", c)
		return
	}

	if err := menuService.DeleteMenu(menuId); err != nil {
		logger.Error("delete menu failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysMenuApi) GetMenuByUserId(c *gin.Context) {
	userId := jwt.GetUserID(c)
	if userId == 0 {
		response.FailWithMessage("请先登录", c)
	}

	var menus []sysModel.SysMenu
	var err error
	// 如果是管理员，返回所有菜单
	if userId == 1 {
		menus, err = menuService.GetAllMenu()
	} else {
		menus, err = menuService.GetMenuByUserId(userId)
	}

	if err != nil {
		logger.Error("get menus failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(menus) == 0 {
		response.OkWithEmptyList(c)
	}

	// 由菜单构建菜单树，parentId为0的是根节点
	root := &sysModel.SysMenu{ParentId: 0}
	m.buildMenuTree(menus, root)

	response.OkWithData(m.buildMenus(root.Children), c)
}

func (m *SysMenuApi) buildMenus(menus []sysModel.SysMenu) []sysVo.RouterVO {
	var routes []sysVo.RouterVO

	if len(menus) == 0 {
		return routes
	}

	for _, menu := range menus {
		r := sysVo.RouterVO{
			Name:      getRouterName(menu),
			Path:      getRouterPath(menu),
			Hidden:    enmu.MenuIsNotVisible.Equals(menu.Visible),
			Component: getComponent(menu),
			Query:     menu.Query,
			MetaVo: sysVo.MetaVO{
				Title:   menu.MenuName,
				Icon:    menu.Icon,
				NoCache: enmu.MenuIsNotCache.Equals(menu.IsCache),
				Link:    ifHttpLink(menu.Path),
			},
		}

		cMenus := menu.Children
		if len(cMenus) > 0 && enmu.MenuTypeM.Equals(menu.MenuType) {
			r.AlwaysShow = true
			r.Redirect = "noRedirect"
			r.Children = m.buildMenus(cMenus)
		} else if isMenuFrame(menu) {
			var childrenList []sysVo.RouterVO

			r.MetaVo = sysVo.MetaVO{}
			children := sysVo.RouterVO{
				Path:      menu.Path,
				Component: menu.Component,
				Name:      strings.ToTitle(menu.Path),
				MetaVo: sysVo.MetaVO{
					Title:   menu.MenuName,
					Icon:    menu.Icon,
					NoCache: enmu.MenuIsCache.Equals(menu.IsCache),
					Link:    ifHttpLink(menu.Path),
				},
				Query: menu.Query,
			}
			childrenList = append(childrenList, children)
			r.Children = childrenList
		} else if menu.ParentId == 0 && isInnerLink(menu) {
			r.MetaVo = sysVo.MetaVO{
				Title: menu.MenuName,
				Icon:  menu.Icon,
			}
			r.Path = "/"

			var childrenList []sysVo.RouterVO
			routerPath := innerLinkReplaceEach(menu.Path)
			children := sysVo.RouterVO{
				Path:      routerPath,
				Component: constant.INNER_LINK,
				Name:      strings.ToTitle(routerPath),
				MetaVo: sysVo.MetaVO{
					Title: menu.MenuName,
					Icon:  menu.Icon,
					Link:  menu.Path,
				},
			}
			childrenList = append(childrenList, children)
			r.Children = childrenList
		}
		routes = append(routes, r)
	}

	return routes
}

func (m *SysMenuApi) buildMenuTree(menus []sysModel.SysMenu, menu *sysModel.SysMenu) {
	// 子节点列表
	children := m.getChildList(menus, menu)

	menu.Children = children
	length := len(menu.Children)
	for i := 0; i < length; i++ {
		m.buildMenuTree(menus, &menu.Children[i])
	}
}

func (m *SysMenuApi) getChildList(menus []sysModel.SysMenu, menu *sysModel.SysMenu) []sysModel.SysMenu {
	var children []sysModel.SysMenu

	if len(menus) == 0 {
		return children
	}

	for _, subMenu := range menus {
		if subMenu.ParentId == menu.MenuId {
			children = append(children, subMenu)
		}
	}

	return children
}

func ifHttpLink(link string) string {
	var l string
	if strings.HasPrefix(link, constant.HTTP) || strings.HasPrefix(link, constant.HTTPS) {
		l = link
	}

	return l
}

func getComponent(menu sysModel.SysMenu) string {
	component := constant.LAYOUT
	c := menu.Component
	if c != "" && !isMenuFrame(menu) {
		component = menu.Component
	} else if c == "" && menu.ParentId != 0 && isInnerLink(menu) {
		component = constant.INNER_LINK
	} else if c == "" && isParentView(menu) {
		component = constant.PARENT_VIEW
	}

	return component
}

func isParentView(menu sysModel.SysMenu) bool {
	return menu.ParentId != 0 && enmu.MenuTypeM.Equals(menu.MenuType)
}

func getRouterPath(menu sysModel.SysMenu) string {
	routerPath := menu.Path

	// 内链打开外网方式
	if menu.ParentId != 0 && isInnerLink(menu) {
		routerPath = innerLinkReplaceEach(routerPath)
	}

	// 非外链并且是一级目录（类型为目录）
	if menu.ParentId == 0 && enmu.MenuTypeM.Equals(menu.MenuType) &&
		enmu.MenuIsNotFrame.Equals(menu.IsFrame) {
		routerPath = "/" + menu.Path
	} else if isMenuFrame(menu) {
		// 非外链并且是一级目录（类型为菜单）
		routerPath = "/"
	}
	return routerPath
}

func innerLinkReplaceEach(path string) string {
	// 将http、https、www替换为空，.替换为/

	r := strings.NewReplacer("http", "", "https", "", "www", "", ".", "/")
	return r.Replace(path)
}

func isInnerLink(menu sysModel.SysMenu) bool {
	return enmu.MenuIsNotFrame.Equals(menu.IsFrame) &&
		(strings.HasPrefix(menu.Path, constant.HTTP) ||
			strings.HasPrefix(menu.Path, constant.HTTPS))
}

func getRouterName(menu sysModel.SysMenu) string {
	routerName := utils.FirstUpper(menu.Path)
	// 非外链并且是一级目录（类型为目录）
	if isMenuFrame(menu) {
		routerName = ""
	}
	return routerName
}

func isMenuFrame(menu sysModel.SysMenu) bool {
	return menu.ParentId == 0 &&
		enmu.MenuTypeC.Equals(menu.MenuType) &&
		enmu.MenuIsNotFrame.Equals(menu.IsFrame)
}
