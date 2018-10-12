package rReadMe

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RReadMe(rGroup *gin.RouterGroup)  *gin.RouterGroup{
	rReadMe := rGroup.Group("/readme")
	{
		rStore := rReadMe.Group("/store")
		{
			rStore.GET("/list", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"result": map[string] string {
						"count": "0",
						"list" : "???",
					},
				})
			})


		}

		rUser := rReadMe.Group("/user")
		{
			rUser.GET("/base", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"message" : "...",
				})
			})
		}
	}

	return rReadMe
}