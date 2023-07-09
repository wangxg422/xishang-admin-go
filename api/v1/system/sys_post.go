package system

import (
	"backend/common/constant"
	"backend/common/enmu"
	"backend/initial/logger"
	"backend/model/common/response"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	"backend/utils"
	"backend/utils/jwt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysPostApi struct{}

func (m *SysPostApi) CreatePost(c *gin.Context) {
	postDto := sysDto.SysCreatePostDTO{}
	if err := c.ShouldBindJSON(&postDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	post := &sysModel.SysPost{}
	postDto.Convert(post)
	post.CreateBy = jwt.GetUserName(c)
	post.DelFlag = enmu.DelFlagNormal.Value()

	if err := postService.CreatePost(post); err != nil {
		logger.Error("create post failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysPostApi) GetPostById(c *gin.Context) {
	id := c.Param("postId")

	if id == "" {
		response.FailWithMessage("postId is null", c)
		return
	}

	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	post, err := postService.GetPostById(postId)
	if err != nil {
		logger.Error("search post failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(post, c)
}

func (m *SysPostApi) UpdatePost(c *gin.Context) {
	postDto := sysDto.SysUpdatePostDTO{}

	if err := c.ShouldBindJSON(&postDto); err != nil {
		logger.Error("parse param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if postDto.PostId == 0 {
		response.FailWithMessage("postId is null", c)
		return
	}

	post := &sysModel.SysPost{}
	postDto.Convert(post)
	post.UpdateBy = jwt.GetUserName(c)

	if err := postService.UpdatePost(post); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysPostApi) DeletePost(c *gin.Context) {
	postIdStr := c.Param("postId")
	if postIdStr == "" {
		response.FailWithMessage("configId is null", c)
		return
	}

	ids := strings.Split(postIdStr, constant.Comma)
	postIds, err := utils.StrToInt64Array(ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = postService.DeletePost(postIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysPostApi) GetPostPage(c *gin.Context) {
	params := &sysDto.SysPostQueryDTO{}
	err := c.ShouldBind(params)

	if err != nil {
		logger.Error("参数解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	posts, err := postService.GetPostPage(params)
	if err != nil {
		logger.Error("查询岗位失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(posts, c)
}
