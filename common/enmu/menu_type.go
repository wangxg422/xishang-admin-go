package enmu

type MenuType string

const (
	MenuType_M = "M"
	MenuType_C = "C"
	MenuType_F = "F"
)

var menuTypeMap map[string]string

func init() {
	menuTypeMap = make(map[string]string)
	menuTypeMap[MenuType_M] = "目录"
	menuTypeMap[MenuType_C] = "菜单"
	menuTypeMap[MenuType_F] = "按钮"
}

func (m MenuType) GetDesc() string {
	return menuTypeMap[string(m)]
}

func (m MenuType) GetValue() string {
	return string(m)
}

func (m MenuType) Size() int {
	return len(menuTypeMap)
}

func (m MenuType) Equals(value string) bool {
	return string(m) == value
}
