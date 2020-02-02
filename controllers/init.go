package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

var (
	webDN          string
	imageSizeLimit int
)

func init() {
	var err error
	webDN = strings.Trim(beego.AppConfig.String("webDN"), "/")
	imageSizeLimit, err = beego.AppConfig.Int("imageSizeLimit")
	if err != nil {
		panic("imageSizeLimit in app.conf is not a number")
	}
}
