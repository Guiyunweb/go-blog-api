package serializer

import (
	"blog-api/model"
	"strconv"
	"time"
)

type Comment struct {
	Id         string    `json:"id"`
	ArticleId  string    `json:"articleId"`
	Comment    string    `json:"comment"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Site       string    `json:"site"`
	FatherId   int64     `json:"fatherId"`
	CreateTime time.Time `json:"createTime"`
}

func BuildComment(comment model.Comment) Comment {
	return Comment{
		Id:         strconv.FormatInt(comment.Id, 10),
		ArticleId:  strconv.FormatInt(comment.ArticleId, 10),
		Comment:    comment.Comment,
		Username:   comment.Username,
		Email:      comment.Email,
		Site:       comment.Site,
		FatherId:   comment.FatherId,
		CreateTime: comment.CreateTime,
	}
}

func BuildComments(items []model.Comment) (comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		comments = append(comments, comment)
	}
	return comments
}

func BuildCommentsResponse(comments []model.Comment) Response {
	return Response{
		Success: true,
		Message: "查询成功",
		Data:    BuildComments(comments),
	}
}
