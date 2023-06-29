package system

import (
	"backend/common/enmu"
	"backend/common/response"
	"backend/initial/logger"
	"backend/model/dto"
	sysModel "backend/model/system"
	"backend/utils"
	"backend/utils/jwt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysUserApi struct {
}

func (m *SysUserApi) CreateUser(c *gin.Context) {
	userDto := dto.SysCreateUserDTO{}
	if err := c.ShouldBindJSON(&userDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	user := &sysModel.SysUser{}
	userDto.Convert(user)

	user.LoginDate = time.Now()
	user.DelFlag = enmu.EnmuGroupApp.DelFlagDelete.GetCode()
	user.Status = enmu.EnmuGroupApp.StatusNormal.GetCode()

	if err := userService.CreateUser(user); err != nil {
		logger.Error("create user failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysUserApi) ListUser(c *gin.Context) {
	response.Ok(c)
}

func (m *SysUserApi) GetUserById(c *gin.Context) {
	userid := c.Param("userId")

	if userid == "" {
		response.FailWithMessage("user id is null", c)
		return
	}

	id, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		response.FailWithMessage("user id convert failed", c)
		return
	}

	user, err := userService.GetUserById(id)
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

func (m *SysUserApi) UpdateUser(c *gin.Context) {
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
	if err := userService.UpdateUser(user); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysUserApi) DeleteUser(c *gin.Context) {
	id := c.Param("userId")

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("user id convert failed", c)
		return
	}

	if err := userService.DeleteUser(userId); err != nil {
		logger.Error("delete user failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysUserApi) GetProfile(c *gin.Context) {
	userId := jwt.GetUserID(c)

	

	id := c.Param("userId")

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("user id convert failed", c)
		return
	}

	if err := userService.DeleteUser(userId); err != nil {
		logger.Error("delete user failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
