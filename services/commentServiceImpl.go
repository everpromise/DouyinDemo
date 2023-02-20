package services

import (
	"time"

	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mapper"
)

type CommentServiceImpl struct {
}

func (u CommentServiceImpl) Commenton(Id string, user entity.UserInfo, content string, Video_Id string) entity.Comment {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	newComment := mapper.InsertComment(Id, user, content, timeStr)
	video, err := mapper.SelectVideoById(Video_Id)
	if err != nil {
		video.CommentCount++
	}
	return newComment
}

func (u CommentServiceImpl) DeleteComment(Id string) {
	mapper.DeleteCommentById(Id)
}

func (u CommentServiceImpl) MakeCommentList() []entity.Comment {
	return mapper.AllComments()
}
