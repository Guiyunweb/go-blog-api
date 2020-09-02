package model

import "time"

type Link struct {
	Id         int64
	Name       string
	Link       string
	Describe   string
	Cover      string
	Review     bool
	CreateTime time.Time
}
