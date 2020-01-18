package utils

import (
	"encoding/base64"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Str2html(raw string) template.HTML {
	return template.HTML(raw)
}

// DateFormat takes a time and a layout string and returns a string with the formatted date. Used by the template parser as "dateformat"
func DateFormat(t time.Time, layout string) string {
	return t.Format(layout)
}

func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func IsNotZero(t time.Time) bool {
	return !t.IsZero()
}

// cache avatar image
// url: https://<static_domain>/static/img/avatar.png
var avatar string

func GetAvatar(domain string) string {
	if avatar == "" {
		resp, err := http.Get("https://" + domain + "/static/img/avatar.png")
		if err != nil {
			log.Println(err)
			return ""
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return ""
		}

		avatar = "data:" + resp.Header.Get("content-type") + ";base64," + base64.StdEncoding.EncodeToString(data)
	}

	return avatar
}
