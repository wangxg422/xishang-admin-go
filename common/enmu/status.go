package enmu

type Status int

const (
	Status_Normal   = 0
	Status_Disabled = 1
)

var statusMap map[int]string

func init() {
	statusMap = make(map[int]string)
	statusMap[Status_Normal] = "正常"
	statusMap[Status_Disabled] = "停用"
}

func (m Status) GetDesc() string {
	return statusMap[int(m)]
}

func (m Status) GetValue() int {
	return int(m)
}

func (m Status) Size() int {
	return len(statusMap)
}

func (m Status) Equals(value int) bool {
	return int(m) == value
}
