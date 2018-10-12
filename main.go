package main

import (
	"./db"
	"./router"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func App() (*gin.Engine, error) {
	//创建数据库连接
	database, err := db.CreateDataBase()

	if(err == nil) {
		return nil, errors.New("connect database fail")
	}
	//创建路由
	return router.CreateRouter(database), nil
}


func main() {
	//连接数据库
	service, err := App()
	if(err != nil) {
		fmt.Errorf("error:", err)
		return
	}

	service.Run()
}

