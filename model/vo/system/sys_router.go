package system

type RouterVO struct {
	// 路由名称
	Name string `json:"name"`
	// 路由地址
	Path string `json:"path"`
	// 是否隐藏路由，当设置 true 的时候该路由不会再侧边栏出现
	Hidden bool `json:"hidden"`
	// 重定向地址，当设置 noRedirect 的时候该路由在面包屑导航中不可被点击
	Redirect string `json:"redirect"`
	// 路由组件
	Component string `json:"component"`
	// 路由参数：如 {"id": 1, "name": "ry"}
	Query string `json:"query"`
	// 当路由下面的 children 声明的路由大于1个时，自动会变成嵌套的模式--如组件页面
	AlwaysShow bool `json:"alwaysShow"`
	// 其他元素
	MetaVo MetaVO `json:"meta"`
	// 子路由
	Children []RouterVO `json:"children"`
}

type MetaVO struct {
	// 路由在侧边栏和面包屑中展示的名字
	Title string `json:"title"`
	// 路由的图标，对应路径src/assets/icons/svg
	Icon string `json:"icon"`
	// 设置为true，则会被 <keep-alive>缓存
	Cached bool `json:"cached"`
	// 内链地址（http(s)://开头）
	Link string `json:"link"`
}
