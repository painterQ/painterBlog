package models

import (
	"encoding/json"
	"time"
)

//DocumentDataBase document data base
type DocumentDataBase interface {
	Init(path string, hookAfterInitDB func() error) (err error)
	Close()

	GetDocument(key []byte) (content []byte, err error)
	GetMate(key []byte, from, to int) ([]*DocumentMate, error)
	GetDocumentByTag(tag ...string) [][]byte
	//相同的key会覆盖
	Push(key, content []byte, isDraft bool) error
}

//DocumentMate mate of document
type DocumentMate struct {
	ID         string    `tag:"json:id"`
	Title      string    `tag:"json:title"`
	SubTitle   string    `tag:"json:subTitle"`
	Tags       []string  `tag:"json:tags"`
	Attr       int       `tag:"json:attr"`
	CreateTime time.Time `tag:"json:createTime"`
	LastTime   time.Time `tag:"json:lastTime"`
	Abstract   string    `tag:"json:abstract"`
}

//Encode encode
func (d *DocumentMate) Encode() ([]byte, error) {
	return json.Marshal(*d)
}

//Decode decode
func (d *DocumentMate) Decode(content []byte) error {
	return json.Unmarshal(content, d)
}
