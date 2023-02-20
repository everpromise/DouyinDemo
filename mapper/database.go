package mapper

import (
	"fmt"

	"github.com/RaymondCode/simple-demo/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() {
	var err error
	Db, err = gorm.Open(mysql.Open("root:ginlater0705@(127.0.0.1:3306)/douyin_demo?charset=utf8mb4"), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败！")
		panic(err)
	}

	// 建表
	Db.AutoMigrate(&entity.UserRegister{})
	Db.AutoMigrate(&entity.UserInfo{})
	Db.AutoMigrate(&entity.Video{})
	Db.AutoMigrate(&entity.Message{})
	Db.AutoMigrate(&entity.Comment{})
	//Db.AutoMigrate(&entity.VideoResponse{})
}
