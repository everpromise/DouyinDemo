package main

import (
	"github.com/RaymondCode/simple-demo/mapper"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	go service.RunMessageServer()

	mapper.InitDatabase() // 初始化数据库
	//mapper.Test()

	r := gin.Default()

	r.StaticFS("/public/covers", http.Dir("./public/covers"))
	r.StaticFS("/public/videos", http.Dir("./public/videos"))
	r.StaticFS("/public/avatars", http.Dir("./public/avatars"))
	r.StaticFS("/public/background", http.Dir("./public/background"))

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
