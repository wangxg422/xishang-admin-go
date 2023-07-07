package enmu

type Status int8

const (
	Status_Normal   = 1
	Status_Disabled = 2
)

var statusMap map[int]string

func init() {
	statusMap = make(map[int]string)
	statusMap[Status_Normal] = "正常"
	statusMap[Status_Disabled] = "停用"
}

func (m Status) Desc() string {
	return statusMap[int(m)]
}

func (m Status) Value() int8 {
	return int8(m)
}

func (m Status) Size() int {
	return len(statusMap)
}

func (m Status) Equals(value int8) bool {
	return int8(m) == value
}
