package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	Response
	UserList []entity.UserInfo `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")

	service := services.RelationImpl{}
	service.Follow(token, toUserId, actionType)

	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	userId := c.Query("user_id")

	service := services.RelationImpl{}
	userList := service.FollowList(userId)

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userList,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	userId := c.Query("user_id")

	service := services.RelationImpl{}
	userList := service.FollowerList(userId)

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userList,
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	userId := c.Query("user_id")

	service := services.RelationImpl{}
	userList := service.FriendList(userId)

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userList,
	})
}
