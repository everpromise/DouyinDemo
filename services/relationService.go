package services

import "github.com/RaymondCode/simple-demo/entity"

type Relation interface {
	// 关注(取消关注) 功能
	Follow(userId string, toUserID string, actionType string)

	// 关注列表
	FollowList(userId string) []entity.UserInfo

	// 粉丝列表
	FollowerList(userId string) []entity.UserInfo

	// 好友列表
	FriendList(userId string) []entity.UserInfo
}
