package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
)

func responseJson(ctx *context.Context, obj interface{})  {
	if a ,ok := obj.(string);ok{
		_, _ = ctx.ResponseWriter.Write([]byte(a))
		return
	}

	if a ,ok := obj.([]byte);ok{
		_, _ = ctx.ResponseWriter.Write(a)
		return
	}

	if a ,ok := obj.(error);ok{
		_, _ = ctx.ResponseWriter.Write([]byte(a.Error()))
		return
	}
	byteArray,err := json.Marshal(obj)
	if err != nil{
		panic(err)
	}
	_, _ = ctx.ResponseWriter.Write(byteArray)
}