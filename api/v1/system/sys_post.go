package system

import (
	"backend/common/enmu"
	"backend/common/response"
	"backend/initial/logger"
	"backend/model/dto"
	sysModel "backend/model/system"
	"backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysPostApi struct{}

func (m *SysPostApi) CreatePost(c *gin.Context) {
	postDto := dto.SysCreatePostDTO{}
	if err := c.ShouldBindJSON(&postDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	post := &sysModel.SysPost{}
	postDto.Convert(post)

	post.Status = enmu.StatusNormal.Value()

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
		response.FailWithMessage("post id is null", c)
		return
	}

	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("post id convert failed", c)
		return
	}

	user, err := postService.GetPostById(postId)
	if err != nil {
		if utils.NoRecord(err) {
			response.OkWithData([]string{}, c)
			return
		}
		logger.Error("search post failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(user, c)
}

func (m *SysPostApi) ListPost(c *gin.Context) {

}

func (m *SysPostApi) UpdatePost(c *gin.Context) {
	postDto := dto.SysUpdatePostDTO{}

	if err := c.ShouldBindJSON(&postDto); err != nil {
		logger.Error("parse param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if postDto.PostId == 0 {
		response.FailWithMessage("post id can not be null", c)
		return
	}

	post := &sysModel.SysPost{}
	postDto.Convert(post)
	if err := postService.UpdatePost(post); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysPostApi) DeletePost(c *gin.Context) {
	id := c.Param("postId")

	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("post id convert failed", c)
		return
	}

	if err := postService.DeletePost(postId); err != nil {
		logger.Error("delete post failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
