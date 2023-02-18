package mapper

import (
	"github.com/RaymondCode/simple-demo/entity"
)

//func InsertVideo(video *entity.Video) bool {
//	Db.Create(&video)
//	return true
//}

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

func InsertVideo(video *entity.Video) error {
	Db.Create(&video)

	if err := Db.Where("play_url = ?", video.PlayUrl).Find(&video).Error; err != nil {
		Db.Create(&video)
		return err
	}
	return nil
}

func VideoListById(userId string) ([]entity.Video, error) {
	var videoList []entity.Video
	if err := Db.Where("author_id = ?", userId).Find(&videoList).Error; err != nil {
		return videoList, err
	}
	return videoList, nil
}

func VideoListAll() []entity.Video {
	var videos []entity.Video
	if err := Db.Find(&videos); err != nil {
		return videos
	}
	return videos
}
