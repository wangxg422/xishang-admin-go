package enmu

type MenuVisible string

const (
	Menu_Visible   = "1"
	Menu_Unvisible = "2"
)

var menuVisibleMap map[string]string

func init() {
	menuVisibleMap = make(map[string]string)
	menuVisibleMap[Menu_Visible] = "显示"
	menuVisibleMap[Menu_Unvisible] = "隐藏"
}

func (m MenuVisible) Desc() string {
	return menuVisibleMap[m.Value()]
}

func (m MenuVisible) Value() string {
	return string(m)
}

func (m MenuVisible) Size() int {
	return len(menuVisibleMap)
}

func (m MenuVisible) Equals(value string) bool {
	return string(m) == value
}
