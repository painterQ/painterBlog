package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/painterQ/painterBlog/models"
	"github.com/painterQ/painterBlog/models/imageStore"
	"regexp"
	"strconv"
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
	d.Mapping("DeleteDoc",d.DeleteDoc)						  //delete	/docs/doc/filter	*
	d.Mapping("GetTags", d.GetTags)                           //get 	/docs/tag
	d.Mapping("UploadImage", d.UploadImage)                   //post	/docs/image/filter 	*
	d.Mapping("GetImageList", d.GetImageList)                 //get		/docs/image/filter 	*
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

//DeleteDoc 删除文章
//method: DELETE
//path /docs/doc/filter
//data {"id":"/doc1"}
//return 200
// @router /doc/filter [delete]
func (d *DocumentsController) DeleteDoc(){
	var para struct {
		ID  string `json:"id"`
	}
	err := json.Unmarshal(d.Ctx.Input.RequestBody, &para)
	if err != nil {
		panic(err)
	}
	err = models.DocumentDataBaseSingleCase.DeleteDoc([]byte(para.ID))
	responseJson(d.Ctx, err)
}

//GetTags 获取全部tag
//method: GET
//path /docs/tag
//para: nil
//return: {"tag1":["",""],"tag2":["",""],"tag3":[]}
// @router /tag [get]
func (d *DocumentsController) GetTags(){
	m := models.DocumentDataBaseSingleCase.GetTags()
	responseJson(d.Ctx, m)
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
		Attr     int      `json:"attr"`
		SubTitle string   `json:"subTitle"`
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
		SubTitle: para.SubTitle,
		Tags:     para.Tag,
		LastTime: time.Now().Unix(),
		Attr:     para.Attr,
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

var reg, _ = regexp.Compile(`filename="(.+)"`)

//uploadImage 上传图片
//method: Post
//path /docs/image/filter
//data: file binary
//return:
// @router /image/filter [post]
func (d *DocumentsController) UploadImage() {
	f, h, _ := d.GetFile("avatar") //获取上传的文件
	defer func() {
		if err := f.Close(); err != nil {
			logs.Error("close file error: " + err.Error())
		}
	}()
	if len(h.Filename) > 64 {
		responseJson(d.Ctx, fmt.Errorf("file name is too long"))
		return
	}
	fmt.Println("name；", h.Filename)

	info, err := models.ImageStoreSingleCase.SaveImage(f, h.Filename, nil)
	if err != nil {
		responseJson(d.Ctx, err)
		return
	}

	go func() {
		img := new(imageStore.ImageInfo)
		_ = json.Unmarshal(info, img)
		if d.Ctx.Input.Header("avatar") == "true" {
			_ = models.AuthorSingleCase.SetAvatar(img.Src)
		}
	}()

	responseJson(d.Ctx, imageStore.AddWebDN2ImgArray(webDN, [][]byte{info}))
}

//GetImageList 获取图片列表
//method: Get
//path /docs/image/filter
//para: start、limit
//return: [{"id":"","name":"","type":"","src":"",}]
// @router /image/filter [get]
func (d *DocumentsController) GetImageList() {
	start, err := strconv.Atoi(d.Input().Get("start"))
	if err != nil {
		responseJson(d.Ctx, fmt.Errorf("parameter [start] error:"+err.Error()))
		return
	}
	limit, err := strconv.Atoi(d.Input().Get("limit"))
	if err != nil {
		responseJson(d.Ctx, fmt.Errorf("parameter [limit] error:"+err.Error()))
		return
	}
	if limit > 20 {
		responseJson(d.Ctx, fmt.Errorf("parameter [limit] have to little then 20"))
		return
	}
	infos := models.ImageStoreSingleCase.GetAllImageInfo(webDN, int64(start), int64(start+limit))
	if infos == nil {
		responseJson(d.Ctx, errors.New("index error, index start with 1"))
		return
	}
	responseJson(d.Ctx, imageStore.AddWebDN2ImgArray(webDN, infos))
}
