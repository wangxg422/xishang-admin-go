package enmu

import "strconv"

type DelFlag int

const (
	DEL_FLAG_NORMAL = iota
	_
	DEL_FLAG_DELETE
)

var delFlagDesc = [...]string{"正常", "", "已删除"}

func (m DelFlag) GetDesc() string {
	return delFlagDesc[m]
}

func (m DelFlag) GetCode() string {
	return strconv.Itoa(int(m))
}

func (m DelFlag) Size() int {
	return len(delFlagDesc)
}
