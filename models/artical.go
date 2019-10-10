package models

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"time"
)

type Article struct {
	// 自增id
	ID int64
	// 作者名
	Author string
	// 标题
	Title string
	// 文章名: how-to-get-girlfriend
	Slug string
	// 评论数量
	Count int
	// markdown文档
	Content string
	// 归属专题
	SerieID int32
	// tagname
	Tags []string
	// 是否是草稿
	IsDraft bool
	// 创建时间
	CreateTime time.Time
	// 更新时间
	UpdateTime time.Time
	// 开始删除时间
	DeleteTime time.Time
	//// 上篇文章
	//Prev *Article `bson:"-"`
	//// 下篇文章
	//Next *Article `bson:"-"`
	// Header
	//Header string `bson:"-"`
	// 预览信息
	Excerpt string `bson:"-"`
	// 一句话描述，文章第一句
	Desc string `bson:"-"`
	// disqus thread
	//Thread string `bson:"-"`
}

type SortArticles []*Article

func (s SortArticles) Len() int           { return len(s) }
func (s SortArticles) Less(i, j int) bool { return s[i].CreateTime.After(s[j].CreateTime) }
func (s SortArticles) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }


func (a *Article) Encode() (v []byte) {
	buf := make([]byte, 0, 512)
	err := gob.NewEncoder(bytes.NewBuffer(buf)).Encode(a)
	if err != nil {
		fmt.Println("####", err)
	}
	return buf
}

func (a *Article) Decode(v []byte)(err error){
	defer func() {
		r := recover()
		if r != nil{
			err = fmt.Errorf("panic when decode,%v",r)
		}
	}()
	dc := gob.NewDecoder(bytes.NewBuffer(v))
	err = dc.Decode(a)
	if err != nil {
		return err
	}
	return nil
}

func (a *Article) GetKey() (k []byte) {
	return []byte(strconv.FormatInt(a.ID,10))
}

func (a *Article) DBName() string {
	return "article"
}