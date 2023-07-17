package enmu

type MenuType string

const (
	MenuType_DIR  = "M"
	MenuType_MENU = "C"
	MenuType_BTN  = "F"
)

var menuTypeMap map[string]string

func init() {
	menuTypeMap = make(map[string]string)
	menuTypeMap[MenuType_DIR] = "目录"
	menuTypeMap[MenuType_MENU] = "菜单"
	menuTypeMap[MenuType_BTN] = "按钮"
}

func (m MenuType) Desc() string {
	return menuTypeMap[string(m)]
}

func (m MenuType) Value() string {
	return string(m)
}

func (m MenuType) Size() int {
	return len(menuTypeMap)
}

func (m MenuType) Equals(value string) bool {
	return string(m) == value
}
