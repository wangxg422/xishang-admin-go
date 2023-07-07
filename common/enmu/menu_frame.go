package enmu

type MenuFrame int8

const (
	Menu_Frame     = 1
	Menu_Not_Frame = 2
)

var menuFrameMap map[int]string

func init() {
	menuFrameMap = make(map[int]string)
	menuFrameMap[Menu_Frame] = "是"
	menuFrameMap[Menu_Not_Frame] = "否"
}

func (m MenuFrame) Desc() string {
	return menuFrameMap[int(m)]
}

func (m MenuFrame) Value() int8 {
	return int8(m)
}

func (m MenuFrame) Size() int {
	return len(menuFrameMap)
}

func (m MenuFrame) Equals(value int8) bool {
	return int8(m) == value
}
