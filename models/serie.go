package models

import "time"

type Serie struct {
	// 自增id
	ID int32
	// 名称unique
	Name string
	// 缩略名
	Slug string
	// 专题描述
	Desc string
	// 创建时间
	CreateTime time.Time
	// 文章
	Articles SortArticles `bson:"-"`
}

type SortSeries []*Serie

func (s SortSeries) Len() int           { return len(s) }
func (s SortSeries) Less(i, j int) bool { return s[i].ID > s[j].ID }
func (s SortSeries) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

