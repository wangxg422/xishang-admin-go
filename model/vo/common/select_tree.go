package common

type TreeSelectVO struct {
	Id       string         `json:"id"`
	Label    string         `json:"label"`
	Children []TreeSelectVO `json:"children"`
}
