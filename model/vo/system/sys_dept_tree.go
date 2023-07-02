package system

type SysDeptTreeVO struct {
	Id       int64           `json:"id"`
	Label    string          `json:"label"`
	Children []SysDeptTreeVO `json:"children"`
}
