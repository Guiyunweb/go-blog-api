package api

import (
	"blog-api/serializer"
	"blog-api/service"
	"github.com/gin-gonic/gin"
)

func SaveLink(c *gin.Context) {
	var service service.SaveLinkService
	if err := c.ShouldBindJSON(&service); err == nil {
		service.SaveLink()
		c.JSON(200, SuccessResponse())
	} else {
		c.JSON(500, err)
	}
}

func ShowLinkList(c *gin.Context) {
	if links, err := service.ShowLinkList(); err == nil {
		c.JSON(200, serializer.BuildLinksResponse(links))
	} else {
		c.JSON(500, err)
	}
}

func LinkList(c *gin.Context) {
	if links, err := service.LinkList(); err == nil {
		c.JSON(200, serializer.BuildLinksResponse(links))
	} else {
		c.JSON(500, err)
	}
}
