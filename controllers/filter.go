package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/painterQ/painterBlog/models"
	"net"
	"net/http"
)

type Filter interface {
	GetPattern() string
	GetPosition() int
	GetFilter() beego.FilterFunc
}

type TokenFilter struct{}

func (TokenFilter) GetPattern() string {
	return `*/filter`
}

func (TokenFilter) GetPosition() int {
	return beego.BeforeRouter
}

func (TokenFilter) GetFilter() beego.FilterFunc {
	return func(ctx *context.Context) {
		ip := net.ParseIP(ctx.Input.IP())
		if !models.AuthorSingleCase.TokenMatch(ctx.Input.Cookie(TokenName), ip){
			logs.Error("access without permission，ip:%v",ctx.Input.IP())
			ctx.Abort(403,"403")
		}
	}
}

func Err403(rw http.ResponseWriter, r *http.Request){
	cookie := http.Cookie{Name: TokenName, Value: "clear", Path: "/", MaxAge: 0}
	http.SetCookie(rw, &cookie)
	_,_ = rw.Write([]byte(`{"err":"access without permission"}`))
}

func Err404(rw http.ResponseWriter, r *http.Request){
	_,_ = rw.Write([]byte(`{"err":"not found"}`))
}