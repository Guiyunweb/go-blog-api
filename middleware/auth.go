package middleware

import (
	"blog-api/serializer"
	"blog-api/util"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(401, serializer.CheckLogin())
			c.Abort()
			return
		}

		if _, err := util.VerifyToken(token); err == nil {
			c.Next()
			return
		} else {
			c.JSON(401, serializer.CheckLogin())
			c.Abort()
			return
		}
	}
}
