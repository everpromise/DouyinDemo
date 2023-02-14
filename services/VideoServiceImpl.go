package services

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mapper"
)

type VideoServiceImpl struct {
}

func (u VideoServiceImpl) FavorVideoList() *[]entity.Video {
	videolist, err := mapper.MakeFavoriteList()
	if err != nil {
		return &videolist
	} else {
		return &videolist
	}
}
