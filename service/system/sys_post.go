package system

import (
	"backend/common/enmu"
	"backend/global"
	sysModel "backend/model/system"
)

type SysPostService struct {
}

func (m *SysPostService) CreatePost(post *sysModel.SysPost) error {
	res := global.DB.Create(&post)

	return res.Error
}

func (m *SysPostService) UpdatePost(post *sysModel.SysPost) error {
	res := global.DB.Model(&sysModel.SysPost{PostId: post.PostId}).Updates(&post)
	return res.Error
}

func (m *SysPostService) DeletePost(id int64) error {
	res := global.DB.Model(&sysModel.SysPost{PostId: id}).Update("del_flag", enmu.DelFlagDeleted.Value())
	return res.Error
}

func (m *SysPostService) ListPost() ([]sysModel.SysPost, error) {
	var list []sysModel.SysPost

	res := global.DB.Where("del_flag = ?", enmu.DelFlagNormal.Value()).Find(&list)

	return list, res.Error
}

func (m *SysPostService) GetPostById(id int64) (sysModel.SysPost, error) {
	Post := sysModel.SysPost{
		PostId: id,
	}

	res := global.DB.Take(&Post, id).Where("del_flag = ?", enmu.DelFlagNormal.Value())
	return Post, res.Error
}

func (m *SysPostService) GetAllPost() ([]sysModel.SysPost, error) {
	var list []sysModel.SysPost
	res := global.DB.Find(&list)
	return list, res.Error
}
