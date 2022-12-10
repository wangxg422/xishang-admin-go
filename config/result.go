package config

const (
	OptCodeSuccess         = 0
	OptCodeSysError        = 50000
	OptCodeParamCanNotNull = 50001
	OptCodeParamParseError = 50002
)

var optCodeMap = make(map[int32]string, 10)

func init() {
	optCodeMap[OptCodeSuccess] = "成功"
	optCodeMap[OptCodeSysError] = "系统错误"
	optCodeMap[OptCodeParamCanNotNull] = "参数不能为空"
	optCodeMap[OptCodeParamParseError] = "参数错误"
}

func OptCodeDesc(code int32) string {
	return optCodeMap[code]
}

type Result struct {
	Code    int32       `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
