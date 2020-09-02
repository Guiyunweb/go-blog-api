package service

import (
	"blog-api/model"
	"blog-api/serializer"
	"blog-api/util"
	"time"
)

type Pagination struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

type ArticleService struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Tags       []string  `json:"tags"`
	Content    string    `json:"content"`
	Thumbnail  string    `json:"thumbnail"`
	Summary    string    `json:"summary"`
	Issued     bool      `json:"issued"`
	CreateTime time.Time `json:"create_time"`
}

type ArticleInfoServer struct {
	Id int64 `form:"id"`
}

func (s *ArticleService) Save() *serializer.Response {

	article := model.Article{
		Id:         util.GetSnowflakeId(),
		Title:      s.Title,
		Content:    s.Content,
		Thumbnail:  s.Thumbnail,
		Summary:    s.Summary,
		Issued:     s.Issued,
		CreateTime: time.Now(),
	}

	if _, err := model.DB.Insert(&article); err != nil {
		panic(err)
		return &serializer.Response{
			Success: false,
			Message: "保存文章失败",
		}
	}
	return nil
}

func (page Pagination) ArticleList() serializer.Response {
	var articles []model.Article

	totalElements, err := model.DB.Table("article").Count()
	if err != nil {
		return serializer.Response{
			Success: false,
			Message: "数据库错误",
			Data:    nil,
		}
	}

	if err := model.DB.Table("article").Desc("create_time").Limit(page.Size, page.Page).Find(&articles); err == nil {
		return serializer.BuildListResponse(serializer.BuildArticleList(articles), totalElements)
	} else {
		return serializer.Response{
			Success: false,
			Message: "数据库错误",
			Data:    nil,
		}
	}
}

func (page Pagination) ArticleShowList() serializer.Response {
	var articles []model.Article

	totalElements, err := model.DB.Table("article").Where("issued = ?", true).Count()
	if err != nil {
		return serializer.Response{
			Success: false,
			Message: "数据库错误",
			Data:    nil,
		}
	}

	if err := model.DB.Table("article").Where("issued = ?", true).Desc("create_time").Limit(page.Size, page.Page).Find(&articles); err == nil {
		return serializer.BuildListResponse(serializer.BuildArticleList(articles), totalElements)
	} else {
		return serializer.Response{
			Success: false,
			Message: "数据库错误",
			Data:    nil,
		}
	}
}

func (service ArticleInfoServer) ArticleInfo() (model.Article, *serializer.Response) {
	var article model.Article
	if _, err := model.DB.Where("id = ?", service.Id).Get(&article); err == nil {
		return article, nil
	} else {
		return article, &serializer.Response{
			Success: false,
			Message: "获取文章错误",
		}
	}
}

func ArchiveList() ([]model.Article, *serializer.Response) {
	var articles []model.Article
	if err := model.DB.Table("article").Desc("create_time").Find(&articles); err == nil {
		return articles, nil
	} else {
		return articles, &serializer.Response{
			Success: false,
			Message: "数据库错误",
			Data:    nil,
		}
	}
}
