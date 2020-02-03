package controllers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/painterQ/painterBlog/models"
	"regexp"
	"strconv"
	"strings"
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
	d.Mapping("UploadImage", d.UploadImage)                   //post	/docs/image/filter 	*
	d.Mapping("GetImageList",d.GetImageList)				  //get		/docs/image/filter 	*
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

//uploadImage 上传图片
//method: Post
//path /docs/image/filter
//data: { img: img, type: blobInfo.blob.type }
//return: {'url':"http://localhost:8080/public/img/background.0ed615ed.jpg"}
// @router /image/filter [post]
func (d *DocumentsController) UploadImage() {
	var para struct {
		Type string `json:"type"` //image/png  image/jpeg   image/jpg
		// dataURL 'data:image/jpeg;base64,'+base64
		Img  string `json:"img"`  //base64
		Name string `json:"name"`
	}
	err := json.Unmarshal(d.Ctx.Input.RequestBody, &para)
	if err != nil {
		responseJson(d.Ctx, err)
		return
	}
	if len(para.Name) > 64 {
		responseJson(d.Ctx, fmt.Errorf("file name is too long"))
		return
	}
	if len(para.Img) > imageSizeLimit {
		responseJson(d.Ctx, fmt.Errorf("file size bigger then %d bytes", imageSizeLimit/4*3))
		return
	}

	imageData, err := base64.StdEncoding.DecodeString(para.Img)
	if err != nil {
		responseJson(d.Ctx, fmt.Errorf("parse error"))
		return
	}

	tmp := strings.Split(para.Type, "/")
	if len(tmp) != 2 {
		responseJson(d.Ctx, fmt.Errorf("parse image type error"))
		return
	}
	info,err := models.ImageStoreSingleCase.SaveImage(imageData, para.Name, tmp[1], nil)
	if err != nil{
		responseJson(d.Ctx, err)
	}
	url := strings.Join([]string{webDN,info.Src}, "/")
	responseJson(d.Ctx, fmt.Sprintf(`{"url":"%s"}`, url))
}

//GetImageList 获取图片列表
//method: Get
//path /docs/image/filter
//para: start、limit
//return: [{"id":"","name":"","type":"","src":"",}]
// @router /image/filter [get]
func (d *DocumentsController) GetImageList() {
	start,err := strconv.Atoi(d.Input().Get("start"))
	if err != nil{
		responseJson(d.Ctx, fmt.Errorf("parameter [start] error:"+err.Error() ))
		return
	}
	limit,err := strconv.Atoi(d.Input().Get("limit"))
	if err != nil{
		responseJson(d.Ctx, fmt.Errorf("parameter [limit] error:"+err.Error() ))
		return
	}
	if limit > 20 {
		responseJson(d.Ctx, fmt.Errorf("parameter [limit] have to little then 20"))
		return
	}
	infos := models.ImageStoreSingleCase.GetAllImageInfo(int64(start),int64(limit))
	for i:= range infos{
		infos[i].Src = strings.Join([]string{webDN,infos[i].Src}, "/")
	}
	fmt.Println(start,limit,infos)

	response,err := json.Marshal(infos)
	if err != nil{
		responseJson(d.Ctx, err)
		return
	}
	responseJson(d.Ctx,response)
}