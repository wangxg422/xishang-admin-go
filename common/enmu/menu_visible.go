package enmu

type MenuVisible int8

const (
	Menu_Visible   = 1
	Menu_Unvisible = 2
)

var eenuVisibleMap map[int]string

func init() {
	eenuVisibleMap = make(map[int]string)
	eenuVisibleMap[Menu_Visible] = "显示"
	eenuVisibleMap[Menu_Unvisible] = "隐藏"
}

func (m MenuVisible) Desc() string {
	return eenuVisibleMap[int(m)]
}

func (m MenuVisible) Value() int8 {
	return int8(m)
}

func (m MenuVisible) Size() int {
	return len(eenuVisibleMap)
}

func (m MenuVisible) Equals(value int8) bool {
	return int8(m) == value
}
