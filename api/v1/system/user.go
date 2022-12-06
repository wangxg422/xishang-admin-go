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

var userSvc sysSvc.UserService = sysSvc.UserService{}

func (u *UserApi) CreateUser(c *gin.Context) {
	user := model.SysUser{}
	err := utils.ReadBodyToModel(&user, c)
	if err != nil {
		utils.FailWithMsg("参数解析错误", c)
		global.Log.Error("parse params failed")
		return
	}

	user.LoginDate = time.Now()
	err = userSvc.CreateUser(user)
	if err != nil {
		utils.FailWithMsg(err.Error(), c)
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
		return
	}

	utils.OkWithData(list, c)
}

func (u *UserApi) GetUserById(c *gin.Context) {
	userid := c.Param("userid")

	if userid == "" {
		utils.FailWithInfo(config.CodeParamNull, nil, "", c)
		return
	}

	id, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		utils.FailWithCodeMsg(config.CodeParamParseError, "参数解析错误", c)
		return
	}

	user, err := userSvc.GetUserById(id)
	if err != nil {
		if utils.NoRecord(err) {
			utils.OkWithEmptyObj(c)
			return
		}
		utils.FailWithMsg("查询失败", c)
		global.Log.Error("", zap.Error(err))
		return
	}

	utils.OkWithInfo(user, "查询成功", c)
}

func (u *UserApi) UpdateUser(c *gin.Context) {
	user := model.SysUser{}

	err := utils.ReadBodyToModel(&user, c)
	if err != nil {
		utils.FailWithMsg("参数解析错误", c)
		global.Log.Error("parse param error", zap.Error(err))
		return
	}

	userid := c.Param("userid")
	if userid == "" {
		utils.FailWithCode(config.CodeParamNull, c)
		global.Log.Error("parse param error", zap.Error(err))
		return
	}

	user.UserId, err = strconv.ParseInt(userid, 10, 64)
	if err != nil {
		utils.FailWithCode(config.CodeParamNull, c)
		return
	}

	err = userSvc.UpdateUser(user)
	if err != nil {
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
		utils.FailWithCode(config.CodeParamParseError, c)
		global.Log.Error("can not parse userid")
		return
	}

	err = userSvc.DeleteUser(userId)
	if err != nil {
		utils.FailWithMsg("删除失败", c)
		global.Log.Error("error when delete user", zap.Error(err))
		return
	}

	utils.Ok(c)
}
