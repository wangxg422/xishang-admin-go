package response

import (
	"backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type Null struct{}
type EmptyList []string

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func buildResult(code int32, data any, msg string, c *gin.Context) {
	res := config.CreateResult(code, data, msg)
	c.JSON(http.StatusOK, res)
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]any{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, message, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data any, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]any{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]any{}, message, c)
}

func FailWithDetailed(data any, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWithInfo(code int, data any, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
