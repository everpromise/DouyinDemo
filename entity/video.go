package entity

type Video struct {
	Id            int64  `gorm:"not null; primary_key; AUTO_INCREMENT" json:"id,omitempty"`
	AuthorId      int64  `gorm:"not null" json:"author_id"`
	PlayUrl       string `gorm:"not null" json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `gorm:"not null" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"not null" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"not null" json:"comment_count,omitempty"`
	IsFavorite    bool   `gorm:"not null" json:"is_favorite,omitempty"`
	Title         string `gorm:"not null" json:"title,omitempty"`
}

type VideoResponse struct {
	Id            int64    `gorm:"not null; primary_key; AUTO_INCREMENT" json:"id,omitempty"`
	Author        UserInfo `gorm:"not null" json:"author"`
	PlayUrl       string   `gorm:"not null" json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string   `gorm:"not null" json:"cover_url,omitempty"`
	FavoriteCount int64    `gorm:"not null" json:"favorite_count,omitempty"`
	CommentCount  int64    `gorm:"not null" json:"comment_count,omitempty"`
	IsFavorite    bool     `gorm:"not null" json:"is_favorite,omitempty"`
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type FeedResponse struct {
	Response
	VideoLists []Video `json:"video_list,omitempty"`
	NextTime   int64   `json:"next_time,omitempty"`
}
