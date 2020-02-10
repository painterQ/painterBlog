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
	if obj == nil {
		_, _ = ctx.ResponseWriter.Write([]byte{})
		return
	}
	ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	switch t := obj.(type) {
	case int:
		switch t {
		case 404:
			ctx.Abort(404, "")
		case 403:
			ctx.Abort(403, "")
		default:
			_, _ = ctx.ResponseWriter.Write([]byte(`{"error":"unknown error"}`))
			ctx.Abort(504,"")
		}
	case []byte:
		_, _ = ctx.ResponseWriter.Write(t)
	case string:
		_, _ = ctx.ResponseWriter.Write([]byte(t))
	case error:
		_, _ = ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"error":"%v"}`, t.Error())))
		ctx.Abort(504, t.Error())
	default:
		byteArray, err := json.Marshal(obj)
		if err != nil {
			_, _ = ctx.ResponseWriter.Write([]byte(`{"error":"unknown error"}`))
			ctx.Abort(504, err.Error())
		}
		_, _ = ctx.ResponseWriter.Write(byteArray)
	}
}

//setToken set token to client
func setToken(token string, w http.ResponseWriter) string {
	cookie := http.Cookie{Name: TokenName, Value: token, Path: "/", MaxAge: 172800} //48hour
	http.SetCookie(w, &cookie)
	return token
}
