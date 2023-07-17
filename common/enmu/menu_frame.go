package enmu

type MenuFrame string

const (
	Menu_Frame     = "1"
	Menu_Not_Frame = "2"
)

var menuFrameMap map[string]string

func init() {
	menuFrameMap = make(map[string]string)
	menuFrameMap[Menu_Frame] = "是"
	menuFrameMap[Menu_Not_Frame] = "否"
}

func (m MenuFrame) Desc() string {
	return menuFrameMap[m.Value()]
}

func (m MenuFrame) Value() string {
	return string(m)
}

func (m MenuFrame) Size() int {
	return len(menuFrameMap)
}

func (m MenuFrame) Equals(value string) bool {
	return string(m) == value
}
