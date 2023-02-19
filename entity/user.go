package entity

//import "github.com/RaymondCode/simple-demo/database"

type UserRegister struct {
	Id       int64  `gorm:"not null; primary_key; AUTO_INCREMENT" json:"id,omitempty"`
	Username string `gorm:"not null; unique" json:"username,omitempty"`
	Password string `gorm:"not null;" json:"password,omitempty"`
}

type UserInfo struct {
	Id              int64  `gorm:"not null; primary_key; AUTO_INCREMENT" json:"id,omitempty"`
	Name            string `gorm:"not null; unique" json:"name,omitempty"`
	FollowCount     int64  `gorm:"not null" json:"follow_count,omitempty"`
	FollowerCount   int64  `gorm:"not null" json:"follower_count,omitempty"`
	IsFollow        bool   `gorm:"not null" json:"is_follow,omitempty"`
	avatar          string `gorm:"not null" json:"avatars,omitempty"`
	backgroundImage string `gorm:"not null" json:"background_image,omitempty"`
	signature       string `gorm:"not null" json:"signature,omitempty"`
	totalFaorited   string `gorm:"not null" json:"total_faorited,omitempty"`
	workCount       int64  `gorm:"not null" json:"work_count,omitempty"`
	favoriteCount   int64  `gorm:"not null" json:"favorite_count,omitempty"`
}
