package enmu

type Sex int8

const (
	Sex_M = 0
	Sex_F = 1
	Sex_U = 2
)

var sexMap map[int]string

func init() {
	sexMap = make(map[int]string)
	sexMap[Sex_M] = "男"
	sexMap[Sex_F] = "女"
	sexMap[Sex_U] = "未知"
}

func (m Sex) Desc() string {
	return sexMap[int(m)]
}

func (m Sex) Value() int8 {
	return int8(m)
}

func (m Sex) Size() int {
	return len(sexMap)
}

func (m Sex) Equals(value int8) bool {
	return int8(m) == value
}
