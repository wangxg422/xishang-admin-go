package request

import (
	"backend/common/constant"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PageInfo struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Limit    int `json:"limit"`
	Offset   int `json:"offset"`
}

func NewPageInfo(c *gin.Context) (PageInfo, error) {
	pageNum, err := strconv.Atoi(c.Query(constant.PageNum))
	pageSize, err := strconv.Atoi(c.Query(constant.PageSize))

	if err != nil {
		return PageInfo{}, err
	}
	if pageNum == 0 || pageSize == 0 {
		return PageInfo{}, nil
	}

	limit := pageSize
	offset := pageSize * (pageNum - 1)

	return PageInfo{
		PageNum:  pageNum,
		PageSize: pageSize,
		Limit:    limit,
		Offset:   offset,
	}, nil
}

func (m *PageInfo) Paging() {
	if m.PageNum == 0 || m.PageSize == 0 {
		m.Limit = 0
		m.Offset = 0
	}

	m.Limit = m.PageSize
	m.Offset = m.PageSize * (m.PageNum - 1)
}
