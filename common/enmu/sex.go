package enmu

type Sex string

const (
	Sex_M = "1"
	Sex_F = "2"
	Sex_U = "3"
)

var sexMap map[string]string

func init() {
	sexMap = make(map[string]string)
	sexMap[Sex_M] = "男"
	sexMap[Sex_F] = "女"
	sexMap[Sex_U] = "未知"
}

func (m Sex) Desc() string {
	return sexMap[string(m)]
}

func (m Sex) Value() string {
	return string(m)
}

func (m Sex) Size() int {
	return len(sexMap)
}

func (m Sex) Equals(value string) bool {
	return string(m) == value
}
