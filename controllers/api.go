package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
	"github.com/painterQ/painterBlog/models"
	"net/http"
	"strconv"
	"time"
)

type APIController struct {
	beego.Controller
}

func (a *APIController) URLMapping() {
	// 更新账号信息
	a.Mapping("ApiAccount", a.ApiAccount) //account
	// 更新博客信息
	a.Mapping("ApiBlog", a.ApiBlog) //blog
	// 更新密码
	a.Mapping("ApiPassword", a.ApiPassword) //password
	// 删除文章
	a.Mapping("ApiPostDelete", a.ApiPostDelete) //post-delete
	// 添加文章
	a.Mapping("ApiPostAdd", a.ApiPostAdd) //post-add
	// 管理文章
	a.Mapping("ApiManagePosts", a.ApiManagePosts) //manage-posts
	// 删除专题
	//a.Mapping("serie-delete", a.ApiSerieDelete)
	//// 添加专题
	//a.Mapping("serie-add", a.ApiSerieAdd)
	//// 专题排序
	//a.Mapping("serie-sort", a.ApiSerieSort)
	//// 删除草稿箱
	//a.Mapping("draft-delete", a.ApiDraftDelete)
	//// 删除回收箱
	//a.Mapping("trash-delete", a.ApiTrashDelete)
	//// 恢复回收箱
	//a.Mapping("trash-recover", a.ApiTrashRecover)
	//// 上传文件
	//a.Mapping("file-upload", a.ApiFileUpload)
	//// 删除文件
	//a.Mapping("file-delete", a.ApiFileDelete)
}

const (
	// 成功
	NOTICE_SUCCESS = "success"
	// 注意
	NOTICE_NOTICE = "notice"
	// 错误
	NOTICE_ERROR = "error"
)

func responseNotice(c *context.Context, typ, content, hl, redrect string) {
	if hl != "" {
		c.SetCookie("notice_highlight", hl, 86400, "/", "", true, false)
	}
	c.SetCookie("notice_type", typ, 86400, "/", "", true, false)
	c.SetCookie("notice", fmt.Sprintf("[\"%s\"]", content), 86400, "/", "", true, false)
	c.Redirect(http.StatusFound, redrect)//c.Request.Referer())
}

//ApiAccount 更新账号信息
// @router /account [post]
func (a *APIController) ApiAccount() {
	email := a.Input().Get("email")
	phoneNumber := a.Input().Get("phoneNumber")
	address := a.Input().Get("address")

	valid := validation.Validation{}
	valid.Email(email, "email").Message("email格式不合法")
	valid.MaxSize(address, 128,"address").Message("address长度不要超过128")
	valid.Phone(phoneNumber, "phoneNumber").Message("phone Number格式不合法")
	if valid.HasErrors() {
		fmt.Println(valid.Errors[0].Message)
		responseNotice(a.Ctx, NOTICE_NOTICE, "参数错误", "","/admin/profile")
		return
	}
	err := models.ManagerSingleCase.NewUpdate().
		Update("email",email).
		Update("phoneNumber",phoneNumber).
		Update("address",address).EndUpdate()
	if err != nil{
		fmt.Println(err)
		responseNotice(a.Ctx, NOTICE_ERROR, "持久化错误", "","/admin/profile")
		return
	}
	responseNotice(a.Ctx, NOTICE_SUCCESS, "更新成功", "","/admin/profile")
}

//ApiBlog 更新博客信息
// @router /blog [post]
func (a *APIController) ApiBlog() {
	bn := a.Input().Get("blogName")
	bt := a.Input().Get("bTitle")
	ba := a.Input().Get("beiAn")
	st := a.Input().Get("subTitle")
	ss := a.Input().Get("seriessay")
	as := a.Input().Get("archivessay")

	valid := validation.Validation{}
	valid.Required(bn, "blogName").Message("缺少blogName")
	valid.Required(bt, "bTitle").Message("缺少bTitle")
	valid.Required(ba, "beiAn").Message("缺少beiAn")
	valid.Required(st, "subTitle").Message("缺少subTitle")
	valid.Required(ss, "seriesSay").Message("缺少seriesSay")
	valid.Required(as, "archivesSay").Message("缺少archivesSay")

	if valid.HasErrors() {
		fmt.Println(valid.Errors[0].Message)
		responseNotice(a.Ctx, NOTICE_NOTICE, "参数错误", "","/admin/profile")
		return
	}
	err := models.BlogSingleCase.NewUpdate().
		Update("blogName",bn).
		Update("bTitle",bt).
		Update("beiAn",ba).
		Update( "subTitle",st).
		Update("seriesSay",ss).
		Update("archivesSay",as).EndUpdate()
	if err != nil{
		fmt.Println(err)
		responseNotice(a.Ctx, NOTICE_ERROR, "持久化错误", "","/admin/profile")
		return
	}
	responseNotice(a.Ctx, NOTICE_SUCCESS, "更新成功", "","/admin/profile")
}

//ApiPassword 更新密码
// @router /password [post]
func (a *APIController) ApiPassword() {

}

//ApiPostDelete 删除文章
// @router /post-delete [post]
func (a *APIController) ApiPostDelete() {

}

//ApiPostAdd 添加文章
// @router /post-add [post]
func (a *APIController) ApiPostAdd() {
	do := a.Input().Get("do") // auto or save or publish
	slug := a.Input().Get("slug")
	title := a.Input().Get("title")
	text := a.Input().Get("text")
	date := a.Input().Get("date")
	serie := a.Input().Get("serie")
	tag := a.Input().Get("tags")
	update := a.Input().Get("update")
	cidStr := a.Input().Get("cid"))

	defer func() {
		switch do {
		case "auto": // 自动保存
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"fail": FAIL, "time": time.Now().Format("15:04:05 PM"), "cid": cid})
				return
			}
			c.JSON(http.StatusOK, gin.H{"success": SUCCESS, "time": time.Now().Format("15:04:05 PM"), "cid": cid})
		case "save", "publish": // 草稿，发布
			if err != nil {
				responseNotice(c, NOTICE_NOTICE, err.Error(), "")
				return
			}
			uri := "/admin/manage-draft"
			if do == "publish" {
				uri = "/admin/manage-posts"
			}
			c.Redirect(http.StatusFound, uri)
		}
	}()

	valid := validation.Validation{}
	valid.Required(do, "do").Message("缺少do") // auto or save or publish
	valid.Required(title, "title").Message("缺少title")
	valid.Required(update, "update").Message("缺少update")
	valid.Required(date, "date").Message("缺少date")
	valid.Required(serie, "serie").Message("缺少serie")

	if valid.HasErrors() {
		fmt.Println(valid.Errors[0].Message)
		responseNotice(a.Ctx, NOTICE_NOTICE, "参数错误", "","/admin/profile/write-post")
		return
	}

	artc := &models.Article{
		Title:      title,
		Content:    text,
		Slug:       slug,
		CreateTime: CheckDate(date),
		IsDraft:    do != "publish",
		Author:     models.ManagerSingleCase.Username(),
		//todo tags
		SerieID:    0,
		Tags:       []string{"壹","贰"},
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		DeleteTime: time.Time{},
	}

	cid,err := strconv.ParseInt(cidStr,10,64)
	if cid < 1 || err != nil{  //旧文章

	}else {	//新文章

	}

}

//ApiManagePosts 管理文章
// @router /manage-posts [post]
func (a *APIController) ApiManagePosts() {

}


// 检查日期
func CheckDate(date string) time.Time {
	if t, err := time.ParseInLocation("2006-01-02 15:04", date, time.Local); err == nil {
		return t
	}
	return time.Now()
}