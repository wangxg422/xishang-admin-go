package public

import (
	"backend/common/response"

	"github.com/gin-gonic/gin"
)

type LoginApi struct{}

func (m *LoginApi) Login(c *gin.Context) {
	response.Ok(c)
}

func (m *LoginApi) Register(c *gin.Context) {
	response.Ok(c)
}
