package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
	"time"
)

var macKey = []byte{'1', '1', 1, '1', 1, '1', 1}

type LoginController struct {
	beego.Controller
}

func (lc *LoginController) URLMapping(){
	lc.Mapping("Login", lc.Login)
}

//Post post
//method	POST
//path		/login
//data		{name:“admin”,password:"admin"}
//return  	{status: 1,message: '登录成功'}
//			{status: 0,message: '账号或者密码错误'}
// @router / [post]
func (lc *LoginController) Login() {
	fmt.Println("###/login")
	var para struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &para)
	if err != nil {
		panic(err)
	}
	if para.Name == "admin" {
		responseJson(lc.Ctx, `{status: 1,message: '登录成功'}`)
		return
	}
	responseJson(lc.Ctx, `{status: 0,message: '账号或者密码错误'}`)
	return
}

func CheckToken(ctx *context.Context) bool {
	if ctx.Request.URL.Path == "/admin/login" {
		return true
	}
	return checkToken(ctx)
}

func checkToken(ctx *context.Context) bool {
	tk, ok := ctx.GetSecureCookie(string(macKey), "tk")
	if !ok {
		fmt.Println("@@@1")
		return false
	}
	if !strings.HasPrefix(tk, "painterBlog") || len(tk) < 12 {
		fmt.Println("@@@2")
		return false
	}
	if t, err := time.Parse(time.ANSIC, tk[11:]); err != nil {
		fmt.Println("@@@3", tk, err)
		return false
	} else {
		if time.Now().Sub(t) > time.Hour {
			fmt.Println("@@@4")
			return false
		}
	}
	//err := models.ManagerSingleCase.NewUpdate().
	//	Update("lastTime",time.Now()).EndUpdate()
	//if err != nil {
	//	logs.Error("更新登陆时间失败："+err.Error())
	//}
	return true
}
