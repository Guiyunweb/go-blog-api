package model

import "time"

type Article struct {
	Id         int64
	Title      string
	Content    string
	Thumbnail  string
	Summary    string
	Issued     bool
	CreateTime time.Time
}
