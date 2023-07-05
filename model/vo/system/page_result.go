package system

type PageResult struct {
	Rows  any   `json:"rows"`
	Total int64 `json:"total"`
}
