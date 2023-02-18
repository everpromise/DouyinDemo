package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	Response
	VideoList []entity.VideoResponse `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// 获取数据
	token := c.PostForm("token")
	title := c.PostForm("title")
	data, err := c.FormFile("data")

	// 更新数据库
	var videoService services.VideoService
	videoService = services.VideoServiceImpl{}
	//user := entity.UserInfo{}
	videoService.VideoPublish(token, data.Filename, title, data)
	// 返回

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//filename := filepath.Base(data.Filename)

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  data.Filename + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userId := c.Query("user_id")
	srv := services.VideoServiceImpl{}
	videoList := srv.VideoListByid(userId)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
