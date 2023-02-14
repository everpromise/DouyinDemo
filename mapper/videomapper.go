package mapper

import (
	"github.com/RaymondCode/simple-demo/entity"
)

func InsertVideo(video *entity.Video) bool {
	Db.Create(&video)
	return true
}

func SelectVideoById(id string) (entity.Video, error) {
	video := entity.Video{}
	if err := Db.First(&video, "id = ?", id).Error; err != nil {
		return video, err
	}
	return video, nil
}

func MakeFavoriteList() ([]entity.Video, error) {
	var videolists []entity.Video
	if err := Db.Where("is_favorite = ?", "1").Find(&videolists).Error; err != nil {
		return videolists, err
	}
	return videolists, nil
}
