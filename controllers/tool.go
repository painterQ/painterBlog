package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"net/http"
)

const (
	TokenName = "painterBlog"
)

func responseJson(ctx *context.Context, obj interface{}) {
	if obj == nil{
		_, _ =ctx.ResponseWriter.Write([]byte{})
		return
	}
	if a, ok := obj.(string); ok {
		_, _ = ctx.ResponseWriter.Write([]byte(a))
		return
	}

	if a, ok := obj.([]byte); ok {
		_, _ = ctx.ResponseWriter.Write(a)
		return
	}

	if a, ok := obj.(error); ok && a != nil {
		_, _ = ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"error":"%v"}`, a.Error())))
		ctx.Abort(504, a.Error())
		return
	}
	byteArray, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	_, _ = ctx.ResponseWriter.Write(byteArray)
}

//setToken set token to client
func setToken(token string, w http.ResponseWriter) string {
	cookie := http.Cookie{Name: TokenName, Value: token, Path: "/", MaxAge: 172800}	//48hour
	http.SetCookie(w, &cookie)
	return token
}