package services

import "github.com/RaymondCode/simple-demo/entity"

type MessageService interface {
	// 存储聊天记录
	ChatList(userId string, toUserId string) []entity.Message

	// 发送聊天信息
	ChatTo(userId string, toUserId string, content string)
}
