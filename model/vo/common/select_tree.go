package common

type TreeSelectVO struct {
	Id       int64          `json:"id"`
	Label    string         `json:"label"`
	Children []TreeSelectVO `json:"children"`
}
