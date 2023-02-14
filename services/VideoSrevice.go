package services

import "github.com/RaymondCode/simple-demo/entity"

/*
所有video相关接口
*/

type VideoService interface {
	//喜欢视频列表
	FavorVideoList() *[]entity.Video
}
