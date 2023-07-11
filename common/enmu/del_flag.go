package enmu

type DelFlag int8

const (
	DelFlag_Normal  = 1
	DelFlag_Deleted = 2
)

var delFlagMap map[int]string

func init() {
	delFlagMap = make(map[int]string)
	delFlagMap[DelFlag_Normal] = "正常"
	delFlagMap[DelFlag_Deleted] = "已删除"
}

func (m DelFlag) Desc() string {
	return delFlagMap[int(m)]
}

func (m DelFlag) Value() int8 {
	return int8(m)
}

func (m DelFlag) Size() int {
	return len(delFlagMap)
}

func (m DelFlag) Equals(value int8) bool {
	return int8(m) == value
}
func (m DelFlag) Name() string {
	return "del_flag"
}
