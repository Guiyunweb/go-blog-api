package server

import (
	"blog-api/api"
	"blog-api/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 设置跨域问题
	r.Use(middleware.Cors())

	// 路由
	v1 := r.Group("")
	{
		// 无需登录的接口
		v1.POST("/auth/register", api.UserRegister)
		v1.POST("/auth/login", api.UserLogin)

		// 获取文章详情与文章列表
		v1.GET("/article/getInfo", api.ArticleInfo)
		v1.GET("/article/showList", api.ArticleShowList)
		// 获得文章归档
		v1.GET("/article/archive", api.ArchiveList)
		// 友情链接 添加与获取
		v1.POST("/link/save", api.SaveLink)
		v1.GET("/link/showList", api.ShowLinkList)
		// 评论查看
		v1.POST("/comment/save", api.SaveComment)
		v1.GET("/comment/show", api.ShowComment)

		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// 用户
			auth.POST("/user/userInfo", api.GetUserInfo)
			// 文章
			auth.POST("/article/save", api.SaveArticle)
			auth.GET("/article/list", api.ArticleList)

			// 友链
			auth.GET("/link/list", api.LinkList)
		}
	}

	return r
}
