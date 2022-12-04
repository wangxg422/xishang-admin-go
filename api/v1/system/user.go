package system

import (
	"backend/config"
	"backend/global"
	model "backend/model/system"
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type UserApi struct {
}

var userSvc = service.AppServiceGroupIns.SysServiceGroup.UserService

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
