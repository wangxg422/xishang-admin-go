package utils

import (
	"backend/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func buildResult(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, config.Result{
		Code:    config.CodeSuccess,
		Data:    data,
		Message: msg,
	})
}

func Ok(c *gin.Context) {
	buildResult(config.CodeSuccess, nil, "成功", c)
}

func OkWithData(data any, c *gin.Context) {
	buildResult(config.CodeSuccess, data, "成功", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	buildResult(config.CodeSuccess, nil, msg, c)
}

func OkWithInfo(data any, msg string, c *gin.Context) {
	buildResult(config.CodeSuccess, data, msg, c)
}

func OkWithEmptyObj(c *gin.Context) {
	buildResult(config.CodeSuccess, Null{}, "成功", c)
}

func OkWithEmptyList(c *gin.Context) {
	buildResult(config.CodeSuccess, []Null{}, "成功", c)
}

func Fail(c *gin.Context) {
	buildResult(config.CodeSysError, nil, "失败", c)
}

func FailWithMsg(msg string, c *gin.Context) {
	buildResult(config.CodeSysError, nil, msg, c)
}

func FailWithCode(code int, c *gin.Context) {
	buildResult(code, nil, "失败", c)
}

func FailWithCodeMsg(code int, msg string, c *gin.Context) {
	buildResult(code, nil, msg, c)
}

func FailWithInfo(code int, data any, msg string, c *gin.Context) {
	buildResult(code, data, msg, c)
}
