package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/painterQ/painterBlog/models"
	"regexp"
	"time"
)

//DocumentsController Documents Controller
type DocumentsController struct {
	beego.Controller
}

// URLMapping /docs
func (d *DocumentsController) URLMapping() {
	d.Mapping("GetDocument", d.GetDocument)                   //get 	/docs
	d.Mapping("GetDocsList", d.GetDocsList)                   //post 	/docs
	d.Mapping("GetDocumentsIDByTags", d.GetDocumentsIDByTags) //get 	/docs/tags
	d.Mapping("PostNewDocument", d.PostNewDocument)           //post 	/docs/doc/filter	*
	d.Mapping("AddTag", d.AddTag)                             //post 	/docs/tag/filter	*
	d.Mapping("GetTags", d.GetTags)                           //get 	/docs/tag
}

//GetDocument 获取文章内容
//method GET
//path /docs
//para id string
//return string
// @router / [get]
func (d *DocumentsController) GetDocument() {
	id := d.Input().Get("doc")
	if id == "" {
		responseJson(d.Ctx, errors.New("para doc error"))
	}
	c, err := models.DocumentDataBaseSingleCase.GetDocument([]byte(id))
	responseJson(d.Ctx, err)
	_, _ = d.Ctx.ResponseWriter.Write(c)
}

//PostDocsList 获取文章元信息
//methos POST
//path /docs
//data {start:"/docc", length: 10}
//return {[
//		{"id":"first","title":"first","subTitle":"blog","tags":["blog","document"],"attr":0,"createTime":"0001-01-01T00:00:00Z","lastTime":"2020-01-26T15:35:30.653602+08:00","abstract":"PHA+Zmlyc3Q8L3A+"},
//				{"id":"first","title":"first","subTitle":"blog","tags":["blog","document"],"attr":0,"createTime":"0001-01-01T00:00:00Z","lastTime":"2020-01-26T15:35:30.653602+08:00","abstract":"PHA+Zmlyc3Q8L3A+"}
//		 ], total: 22}
// @router / [post]
func (d *DocumentsController) GetDocsList() {
	var para struct {
		Start  string `json:"start"`
		Length int    `json:"length"`
	}
	err := json.Unmarshal(d.Ctx.Input.RequestBody, &para)
	if err != nil {
		panic(err)
	}
	c, err := models.DocumentDataBaseSingleCase.GetMate([]byte(para.Start), para.Length)
	if err != nil {
		panic(err)
	}
	_, _ = d.Ctx.ResponseWriter.Write(c)
}

//根据tag获取文章ID
// @router /tags [get]
func (d *DocumentsController) GetDocumentsIDByTags() {

}

//PostNewDocument 发表文章
//method: POST
//path /docs/doc/filter
//data: "{"title":"first","path":"first","document":"<p>第一篇文章</p>\n<p>&nbsp;</p>"}"}
//returm: 200 "{ok:true}"
// @router /doc/filter [post]
func (d *DocumentsController) PostNewDocument() {
	var para struct {
		Title    string   `json:"title"`
		Path     string   `json:"path"`
		Abstract string   `json:"abstract"`
		Tag      []string `json:"tag"`
		Document string   `json:"document"`
	}
	err := json.Unmarshal(d.Ctx.Input.RequestBody, &para)
	if err != nil {
		responseJson(d.Ctx, err)
		return
	}

	if len(para.Title) == 0 || len(para.Path) == 0 || para.Path[0] != '/' {
		responseJson(d.Ctx, errors.New("params have empty member"))
		return
	}

	content := []byte(para.Document)
	abs := para.Abstract
	if len(abs) == 0 {
		abs = getAbstract(para.Document)
	}
	err = models.DocumentDataBaseSingleCase.Push(content, &models.DocumentMate{
		ID:       para.Path,
		Title:    para.Title,
		SubTitle: "blog",
		Tags:     para.Tag,
		LastTime: time.Now().Unix(),
		Attr:     0,
		Abstract: abs,
	})
	if err != nil {
		responseJson(d.Ctx, err)
		return
	}
	_, _ = d.Ctx.ResponseWriter.Write([]byte("{ok:true}"))
}

var r, _ = regexp.Compile(`<p>.+<\\p>`)

func getAbstract(s string) string {
	return r.FindString(s)
}

//GetTags 获取全部tag
//method: GET
//path /docs/tag
//para: nil
//return: ["tag1","tag2","tag3"]
// @router /tag [get]
func (d *DocumentsController) GetTags() {
	//todo 403缓存
	responseJson(d.Ctx, models.DocumentDataBaseSingleCase.GetTag())
}

//AddTag 新增tag
//method: Post
//path /docs/tag/filter
//data: ["tag1","tag2","tag3"]
//return: nil
// @router /tag/filter [post]
func (d *DocumentsController) AddTag() {
	tag := make([]string, 10)
	err := json.Unmarshal(d.Ctx.Input.RequestBody, &tag)
	if err != nil {
		responseJson(d.Ctx, err)
		return
	}
	responseJson(d.Ctx, models.DocumentDataBaseSingleCase.AddTag(tag))
}
