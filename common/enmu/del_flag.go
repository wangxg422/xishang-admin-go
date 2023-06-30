package enmu

type DelFlag int

const (
	DelFlag_Normal  = 0
	DelFlag_Deleted = 2
)

var delFlagMap map[int]string

func init() {
	delFlagMap = make(map[int]string)
	delFlagMap[DelFlag_Normal] = "正常"
	delFlagMap[DelFlag_Deleted] = "已删除"
}

func (m DelFlag) GetDesc() string {
	return delFlagMap[int(m)]
}

func (m DelFlag) GetValue() int {
	return int(m)
}

func (m DelFlag) Size() int {
	return len(delFlagMap)
}

func (m DelFlag) Equals(value int) bool {
	return int(m) == value
}
