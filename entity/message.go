package entity

type Message struct {
	Id         int64  `gorm:"not null; primary_key; AUTO_INCREMENT" json:"id,omitempty"`
	ToUserId   int64  `gorm:"not null"  json:"to_user_id,omitempty"`
	FromUserId int64  `gorm:"not null"  json:"from_user_id,omitempty"`
	Content    string `gorm:"not null"  json:"content,omitempty"`
	CreateTime string `gorm:"not null"  json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
