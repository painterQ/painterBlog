package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/painterQ/painterBlog/models"
	"net"
	"strings"
)

type LoginController struct {
	beego.Controller
}

func (lc *LoginController) URLMapping() {
	lc.Mapping("Login", lc.Login)                     //post	/login
	lc.Mapping("GetAuthorInfo", lc.GetAuthorInfo)     //get	/login
	lc.Mapping("ChangeBaseInfo", lc.ChangeBaseInfo)   //post	/login/base/filter
	lc.Mapping("ChangeBlogInfo", lc.ChangeBlogInfo)   //post	/login/blog/filter
	lc.Mapping("ChangePwdChange", lc.ChangePwdChange) //post	/login/pwd/filter
}

//GetAuthorInfo 获取作者信息（author）和博客信息（header）
//method: get
//path /login
//data: nil
//{title: "Painter Qiao",
//          subTitle: "for dear & love",
//          avatar: "./avatar.jpeg",
//          lastLogin: 123213213,
//          name: "Painter Qiao",
//			ipc: "",
//			github: "",
//          say: "a blog for dear & love"
//          email: ""}
// @router / [get]
func (lc *LoginController) GetAuthorInfo() {
	bs, _ := models.AuthorSingleCase.MarshalJSON(webDN)
	responseJson(lc.Ctx, bs)
}

//Post post
//method	POST
//path		/login
//data		{name:“admin”,password:"admin"}
//return  	{status: 1,message: '登录成功'}
//			{status: 0,message: '账号或者密码错误'}
// @router / [post]
func (lc *LoginController) Login() {
	var para struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &para)
	if err != nil {
		responseJson(lc.Ctx, fmt.Sprintf(`{"status":0,"message":"%v"}`, err.Error()))
		return
	}

	//获取IP，并且记录
	fromIP := net.ParseIP(lc.Ctx.Input.IP())
	models.AuthorSingleCase.Login(para.Password,fromIP) //记录登录日志
	var token string
	var success bool

	defer func() {
		if success {
			setToken(token, lc.Ctx.ResponseWriter) //设置token
			responseJson(lc.Ctx, `{"status":1}`)
			logs.Info("login success from %v", fromIP.String())
		} else {
			logs.Info("login fail from %v", fromIP.String())
			responseJson(lc.Ctx, `{"status":0,"message":"账号或者密码错误"}`)
		}
	}()

	success = strings.ToLower(para.Name) == strings.ToLower(models.AuthorSingleCase.GetEmail())
	if !success {
		logs.Error("login mail error")
		return
	}

	success, token = models.AuthorSingleCase.Login(para.Password, fromIP) //检查密码
	if !success {
		fmt.Println("pwd error", para.Password)
		return
	}
	//success = models.LoginTimesLog.Push(fromIP, success)                                         //检查错误次数

}

//method	POST
//path		/login/base/filter
//data		{mail: "", github: "",}
//return  	200
// @router /base/filter [post]
func (lc *LoginController) ChangeBaseInfo() {
	var para struct {
		Mail   string `json:"mail"`
		Github string `json:"github"`
	}
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &para)
	if err != nil {
		responseJson(lc.Ctx, err)
		return
	}
	err = models.AuthorSingleCase.SetNormal("github", para.Github)
	if err != nil {
		responseJson(lc.Ctx, err)
		return
	}
	err = models.AuthorSingleCase.SetEmail(para.Mail)
	responseJson(lc.Ctx, err)
}

//method	POST
//path		/login/blog/filter
//data		{ name: "", title: "", subTitle: "", IPC: ""}
//return  	200
// @router /blog/filter [post]
func (lc *LoginController) ChangeBlogInfo() {
	var para struct {
		Name     string `json:"name"`
		Title    string `json:"title"`
		SubTitle string `json:"subTitle"`
		IPC      string `json:"ipc"`
	}
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &para)
	if err != nil {
		responseJson(lc.Ctx, err)
		return
	}
	err = models.AuthorSingleCase.SetNormal(
		"name", para.Name, "title", para.Title, "subTitle", para.SubTitle, "ipc", para.IPC)
	if err != nil {
		responseJson(lc.Ctx, fmt.Errorf("set config error: %s", err.Error()))
		return
	}
	responseJson(lc.Ctx, nil)
}

//method	POST
//path		/login/pwd/filter
//data		{pwd: ""}
//return  	200
// @router /pwd/filter [post]
func (lc *LoginController) ChangePwdChange() {
	var para struct {
		PWD string `json:"pwd"`
	}
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &para)
	if err != nil {
		responseJson(lc.Ctx, err)
		return
	}
	models.AuthorSingleCase.ChangePwd(para.PWD)
	responseJson(lc.Ctx, nil)
}
