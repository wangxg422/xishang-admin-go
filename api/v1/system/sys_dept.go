package system

import (
	"backend/common/enmu"
	"backend/initial/logger"
	"backend/model/common/response"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	"backend/utils"
	"backend/utils/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type SysDeptApi struct{}

func (m *SysDeptApi) CreateDept(c *gin.Context) {
	deptDto := sysDto.SysDeptCreateDTO{}
	if err := c.ShouldBindJSON(&deptDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	dept, err := deptDto.Convert()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	dept.DelFlag = enmu.DelFlagNormal.Value()
	dept.CreateBy = jwt.GetUserName(c)

	if err := deptService.CreateDept(&dept); err != nil {
		logger.Error("create dept failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysDeptApi) GetDeptById(c *gin.Context) {
	id := c.Param("deptId")

	if id == "" {
		response.FailWithMessage("dept id is null", c)
		return
	}

	deptId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("dept id convert failed", c)
		return
	}

	user, err := deptService.GetDeptById(deptId)
	if err != nil {
		if utils.NoRecord(err) {
			response.OkWithData([]string{}, c)
			return
		}
		logger.Error("search dept failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(user, c)
}

func (m *SysDeptApi) GetDept(c *gin.Context) {
	params := &sysDto.SysDeptQueryDTO{}
	err := c.ShouldBind(params)

	if err != nil {
		logger.Error("参数解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	depts, err := deptService.GetDept(params)
	if err != nil {
		logger.Error("", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(depts, c)
}

func (m *SysDeptApi) UpdateDept(c *gin.Context) {
	deptDto := sysDto.SysDeptUpdateDTO{}

	if err := c.ShouldBindJSON(&deptDto); err != nil {
		logger.Error("parse param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if deptDto.DeptId == "" {
		response.FailWithMessage("dept id can not be null", c)
		return
	}

	dept, err := deptDto.Convert()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dept.UpdateBy = jwt.GetUserName(c)

	if err := deptService.UpdateDept(&dept); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysDeptApi) DeleteDept(c *gin.Context) {
	id := c.Param("deptId")

	if id == "" {
		response.FailWithMessage("deptId is null", c)
		return
	}

	deptId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("dept id convert failed", c)
		return
	}

	if err := deptService.DeleteDept(deptId); err != nil {
		logger.Error("delete dept failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysDeptApi) GetDeptTree(c *gin.Context) {
	userId := jwt.GetUserID(c)
	if userId == 0 {
		response.FailWithMessage("请先登录", c)
	}

	var depts []sysModel.SysDept
	if userId == 1 {
		temp, err := deptService.GetAllDept()
		if err != nil {
			response.FailWithMessage("查询部门失败", c)
			return
		}
		depts = temp
	} else {
		//TODO 根据用户权限查询部门
		//depts, err := deptService.GetDeptByIds(deptIds)
		//if err != nil {
		//	response.FailWithMessage("查询部门失败", c)
		//}
	}

	vo, err := utils.ListToSelectTree(depts, "parentId", "children", "deptId", "deptName")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(vo.Children, c)
}

/**

func (m *SysDeptApi) buildDeptTree(depts []sysModel.SysDept, dept *sysModel.SysDept, tree *common.TreeSelectVO) {
	children, voChildren := m.getChildren(depts, dept)

	tree.Id = strconv.FormatInt(dept.DeptId, 10)
	tree.Label = dept.DeptName
	tree.Children = voChildren

	length := len(children)
	if length > 0 {
		for i := 0; i < length; i++ {
			m.buildDeptTree(depts, &children[i], &voChildren[i])
		}
	}
}

func (m *SysDeptApi) getChildren(depts []sysModel.SysDept, dept *sysModel.SysDept) ([]sysModel.SysDept, []common.TreeSelectVO) {
	var list []sysModel.SysDept
	var voList []common.TreeSelectVO

	for _, d := range depts {
		if d.ParentId == dept.DeptId {
			list = append(list, d)
			voList = append(voList, common.TreeSelectVO{
				Id:    strconv.FormatInt(d.DeptId, 10),
				Label: d.DeptName,
			})
		}
	}

	return list, voList
}

*/
