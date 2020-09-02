package api

import (
	"blog-api/serializer"
	"blog-api/service"
	"github.com/gin-gonic/gin"
)

// 注册
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	// 判断请求格式是否为JSON
	if err := c.ShouldBindJSON(&service); err == nil {

		if user, err := service.Register(); err != nil {
			c.JSON(500, err)
		} else {
			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 登录用户
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	// 判断请求格式是否为JSON
	if err := c.ShouldBindJSON(&service); err == nil {
		if user, err := service.Login(); err != nil {
			c.JSON(500, err)
		} else {
			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

//获得用户详情
func GetUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if user, err := service.UserInfo(token); err != nil {
		c.JSON(500, err)
	} else {
		res := serializer.BuildUserResponse(user)
		c.JSON(200, res)
	}
}
