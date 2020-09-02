package service

import (
	"blog-api/model"
	"blog-api/serializer"
	"blog-api/util"
	"time"
)

type SaveLinkService struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Link       string    `json:"link"`
	Describe   string    `json:"describe"`
	Cover      string    `json:"cover"`
	Review     bool      `json:"review"`
	CreateTime time.Time `json:"create_time"`
}

func (service SaveLinkService) SaveLink() *serializer.Response {
	link := model.Link{
		Id:         util.GetSnowflakeId(),
		Name:       service.Name,
		Link:       service.Link,
		Describe:   service.Describe,
		Cover:      service.Cover,
		Review:     service.Review,
		CreateTime: time.Now(),
	}
	if _, err := model.DB.Insert(&link); err != nil {
		return &serializer.Response{
			Success: false,
			Message: "保存文章失败",
		}
	} else {
		return nil
	}
}

func ShowLinkList() ([]model.Link, *serializer.Response) {
	var links []model.Link
	if err := model.DB.Where("review = ?", true).Find(&links); err == nil {
		return links, nil
	} else {
		return nil, &serializer.Response{
			Success: false,
			Message: "数据库错误",
			Data:    nil,
		}
	}

}
func LinkList() ([]model.Link, *serializer.Response) {
	var links []model.Link
	if err := model.DB.Table("link").Find(&links); err == nil {
		return links, nil
	} else {
		return nil, &serializer.Response{
			Success: false,
			Message: "数据库错误",
			Data:    nil,
		}
	}
}
