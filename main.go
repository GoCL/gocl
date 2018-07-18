package main

import "gocl_server/router"

func main() {
	router := router.RouterInit()
	router.Run(":3000")
}
