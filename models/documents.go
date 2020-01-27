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
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	SubTitle   string    `json:"subTitle"`
	Tags       []string  `json:"tags"`
	Attr       int       `json:"attr"`
	LastTime   time.Time `json:"lastTime"`
	Abstract   string    `json:"abstract"`
}

//Encode encode
func (d *DocumentMate) Encode() ([]byte, error) {
	return json.Marshal(*d)
}

//Decode decode
func (d *DocumentMate) Decode(content []byte) error {
	return json.Unmarshal(content, d)
}
