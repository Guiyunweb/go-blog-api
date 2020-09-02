package serializer

import (
	"blog-api/model"
	"strconv"
	"time"
)

type Link struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Link       string    `json:"link"`
	Describe   string    `json:"describe"`
	Cover      string    `json:"cover"`
	Show       bool      `json:"show"`
	CreateTime time.Time `json:"create_time"`
}

type LinkResponse struct {
	Response
	Success bool        `json:"success"`
	Data    ArticleInfo `json:"data"`
}

func BuildLink(link model.Link) Link {
	return Link{
		Id:         strconv.FormatInt(link.Id, 10),
		Name:       link.Name,
		Link:       link.Link,
		Describe:   link.Describe,
		Cover:      link.Cover,
		Show:       link.Review,
		CreateTime: link.CreateTime,
	}
}

func BuildLinkList(items []model.Link) (links []Link) {
	for _, item := range items {
		link := BuildLink(item)
		links = append(links, link)
	}
	return links
}

func BuildLinksResponse(links []model.Link) Response {
	return Response{
		Success: true,
		Message: "查询成功",
		Data:    BuildLinkList(links),
	}
}
