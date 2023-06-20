package system

import (
	"backend/common/response"
	"backend/initial/logger"
	"backend/model/dto"
	sysModel "backend/model/system"
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
	userDto := dto.SysCreateUserDTO{}
	if err := c.ShouldBindJSON(&userDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	user := &sysModel.SysUser{}
	userDto.Convert(user)

	user.LoginDate = time.Now()
	if err := userSvc.CreateUser(user); err != nil {
		logger.Error("create user failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
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
		response.FailWithMessage("user id is null", c)
		return
	}

	id, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		response.FailWithMessage("user id convert failed", c)
		return
	}

	user, err := userSvc.GetUserById(id)
	if err != nil {
		if utils.NoRecord(err) {
			response.OkWithData([]string{}, c)
			return
		}
		logger.Error("search user failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(user, c)
}

func (u *UserApi) UpdateUser(c *gin.Context) {
	userDto := dto.SysUpdateUserDTO{}

	if err := c.ShouldBindJSON(&userDto); err != nil {
		logger.Error("parse param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if userDto.UserId == 0 {
		response.FailWithMessage("user id can not be null", c)
		return
	}

	user := &sysModel.SysUser{}
	userDto.Convert(user)
	if err := userSvc.UpdateUser(user); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (u *UserApi) DeleteUser(c *gin.Context) {
	id := c.Param("userid")

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("user id convert failed", c)
		return
	}

	if err := userSvc.DeleteUser(userId); err != nil {
		logger.Error("delete user failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
