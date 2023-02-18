package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]entity.UserInfo{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	UserInfo entity.UserInfo `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	//user is valid ? database insert : return
	var usi services.UserService
	usi = services.UserServiceImpl{}
	userNew := entity.UserRegister{Username: username, Password: password}
	userId, err := usi.RegisterUser(&userNew)

	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userId,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//token := username + password

	var usi services.UserService
	usi = services.UserServiceImpl{}
	userTemp := entity.UserRegister{Username: username, Password: password}
	userId, err := usi.LoginUser(&userTemp)

	fmt.Printf("用户id%v\n", userId)
	if err == 1 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exit!"},
		})
	} else if err == 2 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 2, StatusMsg: "password error"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "登陆成功！"},
			UserId:   userId,
			Token:    strconv.Itoa(int(userId)),
		})
	}
}

func UserInfo(c *gin.Context) {
	userId := c.Query("token")

	var usi services.UserService
	usi = services.UserServiceImpl{}
	userInfo := usi.UserInfo(userId)

	if userInfo != nil {
		fmt.Printf("%+v", *userInfo)
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			UserInfo: *userInfo,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist11"},
		})
	}
}
