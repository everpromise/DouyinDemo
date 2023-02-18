package services

import (
	"github.com/RaymondCode/simple-demo/entity"
	"mime/multipart"
)

/*
所有video相关接口
*/

type VideoService interface {
	//喜欢视频列表
	FavorVideoList() *[]entity.Video

	// 视频发布
	VideoPublish(token string, title string, desc string, file *multipart.FileHeader) (*entity.UserInfo, *entity.Video, error)

	// 视频列表 相同用户
	VideoListByid(userId string) []entity.VideoResponse

	// 视频列表，全部用户
	VideoList() []entity.VideoResponse
}
