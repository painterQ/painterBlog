package models

import (
	"encoding/json"
	"time"
)

type Document struct {
	ID       string`tag:"json:id"`
	Title    string`tag:"json:title"`
	SubTitle string`tag:"json:subTitle"`
	Tags     []string`tag:"json:tags"`
	Attr     int`tag:"json:attr"`
	CreateTime   time.Time `tag:"json:createTime"`
	LastTime     time.Time`tag:"json:lastTime"`
	Abstract []byte`tag:"json:abstract"`
	NextDoc  string`tag:"json:nextDoc"`
	PrefDoc  string`tag:"json:prevDoc"`
	Content  string `tag:"json:content"`
}

func (d * Document)Encode() ([]byte,error){
	return json.Marshal(*d)
}

func (d * Document)Decode(content []byte) error{
	return json.Unmarshal(content, d)
}