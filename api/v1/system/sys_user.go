package system

import (
	"backend/initial/logger"
	"backend/model/common/response"
	sysDto "backend/model/dto/system"
	"backend/utils/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type SysUserApi struct {
}

func (m *SysUserApi) CreateUser(c *gin.Context) {
	userDto := sysDto.SysUserCreateDTO{}
	if err := c.ShouldBindJSON(&userDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.CreateUser(userDto); err != nil {
		logger.Error("create user failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysUserApi) GetUserPage(c *gin.Context) {
	params := &sysDto.SysUserQueryDTO{}
	err := c.ShouldBind(params)

	if err != nil {
		logger.Error("参数解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	users, err := userService.GetUserPage(params)
	if err != nil {
		logger.Error("查询失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(users, c)
}

func (m *SysUserApi) GetUserById(c *gin.Context) {
	userIdStr := c.Param("userId")

	res := make(map[string]any)
	if userIdStr == "" {
		response.FailWithMessage("请先登录", c)
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		logger.Error("userId 转换失败", zap.Error(err))
		response.FailWithMessage("userId 转换失败", c)
		return
	}

	userInfo, err := userService.GetUserInfo(userId)
	if err != nil {
		logger.Error("查询用户信息失败,userId:%s", zap.String("userId", userIdStr), zap.Error(err))
		response.FailWithMessage("查询用户信息失败", c)
		return
	}
	res["userInfo"] = userInfo

	postList := userInfo.SysPosts
	if postList != nil {
		var postIds []int64
		for _, p := range postList {
			postIds = append(postIds, p.PostId)
		}
		res["postIds"] = postIds
	}

	roleList := userInfo.SysRoles
	if roleList != nil {
		var roleIds []int64
		for _, r := range roleList {
			roleIds = append(roleIds, r.RoleId)
		}
		res["roleIds"] = roleIds
	}

	response.OkWithData(res, c)
}

func (m *SysUserApi) UpdateUser(c *gin.Context) {
	userDto := sysDto.SysUserUpdateDTO{}

	if err := c.ShouldBindJSON(&userDto); err != nil {
		logger.Error("parse param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if userDto.UserId == 0 {
		response.FailWithMessage("user id can not be null", c)
		return
	}

	userDto.UpdateBy = jwt.GetUserName(c)
	userDto.UpdateTime = time.Now()

	if err := userService.UpdateUser(&userDto); err != nil {
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
