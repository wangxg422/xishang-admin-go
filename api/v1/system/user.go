package system

import (
	"backend/config"
	"backend/global"
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
		utils.FailWithCode(config.OptCodeParamParseError, c)
		global.Log.Error("parse params failed", zap.Error(err))
		return
	}

	user.LoginDate = time.Now()
	if err := userSvc.CreateUser(user); err != nil {
		utils.FailWithMsg(err.Error(), c)
		global.Log.Error("create user failed", zap.Error(err))
		return
	}

	utils.Ok(c)
}

func (u *UserApi) ListUser(c *gin.Context) {
	list, err := userSvc.ListUser()
	if err != nil {
		if utils.NoRecord(err) {
			utils.OkWithEmptyList(c)
			return
		}
		utils.FailWithMsg("查询失败", c)
		global.Log.Error("search user failed", zap.Error(err))
		return
	}

	utils.OkWithData(list, c)
}

func (u *UserApi) GetUserById(c *gin.Context) {
	userid := c.Param("userid")

	if userid == "" {
		utils.FailWithCode(config.OptCodeParamCanNotNull, c)
		global.Log.Error("userid is null")
		return
	}

	id, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		utils.FailWithCodeMsg(config.OptCodeParamParseError, err.Error(), c)
		global.Log.Error("search user failed", zap.Error(err))
		return
	}

	user, err := userSvc.GetUserById(id)
	if err != nil {
		if utils.NoRecord(err) {
			utils.OkWithEmptyObj(c)
			return
		}
		utils.FailWithMsg("查询失败", c)
		global.Log.Error("search user failed", zap.Error(err))
		return
	}

	utils.OkWithInfo(user, "查询成功", c)
}

func (u *UserApi) UpdateUser(c *gin.Context) {
	user := model.SysUser{}

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.FailWithCode(config.OptCodeParamParseError, c)
		global.Log.Error("parse param error", zap.Error(err))
		return
	}

	userid := c.Param("userid")
	if userid == "" {
		utils.FailWithCode(config.OptCodeParamCanNotNull, c)
		global.Log.Error("userid is null")
		return
	}

	userIdInt, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		utils.FailWithCode(config.OptCodeParamParseError, c)
		global.Log.Error("userid parse error", zap.Error(err))
		return
	}

	user.UserId = userIdInt

	if err = userSvc.UpdateUser(user); err != nil {
		utils.FailWithMsg("更新失败", c)
		global.Log.Error("update error", zap.Error(err))
		return
	}

	utils.Ok(c)
}

func (u *UserApi) DeleteUser(c *gin.Context) {
	id := c.Param("userid")

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.FailWithCode(config.OptCodeParamParseError, c)
		global.Log.Error("parse userid failed", zap.Error(err))
		return
	}

	if err := userSvc.DeleteUser(userId); err != nil {
		utils.FailWithMsg("删除失败", c)
		global.Log.Error("delete user failed", zap.Error(err))
		return
	}

	utils.Ok(c)
}
