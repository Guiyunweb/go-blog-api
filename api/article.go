package api

import (
	"blog-api/serializer"
	"blog-api/service"
	"github.com/gin-gonic/gin"
)

// 文章列表
func ArticleList(c *gin.Context) {

	var service service.Pagination
	if err := c.ShouldBindQuery(&service); err == nil {
		res := service.ArticleList()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func ArticleShowList(c *gin.Context) {

	var service service.Pagination
	if err := c.ShouldBindQuery(&service); err == nil {
		res := service.ArticleShowList()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ArchiveList(c *gin.Context) {
	if archiveList, err := service.ArchiveList(); err == nil {
		res := serializer.BuildArchiveResponse(archiveList)
		c.JSON(200, res)
	} else {
		c.JSON(200, err)
	}
}

// 保存文章
func SaveArticle(c *gin.Context) {
	var service service.ArticleService
	if err := c.ShouldBind(&service); err == nil {

		if err := service.Save(); err != nil {
			c.JSON(500, err)
		} else {
			c.JSON(200, SuccessResponse())
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//查询文章详情
func ArticleInfo(c *gin.Context) {
	var service service.ArticleInfoServer
	if err := c.ShouldBindQuery(&service); err == nil {
		if article, err := service.ArticleInfo(); err != nil {
			c.JSON(500, err)
		} else {
			res := serializer.BuildArticleInfoResponse(article)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
