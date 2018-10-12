package router

import (
	"github.com/gin-gonic/gin"
	"../handler"
	"../middleware/jwt"
	"./rReadMe"
	"./rMoive"
	"./rPublic"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	rAdmin := router.Group("/admin")
	{
		rAdmin.Use(jwtauth.JWTAuthRequired())
		{
			rAdmin.POST("/login", handler.LoginEndpoint)
			rAdmin.POST("/submit", handler.LoginEndpoint)
		}
		rAdmin.POST("/register", handler.AddUserEndpoint)
	}

	rApps := router.Group("/apps", jwtauth.JWTAuth())
	{
		//app:微信小程(read me)
		rReadMe.RReadMe(rApps)

		//app:webapp(movie)
		rMoive.RMovie(rApps)

		//公共资源： 验证来源
		rPublic.RPublic(rApps)
	}

	return  router
}
