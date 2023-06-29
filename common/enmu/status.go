package enmu

type Status string

const (
	STATUS_NORMAL   = "0"
	STATUS_DISABLED = "1"
)

var statusDesc = [...]string{"正常", "", "停用"}

func (m Status) GetDesc() string {
	return ""
}

func (m Status) GetCode() string {
	return string(m)
}

func (m Status) Size() int {
	return 2
}

func (m Status) Compare(value string) bool {
	return string(m) == value
}
