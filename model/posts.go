package model

import "time"

type Posts struct {
	Id              int
	CreateTime      time.Time
	UpdateTime      time.Time
	DisallowComment int
	EditTime        time.Time
	EditorType      int
	FormatContent   string
	MetaDescription string
	MetaKeywords    string
	OriginalContent string
	Password        string
	Slug            string
	Status          int
	Summary         string
	Template        string
	Thumbnail       string
	Title           string
	TopPriority     int
	Url             string
	Visits          string
	WordCount       string
}
