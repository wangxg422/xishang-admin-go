package system

import (
	"backend/common/response"
	"backend/initial/logger"
	model "backend/model/system"
	sysSvc "backend/service/system"
	"backend/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
}

var userSvc = sysSvc.UserService{}

func (u *UserApi) CreateUser(c *gin.Context) {
	user := model.SysUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		return
	}

	user.LoginDate = time.Now()
	if err := userSvc.CreateUser(user); err != nil {
		logger.Error("create user failed", zap.Error(err))
		return
	}

	response.Ok(c)
}

func (u *UserApi) ListUser(c *gin.Context) {
	response.Ok(c)
}

func (u *UserApi) GetUserById(c *gin.Context) {
	userid := c.Param("userid")

	if userid == "" {
		logger.Error("userid is null")
		return
	}

	id, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		logger.Error("search user failed", zap.Error(err))
		return
	}

	user, err := userSvc.GetUserById(id)
	if err != nil {
		if utils.NoRecord(err) {
			return
		}
		logger.Error("search user failed", zap.Error(err))
		return
	}

	response.OkWithData(user, c)
}

func (u *UserApi) UpdateUser(c *gin.Context) {
	user := model.SysUser{}

	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Error("parse param error", zap.Error(err))
		return
	}

	userid := c.Param("userid")
	if userid == "" {
		logger.Error("userid is null")
		return
	}

	userIdInt, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		logger.Error("userid parse error", zap.Error(err))
		return
	}

	user.UserId = userIdInt

	if err = userSvc.UpdateUser(user); err != nil {
		logger.Error("update error", zap.Error(err))
		return
	}

	response.Ok(c)
}

func (u *UserApi) DeleteUser(c *gin.Context) {
	id := c.Param("userid")

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logger.Error("parse userid failed", zap.Error(err))
		return
	}

	if err := userSvc.DeleteUser(userId); err != nil {
		logger.Error("delete user failed", zap.Error(err))
		return
	}

	response.Ok(c)
}
