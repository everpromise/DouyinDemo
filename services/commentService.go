package services

import "github.com/RaymondCode/simple-demo/entity"

type CommentService interface {
	//评论视频
	Commenton(Id string, user entity.UserInfo, content string, Video_Id string) entity.Comment
	//删除评论
	DeleteComment(Id string)
	//生成评论列表
	MakeCommentList() []entity.Comment
}
