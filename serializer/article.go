package serializer

import (
	"blog-api/model"
	"strconv"
	"time"
)

// 文章序列器
type Article struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	Thumbnail  string    `json:"thumbnail"`
	Summary    string    `json:"summary"`
	Issued     bool      `json:"issued"`
	CreateTime time.Time `json:"create_time"`
}

type ArticleInfo struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Thumbnail  string    `json:"thumbnail"`
	Summary    string    `json:"summary"`
	Issued     bool      `json:"issued"`
	CreateTime time.Time `json:"create_time"`
}

type ArticleResponse struct {
	Response
	Success bool        `json:"success"`
	Data    ArticleInfo `json:"data"`
}

type Archive struct {
	Archives []Article `json:"archives"`
	Year     int       `json:"year"`
}

type ArchiveResponse struct {
	Response
	Success bool      `json:"success"`
	Data    []Archive `json:"data"`
}

func BuildArticle(article model.Article) Article {
	return Article{
		Id:         strconv.FormatInt(article.Id, 10),
		Title:      article.Title,
		Thumbnail:  article.Thumbnail,
		Summary:    article.Summary,
		Issued:     article.Issued,
		CreateTime: article.CreateTime,
	}
}

func BuildArticleList(items []model.Article) (articles []Article) {
	for _, item := range items {
		article := BuildArticle(item)
		articles = append(articles, article)
	}
	return articles
}

func BuildArticleInfo(article model.Article) ArticleInfo {
	return ArticleInfo{
		Id:         strconv.FormatInt(article.Id, 10),
		Title:      article.Title,
		Content:    article.Content,
		Thumbnail:  article.Thumbnail,
		Summary:    article.Summary,
		Issued:     article.Issued,
		CreateTime: article.CreateTime,
	}
}

func BuildArticleInfoResponse(article model.Article) ArticleResponse {
	return ArticleResponse{
		Success: true,
		Data:    BuildArticleInfo(article),
	}
}

func BuildArchiveResponse(articles []model.Article) ArchiveResponse {
	return ArchiveResponse{
		Success: true,
		Data:    BuildArchive(articles),
	}
}

func BuildArchive(items []model.Article) (archives []Archive) {
	var archive Archive
	for index, item := range items {
		year := item.CreateTime.Year()
		if index == 0 {
			archive.Year = year
			archive.Archives = append(archive.Archives, BuildArticle(item))
		} else if year == archive.Year {
			archive.Archives = append(archive.Archives, BuildArticle(item))
		} else {
			archives = append(archives, archive)
			archive = Archive{}
			archive.Year = year
			archive.Archives = append(archive.Archives, BuildArticle(item))
		}
	}
	archives = append(archives, archive)
	return archives
}
