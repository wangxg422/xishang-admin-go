package utils

import (
	"backend/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Null struct{}

func buildResult(code int32, data any, msg string, c *gin.Context) {
	res := config.CreateResult(code, data, msg)
	c.JSON(http.StatusOK, res)
}

func Ok(c *gin.Context) {
	buildResult(config.OptCodeSuccess, nil, "成功", c)
}

func OkWithData(data any, c *gin.Context) {
	buildResult(config.OptCodeSuccess, data, "成功", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	buildResult(config.OptCodeSuccess, nil, msg, c)
}

func OkWithInfo(data any, msg string, c *gin.Context) {
	buildResult(config.OptCodeSuccess, data, msg, c)
}

func OkWithEmptyObj(c *gin.Context) {
	buildResult(config.OptCodeSuccess, Null{}, "成功", c)
}

func OkWithEmptyList(c *gin.Context) {
	buildResult(config.OptCodeSuccess, []Null{}, "成功", c)
}

func Fail(c *gin.Context) {
	buildResult(config.OptCodeSysError, nil, "失败", c)
}

func FailWithMsg(msg string, c *gin.Context) {
	buildResult(config.OptCodeSysError, nil, msg, c)
}

func FailWithCode(code int32, c *gin.Context) {
	desc := config.OptCodeDesc(code)

	buildResult(code, nil, desc, c)
}

func FailWithCodeMsg(code int32, msg string, c *gin.Context) {
	buildResult(code, nil, msg, c)
}

func FailWithInfo(code int32, data any, msg string, c *gin.Context) {
	buildResult(code, data, msg, c)
}
