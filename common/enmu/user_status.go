package enmu

import "strconv"

type UserStatus int

const (
	USER_STATUS_NORMAL = iota
	USER_STATUS_DISABLED
)

var userStatusDesc = [...]string{"正常", "", "停用"}

func (m UserStatus) GetDesc() string {
	return userStatusDesc[m]
}

func (m UserStatus) GetCode() string {
	return strconv.Itoa(int(m))
}

func (m UserStatus) Size() int {
	return len(userStatusDesc)
}
