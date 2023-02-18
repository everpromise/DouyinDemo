package services

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mapper"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/disintegration/imaging"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
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

func (u VideoServiceImpl) VideoPublish(token string, title string, desc string, file *multipart.FileHeader) (*entity.UserInfo, *entity.Video, error) {
	userInfo := mapper.SelectUserInfoById(token)

	// 存储视频
	filename := filepath.Base(file.Filename)
	finalName := fmt.Sprintf("%d_%s", userInfo.Id, filename)
	saveFile := filepath.Join(config.LocalVideoPrefixPath, finalName)
	utils.SaveUploadedFile(file, saveFile)

	var video = entity.Video{
		AuthorId:      userInfo.Id,
		PlayUrl:       config.PlayUrlPrefix + finalName,
		CoverUrl:      config.CoverUrl + strings.Split(title, ".")[0] + ".png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         desc,
	}
	fmt.Printf("%+v", video)
	fmt.Printf("\n%+v", config.LocalVideoPrefixPath+finalName)
	//fmt.Printf("\n%+v", config.LocalCoverPrefixPath+strings.Split(title, ".")[0]+".jpg")
	// 存储封面
	reader := utils.ReadFrameAsJpeg(config.LocalVideoPrefixPath+finalName, 1)
	img, err := imaging.Decode(reader)
	if err != nil {
		return userInfo, &video, err
	}
	err = imaging.Save(img, config.LocalCoverPrefixPath+strings.Split(title, ".")[0]+".png")
	if err != nil {
		return userInfo, &video, err
	}

	if err := mapper.InsertVideo(&video); err != nil {
		fmt.Printf("%#v", err)
		return userInfo, &video, err
	}
	return userInfo, &video, nil
}

func (u VideoServiceImpl) VideoListByid(userId string) []entity.VideoResponse {
	videoList, err := mapper.VideoListById(userId)
	n := len(videoList)
	if n == 0 {
		return nil
	}
	var videoResponseList []entity.VideoResponse

	for _, v := range videoList {
		user := mapper.SelectUserInfoById(strconv.Itoa(int(v.AuthorId)))
		videoResponseList = append(videoResponseList, entity.VideoResponse{
			Id:            v.Id,
			Author:        *user,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
		})
	}

	if err != nil {
		return videoResponseList
	}
	return videoResponseList
}

func (u VideoServiceImpl) VideoList() []entity.VideoResponse {
	var videos []entity.Video
	videos = mapper.VideoListAll()
	var videoRespons []entity.VideoResponse
	for _, v := range videos {
		videoRespons = append(videoRespons, entity.VideoResponse{
			Id:            v.Id,
			Author:        *mapper.SelectUserInfoById(strconv.Itoa(int(v.AuthorId))),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
		})
	}
	return videoRespons
}
