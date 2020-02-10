package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

var (
	webDN          string

)

func init() {
	webDN = strings.Trim(beego.AppConfig.String("webDN"), "/")
}
