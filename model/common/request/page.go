package request

import (
	"backend/common/constant"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PageInfo struct {
	PageNum  int `form:"pageNum" json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

func NewPageInfo(c *gin.Context) (PageInfo, error) {
	pageNum, err := strconv.Atoi(c.Query(constant.PageNum))
	pageSize, err := strconv.Atoi(c.Query(constant.PageSize))

	if err != nil {
		return PageInfo{}, err
	}

	return PageInfo{
		PageNum:  pageNum,
		PageSize: pageSize,
	}, nil
}

func (m *PageInfo) Paging() (int, int) {
	if m.PageNum == 0 || m.PageSize == 0 {
		return 0, 0
	}

	limit := m.PageSize
	offset := m.PageSize * (m.PageNum - 1)

	return limit, offset
}
