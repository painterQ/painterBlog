package models

import (
	"encoding/json"
)

//DocumentDataBase document data base
type DocumentDataBase interface {
	Init(dbPath string, hookAfterInitDB func(dbPath string)) (err error)
	Close()

	GetDocument(key []byte) (content []byte, err error)
	GetMate(key []byte, length int) ([]byte, error)
	GetDocumentByTag(tag []string) []string
	//相同的key会覆盖
	Push(content []byte, mate *DocumentMate) (err error)
	//tag
	GetTag() []string
	AddTag([]string) error
}

//DocumentMate mate of document
type DocumentMate struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	SubTitle string   `json:"subTitle"`
	Tags     []string `json:"tags"`
	Attr     int      `json:"attr"`
	LastTime int64    `json:"lastTime"`
	Abstract string   `json:"abstract"`
}

//Encode encode
func (d *DocumentMate) Encode() ([]byte, error) {
	return json.Marshal(*d)
}

//Decode decode
func (d *DocumentMate) Decode(content []byte) error {
	return json.Unmarshal(content, d)
}
