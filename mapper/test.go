package mapper

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
)

func Test() {
	if Db == nil {
		fmt.Printf("db null")
	}
	u1 := entity.UserRegister{Username: "冯12k", Password: "男"}
	Db.Create(u1)
}
