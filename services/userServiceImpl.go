package services

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mapper"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) RegisterUser(userRegister *entity.UserRegister) (int64, error) {
	user, err := mapper.SelectUserByName(userRegister.Username)
	if err != nil { // 用户不存在， 可以注册
		mapper.InsertUser(userRegister) // 添加用户
		userT, _ := mapper.SelectUserByName(userRegister.Username)
		mapper.InsertUserInfo(entity.UserInfo{Id: userT.Id, Name: userRegister.Username})
		return user.Id, err
	}

	return user.Id, nil
}

func (u UserServiceImpl) LoginUser(userLogin *entity.UserRegister) (int64, int64) {
	user, err := mapper.SelectUserByName(userLogin.Username)
	if err != nil { // 用户不存在
		return user.Id, 1
	} else if user.Password != userLogin.Password {
		return user.Id, 2
	}
	return user.Id, 0
}

func (u UserServiceImpl) UserInfo(id string) *entity.UserInfo {
	userInfo := mapper.SelectUserInfoById(id)
	return userInfo
}

func (u UserServiceImpl) VideoFavorAct(video_id string) int64 {
	video, err := mapper.SelectVideoById(video_id)
	if err != nil { //视频不存在
		return 0
	}
	video.IsFavorite = !video.IsFavorite
	return 1
}
