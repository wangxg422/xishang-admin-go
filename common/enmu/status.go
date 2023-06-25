package enmu

import "strconv"

type Status int

const (
	STATUS_NORMAL = iota
	STATUS_DISABLED
)

var statusDesc = [...]string{"正常", "", "停用"}

func (m Status) GetDesc() string {
	return statusDesc[m]
}

func (m Status) GetCode() string {
	return strconv.Itoa(int(m))
}

func (m Status) Size() int {
	return len(statusDesc)
}
