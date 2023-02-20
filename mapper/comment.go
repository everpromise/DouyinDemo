package mapper

import (
	"strconv"

	"github.com/RaymondCode/simple-demo/entity"
)

func InsertComment(Id string, user entity.UserInfo, content string, createTime string) entity.Comment {
	Id64, _ := strconv.ParseInt(Id, 10, 64)
	newComment := entity.Comment{
		Id:         Id64,
		User:       user,
		Content:    content,
		CreateDate: createTime,
	}
	Db.Create(&newComment)
	return newComment
}

func DeleteCommentById(comment_id string) {
	comment := entity.Comment{}
	if err := Db.First(&comment, "id = ?", comment_id).Error; err != nil {
		Db.Delete(comment)
	}
}

func AllComments() []entity.Comment {
	var commentlist []entity.Comment
	Db.Find(&commentlist)
	return commentlist
}
