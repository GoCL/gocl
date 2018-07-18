package router

import (
	"github.com/gin-gonic/gin"
	user "gocl_server/router/rUser"
)
func RouterInit() *gin.Engine{
	router := gin.Default()

	rUser := router.Group("/user")
	{
		user.RUser(rUser)
	}

	return  router
}