package system

type PageResult struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}
