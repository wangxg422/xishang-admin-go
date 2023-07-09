package system

import (
	"backend/common/enmu"
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
	"errors"
)

type SysPostService struct {
}

func (m *SysPostService) CreatePost(post *sysModel.SysPost) error {
	var existPost sysModel.SysPost
	err := global.DB.Where("post_code = ? AND del_flag = ?", post.PostCode, enmu.DelFlagNormal.Value()).
		Find(&existPost).Limit(1).Error
	if err != nil {
		return err
	}

	if existPost.PostId != 0 {
		return errors.New("职位编码 " + post.PostCode + " 已经存在")
	}

	return global.DB.Create(post).Error
}

func (m *SysPostService) UpdatePost(data *sysModel.SysPost) error {
	// 查询是否存在postCode的岗位
	existPost, err := m.GetPostByCode(data.PostCode)
	if err != nil {
		return err
	}

	if existPost.PostId != 0 && existPost.PostId != data.PostId {
		return errors.New("职位编码 " + data.PostCode + " 已经存在")
	}

	vMap, err := utils.StructToMap(data)
	if err != nil {
		return err
	}

	utils.DeleteKvWhenUpdate(vMap)
	utils.DeleteKv(vMap, "sys_users", "del_flag")

	return global.DB.Model(&sysModel.SysPost{PostId: data.PostId}).Updates(vMap).Error
}

func (m *SysPostService) DeletePost(ids []int64) error {
	return global.DB.Model(&sysModel.SysPost{}).
		Where("post_id IN ?", ids).
		Update("del_flag", enmu.DelFlagDeleted.Value()).Error
}

func (m *SysPostService) GetPostById(id int64) (sysModel.SysPost, error) {
	Post := sysModel.SysPost{
		PostId: id,
	}

	res := global.DB.Find(&Post).Where("post_id = ? AND del_flag = ?", id, enmu.DelFlagNormal.Value()).Limit(1)
	return Post, res.Error
}

func (m *SysPostService) GetPostByCode(code string) (sysModel.SysPost, error) {
	Post := sysModel.SysPost{
		PostCode: code,
	}

	res := global.DB.Where("post_code = ? AND del_flag = ?", code, enmu.DelFlagNormal.Value()).Limit(1).Find(&Post)
	return Post, res.Error
}

func (m *SysPostService) GetAllPost() ([]sysModel.SysPost, error) {
	var list []sysModel.SysPost
	res := global.DB.Where("del_flag = ?", enmu.DelFlagNormal.Value()).Find(&list)
	return list, res.Error
}

func (m *SysPostService) GetPostPage(params *sysDto.SysPostQueryDTO) (sysVo.PageResult, error) {
	pageResult := sysVo.PageResult{}

	db := global.DB.Model(&sysModel.SysPost{})

	likeArr := []string{
		"post_code",
		"post_name",
	}

	utils.ConcatLikeWhereCondition(db, likeArr, params.PostCode, params.PostName)
	utils.ConcatOneEqualsStrWhereCondition(db, "status", params.Status)
	db.Where("del_flag = ?", enmu.DelFlagNormal.Value()).Order("post_sort, post_code")

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	limit, offset := params.PageInfo.Paging()
	db = db.Limit(limit).Offset(offset)

	var posts []sysModel.SysPost
	res := db.Find(&posts)

	pageResult.Total = total
	pageResult.Rows = posts

	return pageResult, res.Error
}
