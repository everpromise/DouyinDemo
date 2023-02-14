package mapper

import (
	"github.com/RaymondCode/simple-demo/entity"
)

func InsertUser(userNew *entity.UserRegister) bool {
	Db.Create(&userNew)
	return true
}

// func DeleteUserById(id string) bool {
// 	userRegister := entity.UserRegister{}
// 	if err := Db.First(&userRegister, "id = ?", id).Error; err != nil {
// 		Db.Delete(&userRegister)
// 		return true
// 	}
// 	return false
// }

func SelectUserById(id string) (entity.UserRegister, error) {
	userRegister := entity.UserRegister{}
	if err := Db.First(&userRegister, "id = ?", id).Error; err != nil {
		return userRegister, err
	}
	return userRegister, nil
}

func SelectUserByName(name string) (*entity.UserRegister, error) {
	userRegister := entity.UserRegister{}
	if err := Db.First(&userRegister, "username = ?", name).Error; err != nil {
		return &userRegister, err
	}
	return &userRegister, nil
}

func InsertUserInfo(userinfo entity.UserInfo) {
	Db.Create(&userinfo)
}

func SelectUserInfoById(id string) *entity.UserInfo {
	userInfo := entity.UserInfo{}
	if err := Db.First(&userInfo, "id = ?", id).Error; err != nil {
		return &userInfo
	}
	return &userInfo
}

func SelectUserInfoByName(name string) (*entity.UserInfo, error) {
	userInfo := entity.UserInfo{}
	if err := Db.First(&userInfo, "username = ?", name).Error; err != nil {
		return &userInfo, err
	}
	return &userInfo, nil
}
