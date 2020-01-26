package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/painterQ/painterBlog/models"
	"time"
)

//DocumentsController Documents Controller
type DocumentsController struct {
	beego.Controller
}

var dbImpl *models.DocumentLevelDB

func init() {
	dbImpl = new(models.DocumentLevelDB)
	err := dbImpl.Init("/tmp/123123", nil)
	if err != nil {
		panic(err)
	}
}

// URLMapping /docs
func (d *DocumentsController) URLMapping() {
	d.Mapping("GetDocument", d.GetDocument)
	d.Mapping("PostDocsList", d.PostDocsList)
	d.Mapping("GetDocumentsIDByTags", d.GetDocumentsIDByTags)
}

//GetDocument GetDocument
// @router / [get]
func (d *DocumentsController) GetDocument() {
	id := d.Input().Get("id")
	if id == "" {

	}
	c, err := dbImpl.GetDocument([]byte(id))
	if err != nil {
		panic(err)
	}
	_, _ = d.Ctx.ResponseWriter.Write(c)
}

//PostDocsList PostDocsList
// @router / [post]
func (d *DocumentsController) PostDocsList() {
	fmt.Println("###GetDocsList")
	var para struct {
		Start  string `json:"start"`
		Length int    `json:"length"`
	}
	err := json.Unmarshal(d.Ctx.Input.RequestBody, &para)
	if err != nil {
		panic(err)
	}
	c, err := dbImpl.GetMate([]byte(para.Start), para.Length)
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
//path /docs/doc
//data: "{"title":"first","path":"first","document":"<p>第一篇文章</p>\n<p>&nbsp;</p>"}"}
//returm: 200 "{ok:true}"
// @router /doc [post]
func (d *DocumentsController) PostNewDocument() {
	fmt.Println("###PostNewDocument")
	var para struct {
		Tital    string `json:"title"`
		Path     string `json:"path"`
		Document string `json:"document"`
	}
	err := json.Unmarshal(d.Ctx.Input.RequestBody, &para)
	if err != nil {
		panic(err)
	}
	content := []byte(para.Document)
	path := []byte(para.Path)
	abs := para.Document
	if len(abs) > 20 {
		abs = abs[:20] + "..."
	}
	err = dbImpl.Push(path, content, &models.DocumentMate{
		ID:       para.Path,
		Title:    para.Tital,
		SubTitle: "blog",
		Tags:     []string{"blog", "document"},
		LastTime: time.Now(),
		Abstract: abs,
	})
	if err != nil {
		panic(err)
	}
	_, _ = d.Ctx.ResponseWriter.Write([]byte("{ok:true}"))
}

//------------------tool------------------
func (d *DocumentsController) ErrResponse(code int, err error) {
	if code == 401 {

	}
}

func (d *DocumentsController) Cache() bool {
	//d.Ctx.Request.Header.Get("If-Modified-Since")
	//d.Ctx.Request.Header.Get()
	return true
}
