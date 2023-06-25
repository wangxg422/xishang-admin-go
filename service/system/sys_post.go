package system

import (
	"backend/common/enmu"
	"backend/global"
	sysModel "backend/model/system"
)

type SysPostService struct {
}

func (m *SysPostService) CreatePost(Post *sysModel.SysPost) error {
	res := global.DB.Create(&Post)

	return res.Error
}

func (m *SysPostService) UpdatePost(Post *sysModel.SysPost) error {
	res := global.DB.Model(&sysModel.SysPost{PostId: Post.PostId}).Updates(&Post)
	return res.Error
}

func (m *SysPostService) DeletePost(Postid int64) error {
	res := global.DB.Model(&sysModel.SysPost{PostId: Postid}).Update("del_flag", enmu.EnmuGroupApp.DelFlagDelete.GetCode())
	return res.Error
}

func (m *SysPostService) ListPost() ([]sysModel.SysPost, error) {
	var list []sysModel.SysPost

	res := global.DB.Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode()).Find(&list)

	return list, res.Error
}

func (m *SysPostService) GetPostById(Postid int64) (sysModel.SysPost, error) {
	Post := sysModel.SysPost{
		PostId: Postid,
	}

	res := global.DB.Take(&Post, Postid).Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode())
	return Post, res.Error
}
