package router

import (
	"github.com/gin-gonic/gin"
	"../controller"
	"../middleware/jwt"
	"./rReadMe"
	"./rMoive"
	"./rPublic"
	"gocl/db"
)

func CreateRouter(database db.GoclDatabase) *gin.Engine {
	router := gin.Default()

	rAdmin := router.Group("/admin")
	{
		rAdmin.Use(jwtauth.JWTAuthRequired())
		{
			rAdmin.POST("/login", controller.LoginEndpoint)
			rAdmin.POST("/submit", controller.LoginEndpoint)
		}
		rAdmin.POST("/register", controller.AddUserEndpoint)
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
