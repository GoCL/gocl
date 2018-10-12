package rMoive

import "github.com/gin-gonic/gin"

func RMovie(rGroup *gin.RouterGroup)  *gin.RouterGroup{
	rMovie := rGroup.Group("/movie")
	{
		rMovie.GET("/test", )
	}

	return rMovie
}