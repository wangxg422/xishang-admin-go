package system

import (
	"backend/common/constant"
	"backend/common/enmu"
	"backend/initial/logger"
	"backend/model/common/response"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
	"backend/utils/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

type SysMenuApi struct{}

func (m *SysMenuApi) CreateMenu(c *gin.Context) {
	menuDto := sysDto.SysMenuCreateDTO{}
	if err := c.ShouldBindJSON(&menuDto); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	menu := &sysModel.SysMenu{}
	err := menuDto.Convert(menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	menu.CreateBy = jwt.GetUserName(c)

	if err := menuService.CreateMenu(menu); err != nil {
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

func (m *SysMenuApi) GetMenu(c *gin.Context) {
	params := &sysDto.SysMenuQuery{}
	err := c.ShouldBind(params)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	menus, err := menuService.GetMenu(params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(menus, c)
}

func (m *SysMenuApi) UpdateMenu(c *gin.Context) {
	menuDto := sysDto.SysMenuUpdateDTO{}

	if err := c.ShouldBindJSON(&menuDto); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	menu := &sysModel.SysMenu{}
	err := menuDto.Convert(menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menu.UpdateBy = jwt.GetUserName(c)
	menu.UpdateTime = time.Now()

	if err := menuService.UpdateMenu(menu); err != nil {
		logger.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysMenuApi) DeleteMenu(c *gin.Context) {
	id := c.Param("menuId")

	if id == "" {
		response.FailWithMessage("menuId is null", c)
		return
	}

	menuId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("menuId convert failed", c)
		return
	}

	if err := menuService.DeleteMenu(menuId); err != nil {
		logger.Error("delete menu failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

// GetRouterByUserId
/*
* 路由配置项
*
* alwaysShow: true                 // 当一个路由下面的 children 声明的路由大于1个时，自动会变成嵌套的模式--如组件页面
*                                  // 只有一个时，会将那个子路由当做根路由显示在侧边栏--如引导页面
*                                  // 若你想不管路由下面的 children 声明的个数都显示你的根路由
*                                  // 你可以设置 alwaysShow: true，这样它就会忽略之前定义的规则，一直显示根路由
* redirect: noRedirect             // 当设置 noRedirect 的时候该路由在面包屑导航中不可被点击
* name:'router-name'               // 设定路由的名字，一定要填写不然使用<keep-alive>时会出现各种问题
* path: '/system'                  // 路由路径
* hidden: true                     // 当设置 true 的时候该路由不会再侧边栏出现 如401，login等页面，或者如一些编辑页面/edit/1
* component: 'system/user/index'   // 路由组件地址
* query: '{"id": 1, "name": "ry"}' // 访问路由的默认传递参数
* roles: ['admin', 'common']       // 访问路由的角色权限
* permissions: ['a:a:a', 'b:b:b']  // 访问路由的菜单权限
* meta : {
   noCache: true                   // 如果设置为true，则不会被 <keep-alive> 缓存(默认 false)
   title: 'title'                  // 设置该路由在侧边栏和面包屑中展示的名字
   icon: 'svg-name'                // 设置该路由的图标，对应路径src/assets/icons/svg
   link: ''                         // 内链地址（http(s)://开头）
   breadcrumb: false               // 如果设置为false，则不会在breadcrumb面包屑中显示
   activeMenu: '/system/user'      // 当路由设置了该属性，则会高亮相对应的侧边栏。
 }
*/
func (m *SysMenuApi) GetRouterByUserId(c *gin.Context) {
	userId := jwt.GetUserID(c)
	if userId == 0 {
		response.FailWithMessage("请先登录", c)
		return
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
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(menus) == 0 {
		response.OkWithEmptyList(c)
		return
	}

	// 由菜单构建菜单树，parentId为0的是根节点
	root := &sysModel.SysMenu{ParentId: 0}
	m.buildMenuTree(menus, root)

	response.OkWithData(m.buildMenus(root.Children), c)
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

func (m *SysMenuApi) buildMenus(menus []sysModel.SysMenu) []sysVo.RouterVO {
	var routes []sysVo.RouterVO

	if len(menus) == 0 {
		return routes
	}

	for _, menu := range menus {
		r := sysVo.RouterVO{
			Name:      menu.Name,
			Path:      getRouterPath(menu),
			Hidden:    menu.Hidden == "1",
			Component: getComponent(menu),
			Query:     menu.Query,
			MetaVo: sysVo.MetaVO{
				Title:  menu.Title,
				Icon:   menu.Icon,
				Cached: enmu.MenuIsCache.Equals(menu.Cached),
				Link:   ifHttpLink(menu.Path),
			},
		}

		cMenus := menu.Children
		if len(cMenus) > 0 && enmu.MenuTypeDir.Equals(menu.Type) {
			r.AlwaysShow = true
			r.Redirect = constant.NoRedirect
			r.Children = m.buildMenus(cMenus)
		} else if isMenuFrame(menu) {
			var childrenList []sysVo.RouterVO

			r.MetaVo = sysVo.MetaVO{}
			children := sysVo.RouterVO{
				Name:      menu.Name,
				Path:      menu.Path,
				Component: menu.Component,
				MetaVo: sysVo.MetaVO{
					Title:  menu.Title,
					Icon:   menu.Icon,
					Cached: enmu.MenuIsCache.Equals(menu.Cached),
					Link:   ifHttpLink(menu.Path),
				},
				Query: menu.Query,
			}
			childrenList = append(childrenList, children)
			r.Children = childrenList
		} else if menu.ParentId == 0 && isInnerLink(menu) {
			r.MetaVo = sysVo.MetaVO{
				Title: menu.Title,
				Icon:  menu.Icon,
			}
			r.Path = "/"

			var childrenList []sysVo.RouterVO
			children := sysVo.RouterVO{
				Name:      menu.Name,
				Path:      innerLinkReplaceEach(menu.Path),
				Component: constant.INNER_LINK,
				MetaVo: sysVo.MetaVO{
					Title: menu.Title,
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

func (m *SysMenuApi) getChildList(menus []sysModel.SysMenu, menu *sysModel.SysMenu) []sysModel.SysMenu {
	var children []sysModel.SysMenu

	if menus == nil || len(menus) == 0 {
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
	if strings.HasPrefix(link, constant.HTTP) || strings.HasPrefix(link, constant.HTTPS) {
		return link
	}

	return ""
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
	return menu.ParentId != 0 && enmu.MenuTypeDir.Equals(menu.Type)
}

func isInnerLink(menu sysModel.SysMenu) bool {
	return enmu.MenuIsNotFrame.Equals(menu.Frame) &&
		(strings.HasPrefix(menu.Path, constant.HTTP) ||
			strings.HasPrefix(menu.Path, constant.HTTPS))
}

func getRouterPath(menu sysModel.SysMenu) string {
	routerPath := menu.Path

	// 内链打开外网方式
	if menu.ParentId != 0 && isInnerLink(menu) {
		routerPath = innerLinkReplaceEach(routerPath)
	}

	// 非外链并且是一级目录（类型为目录）
	if menu.ParentId == 0 && enmu.MenuTypeDir.Equals(menu.Type) &&
		enmu.MenuIsNotFrame.Equals(menu.Frame) {
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

func isMenuFrame(menu sysModel.SysMenu) bool {
	return menu.ParentId == 0 &&
		enmu.MenuTypeMenu.Equals(menu.Type) &&
		enmu.MenuIsNotFrame.Equals(menu.Frame)
}
