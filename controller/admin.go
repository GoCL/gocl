package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../model"
	"../middleware/jwt"
)

func LoginEndpoint(c *gin.Context) {
	username, u:= c.GetPostForm("username")
	password, p := c.GetPostForm("password")

	if (!p || !u) {
		c.JSON(http.StatusOK, gin.H{
			"message": "登陆参数缺失，请检查参数",
		})
		return
	}

	err, user := model(username, password)

	if (err != nil) {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户不存在",
		})
		return
	}

	jwtUser , _ := c.Get("JWTUser")
	token, err := jwtUser.(jwtauth.JWTAuthCallback).NewToken(int(user.Uid.Pid()), user.Username, user.Email, "token")

	if (err != nil) {
		c.JSON(http.StatusOK, gin.H{
			"message": "token 创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登陆成功",
		"result": gin.H{
			"token": token,
		},
	})
}

func AddUserEndpoint(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email    := c.PostForm("email")

	err := model.AddUser(username, password, email)
	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "注册用户异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "注册用户成功",
	})
}
