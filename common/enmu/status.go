package enmu

type Status string

const (
	Status_Normal   = "1"
	Status_Disabled = "2"
)

var statusMap map[string]string

func init() {
	statusMap = make(map[string]string)
	statusMap[Status_Normal] = "正常"
	statusMap[Status_Disabled] = "停用"
}

func (m Status) Desc() string {
	return statusMap[string(m)]
}

func (m Status) Value() string {
	return string(m)
}

func (m Status) Size() int {
	return len(statusMap)
}

func (m Status) Equals(value string) bool {
	return string(m) == value
}
