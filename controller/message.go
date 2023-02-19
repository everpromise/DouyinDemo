package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var tempChat = map[string][]Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	Response
	MessageList []entity.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	service := services.MessageServiceImpl{}
	service.ChatTo(token, toUserId, content)

	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	service := services.MessageServiceImpl{}
	chatList := service.ChatList(token, toUserId)

	if len(chatList) != 0 {
		c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0},
			MessageList: chatList})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "no chat message!"})
	}
}
