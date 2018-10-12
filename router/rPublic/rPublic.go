package rPublic

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RPublic(rGroup *gin.RouterGroup)  *gin.RouterGroup{
	rPublic := rGroup.Group("/readme")
	{
		rPublic.GET("/image", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "/images",
			})
		})

		rPublic.GET("/music", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message" : "/music",
			})
		})

		rPublic.GET("/video", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message" : "/video",
			})
		})
	}

	return rPublic
}