package services

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mapper"
	"time"
)

type MessageServiceImpl struct {
}

func (u MessageServiceImpl) ChatList(userId string, toUserId string) []entity.Message {
	return mapper.SelectMessageList(userId, toUserId)
}

func (u MessageServiceImpl) ChatTo(userId string, toUserId string, content string) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	mapper.InsertMessage(userId, toUserId, content, timeStr)
}
