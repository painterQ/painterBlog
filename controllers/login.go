package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/painterQ/painterBlog/models"
	"net"
	"path"
	"strings"
)

var accessDBImpl models.AccessLog
var mail string

func init() {
	accessDBImpl = new(models.AccessLogLevelDBImpl)
	dbPath := beego.AppConfig.DefaultString(ConfigDBPath, defaultDBPath)
	dbPath = path.Join(dbPath, AccessDBPath)
	err := accessDBImpl.Start(dbPath)
	if err != nil {
		panic(err)
	}

	mail = beego.AppConfig.String("mail")
}

type LoginController struct {
	beego.Controller
}

func (lc *LoginController) URLMapping() {
	lc.Mapping("Login", lc.Login)
	lc.Mapping("GetAuthorInfo", lc.GetAuthorInfo)
}

//GetAuthorInfo 获取作者信息（author）和博客信息（header）
//method: get
//path /login
//data: nil
//return: {title: "Painter Qiao",
//          subTitle: "for dear & love",
//          avatar: "./avatar.jpeg",
//          lastLogin: 123213213,
//          name: "Painter Qiao",
//          say: "a blog for dear & love"}
// @router / [get]
func (lc *LoginController) GetAuthorInfo() {
	fmt.Println(models.AuthorSingle.Say)
	bs, err := json.Marshal(*models.AuthorSingle)
	fmt.Println(string(bs))
	if err != nil {
		responseJson(lc.Ctx, err)
		return
	}
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
	logs.Error("1")
	var para struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &para)
	if err != nil {
		responseJson(lc.Ctx, fmt.Sprintf(`{"status":0,"message":"%v"}`, err.Error()))
		return
	}
	logs.Error("2")
	//获取IP，并且记录
	fromIP := net.ParseIP(lc.Ctx.Input.IP())
	accessDBImpl.Login(fromIP)	//记录登录日志

	success, token := models.AuthorSingle.Login(para.Password, fromIP)                           //检查密码
	success = strings.ToLower(para.Name) == strings.ToLower(mail) && success //检查mail
	success = models.LoginTimesLog.Push(fromIP, success)                                         //检查错误次数
	logs.Error("3 %v", models.AuthorSingle.Say)
	if success {
		models.SetToken(token, lc.Ctx.ResponseWriter) //设置token
		responseJson(lc.Ctx, `{"status":1}`)
		logs.Info("login success from %v", fromIP.String())
		return
	}

	logs.Info("login fail from %v", fromIP.String())
	responseJson(lc.Ctx, `{"status":0,"message":"账号或者密码错误"}`)
}
