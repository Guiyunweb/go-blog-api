package service

import (
	"blog-api/model"
	"blog-api/serializer"
	"blog-api/util"
	"strconv"
	"time"
)

type SaveCommentService struct {
	Id         int64     `json:"id"`
	ArticleId  string    `json:"articleId"`
	Comment    string    `json:"comment"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Site       string    `json:"site"`
	FatherId   int64     `json:"fatherId"`
	CreateTime time.Time `json:"createDate"`
}

type ShowCommentService struct {
	Id string `form:"id"`
}

func (service SaveCommentService) SaveComment() *serializer.Response {
	articleId, _ := strconv.ParseInt(service.ArticleId, 10, 64)
	comment := model.Comment{
		Id:         util.GetSnowflakeId(),
		ArticleId:  articleId,
		Comment:    service.Comment,
		Username:   service.Username,
		Email:      service.Email,
		Site:       service.Site,
		FatherId:   service.FatherId,
		CreateTime: time.Now(),
	}
	if _, err := model.DB.Insert(&comment); err != nil {
		panic(err)
		return &serializer.Response{
			Success: false,
			Message: "保存评论失败",
		}
	} else {
		return nil
	}
}

func (service ShowCommentService) ShowComment() ([]model.Comment, *serializer.Response) {
	var comment []model.Comment
	if err := model.DB.Where("article_id = ?", service.Id).Find(&comment); err == nil {
		return comment, nil
	} else {
		return comment, &serializer.Response{
			Success: false,
			Message: "查询失败",
			Data:    nil,
		}
	}

}
