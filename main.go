package main

import (
	"./router"
	"./db"
)
func main() {
	//连接数据库
	db.InitMongodb()

	router := router.InitRouter()
	router.Run()
}
