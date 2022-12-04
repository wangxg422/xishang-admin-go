package config

const (
	CodeSuccess         = 200
	CodeSysError        = 50000
	CodeParamNull       = 50001
	CodeParamParseError = 50002
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
