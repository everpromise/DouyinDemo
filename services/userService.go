package services

import "github.com/RaymondCode/simple-demo/entity"

/*
所有 user 相关操作姐接口
*/
type UserService interface {
	// 注册，登陆，查询用户信息
	RegisterUser(userRegister *entity.UserRegister) (int64, error)

	LoginUser(userLogin *entity.UserRegister) (int64, int64)

	UserInfo(id string) *entity.UserInfo

	//赞操作
	VideoFavorAct(video_id string) int64
}
