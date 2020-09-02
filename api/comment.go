package api

import (
	"blog-api/serializer"
	"blog-api/service"
	"github.com/gin-gonic/gin"
)

func SaveComment(c *gin.Context) {
	var service service.SaveCommentService
	if err := c.ShouldBindJSON(&service); err == nil {
		if err := service.SaveComment(); err == nil {
			c.JSON(200, SuccessResponse())
		} else {
			c.JSON(500, err)
		}
	} else {
		c.JSON(500, err)
	}
}
func ShowComment(c *gin.Context) {
	var service service.ShowCommentService
	if err := c.ShouldBindQuery(&service); err == nil {
		if comments, err := service.ShowComment(); err == nil {
			res := serializer.BuildCommentsResponse(comments)
			c.JSON(200, res)
		} else {
			c.JSON(200, comments)
		}
	}
}
