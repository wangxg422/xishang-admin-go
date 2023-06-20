package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type EmptyMap struct{}
type EmptyList []string

func buildResult(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	buildResult(SUCCESS, map[string]any{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	buildResult(SUCCESS, map[string]any{}, message, c)
}

func OkWithData(data any, c *gin.Context) {
	buildResult(SUCCESS, data, "成功", c)
}

func OkWithDetailed(data any, message string, c *gin.Context) {
	buildResult(SUCCESS, data, message, c)
}

func OkWithInfo(code int, data any, message string, c *gin.Context) {
	buildResult(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	buildResult(ERROR, map[string]any{}, "失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	buildResult(ERROR, nil, message, c)
}

func FailWithDetailed(data any, message string, c *gin.Context) {
	buildResult(ERROR, data, message, c)
}

func FailWithInfo(code int, data any, message string, c *gin.Context) {
	buildResult(ERROR, data, message, c)
}
