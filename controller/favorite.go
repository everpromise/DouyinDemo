package controller

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/services"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	userId := c.Query("token")
	video_id := c.Query("video_id")
	//action_type := c.Query("action_type")
	var usi services.UserService
	usi = services.UserServiceImpl{}
	userInfo := usi.UserInfo(userId)
	if userInfo != nil {
		err := usi.VideoFavorAct(video_id)
		if err == 1 {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		} else if err == 0 {
			c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "Video doesn not exist"})
		}

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	//user_id := c.Query("user_id")
	token := c.Query("token")
	var usi services.UserService
	usi = services.UserServiceImpl{}
	var vsi services.VideoService
	vsi = services.VideoServiceImpl{}
	//DemoVideos0 := vsi.FavorVideoList()
	vsi.FavorVideoList()
	userInfo := usi.UserInfo(token)
	if userInfo != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: DemoVideos,
		})
	}
}
