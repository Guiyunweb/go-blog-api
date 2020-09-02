package model

import "time"

type Comment struct {
	Id         int64
	ArticleId  int64 `xorm:"index notnull"`
	Comment    string
	Username   string
	Email      string
	Site       string
	FatherId   int64
	CreateTime time.Time
}
