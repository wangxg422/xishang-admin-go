package enmu

type MenuType string

const (
	DIR  = "M"
	MENU = "C"
	BTN  = "F"
)

func (m MenuType) GetDesc() string {
	switch string(m) {
	case "M":
		return "目录"
	case "C":
		return "菜单"
	case "F":
		return "按钮"
	default:
		return "undefined menu type"
	}
}

func (m MenuType) GetCode() string {
	return string(m)
}

func (m MenuType) Size() int {
	return 3
}
