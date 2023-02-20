package mapper

import (
	"strconv"

	"github.com/RaymondCode/simple-demo/entity"
)

func InsertMessage(userId string, toUserId string, content string, createTime string) {
	userId64, _ := strconv.ParseInt(userId, 10, 64)
	toUserId64, _ := strconv.ParseInt(toUserId, 10, 64)
	newMessage := entity.Message{
		ToUserId:   toUserId64,
		FromUserId: userId64,
		Content:    content,
		CreateTime: createTime,
	}
	Db.Create(&newMessage)
}

func SelectMessageList(userId string, toUserId string) []entity.Message {
	var messageList []entity.Message
	Db.Where("from_user_id=?", userId).Or("to_user_id=?", toUserId).Find(&messageList)
	return messageList
}
