package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []entity.VideoResponse `json:"video_list,omitempty"`
	NextTime  int64                  `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	fmt.Printf("%v\n", c.Query("latest_time"))
	service := services.VideoServiceImpl{}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: service.VideoList(),
		NextTime:  time.Now().Unix(),
	})
}
