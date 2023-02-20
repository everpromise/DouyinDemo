package controller

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/services"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []entity.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment entity.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	video_id := c.Query("video_id")
	comment_id := c.Query("comment_id")
	var usi services.UserService
	usi = services.UserServiceImpl{}
	userInfo := usi.UserInfo(token)
	var csi services.CommentService
	csi = services.CommentServiceImpl{}
	if userInfo != nil {
		if actionType == "1" {
			text := c.Query("comment_text")
			Comment0 := csi.Commenton(comment_id, *userInfo, text, video_id)
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0, StatusMsg: "ok"},
				Comment: Comment0})
			return
		}
		if actionType == "2" {
			csi.DeleteComment(comment_id)
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	//video_id := c.Query("video_id")
	var usi services.UserService
	usi = services.UserServiceImpl{}
	userInfo := usi.UserInfo(token)
	var csi services.CommentService
	csi = services.CommentServiceImpl{}
	DemoComments := csi.MakeCommentList()
	if userInfo != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0},
			CommentList: DemoComments,
		})
	}
}
