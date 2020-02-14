package appConfig

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

var conf = `
{
    "normalConfig": {
        "title": "painter qiao",
        "subTitle": "for dear \u0026 love",
        "name": "painter qiao",
        "say": "a blog for dear \u0026 love",
        "ipc": "浙江备案",
        "github": "https://github.com/painterQ",
        "avatar": "../avatar.jpeg",
        "email": "painterqiao@gmail.com"
    },
    "pwd": {
        "salt": "",
        "key": "123456"
    },
	"token": "",
    "lastLogin": 1580185668,
	"friendLinks":[
        {
            "name":"github",
            "url": "https://www.github.com",
            "header": true
        },{
            "name":"百度",
            "url": "https://www.baidu.com"
        },{
            "name":"知乎",
            "url": "https://www.zhihu.com"
        }
    ]
}
`

var fileName = ""

func TestMain(m *testing.M) {
	var err error
	var f *os.File
	for {
		r := make([]byte, 32)
		_, _ = rand.Read(r)
		fileName = fmt.Sprintf("./%s.json", hex.EncodeToString(r))
		f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_EXCL, 0666)
		if err == nil {
			break
		}
	}
	_, _ = f.WriteString(conf)
	m.Run()
	_ = os.RemoveAll(fileName)

}

func TestFriendLinks(t *testing.T) {
	testLinks :=[10]LinksItem{
		{Name: "1", URL:"http:www.dd1.cc", Header:false},
		{Name: "2", URL:"http:www.dd2.cc", Header:false},
		{Name: "3", URL:"http:www.dd3.cc", Header:false},
		{Name: "4", URL:"http:www.dd4.cc", Header:false},
		{Name: "5", URL:"http:www.dd5.cc", Header:false},
		{Name: "6", URL:"http:www.dd6.cc", Header:false},
		{Name: "7", URL:"http:www.dd7.cc", Header:false},
		{Name: "8", URL:"http:www.dd8.cc", Header:false},
		{Name: "9", URL:"http:www.dd9.cc", Header:false},
		{Name: "10", URL:"http:www0.dd1.cc", Header:false},
	}
	c := new(AppConfig)
	err := c.Start(fileName)
	assert.Nil(t, err)
	defer c.Close()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		for i:=0;i<100;i++{
			r := c.GetLinks()
			assert.True(t,len(r)>=3)
			assert.True(t,len(r)<=13)
		}
		wg.Done()
	}()

	go func() {
		for i:=0;i<100;i++{
			c.AddLinks(&testLinks[i % 10])
		}
		wg.Done()
	}()
	go func() {
		for i:=0;i<100;i++{
			assert.False(t,c.RemoveLinks("www.not.exist"))
		}
		wg.Done()
	}()
	wg.Wait()
	r:=c.GetLinks()
	assert.True(t,len(r) == 13)
}

func TestAppConfig_StartClose(t *testing.T) {
	c := new(AppConfig)
	err := c.Start(fileName)
	assert.Nil(t, err)
	defer c.Close()
}

func TestAppConfig_GetLastLogin(t *testing.T) {
	c := new(AppConfig)
	err := c.Start(fileName)
	assert.Nil(t, err)
	defer c.Close()

	assert.Equal(t, c.GetLastLogin().Unix(), 1580185668)
}

func TestAppConfig_LastLogin(t *testing.T) {
	c := new(AppConfig)
	err := c.Start(fileName)
	assert.Nil(t, err)
	defer c.Close()

	assert.Equal(t, c.GetLastLogin().Unix(), int64(1580185668))
	c.SetLastLogin(time.Date(1970, 1, 1, 0, 0, 37, 0, time.UTC))
	assert.Equal(t, c.GetLastLogin().Unix(), int64(37))
}

func TestAppConfig_PWD(t *testing.T) {
	c := new(AppConfig)
	err := c.Start(fileName)
	assert.Nil(t, err)
	defer c.Close()

	pwd, b := c.GetPWD()
	assert.True(t, len(pwd.Salt) > 0)
	assert.True(t, b)
}

func TestAppConfig_NormalConfig(t *testing.T) {
	c := new(AppConfig)
	err := c.Start(fileName)
	assert.Nil(t, err)

	nc := c.GetNormalConfig()
	assert.Equal(t, "painter qiao", nc.Title)
	assert.Equal(t, "for dear & love", nc.SubTitle)
	assert.Equal(t, "painter qiao", nc.Name)
	assert.Equal(t, "a blog for dear \u0026 love", nc.Say)
	assert.Equal(t, "浙江备案", nc.IPC)
	assert.Equal(t, "https://github.com/painterQ", nc.Github)
	assert.Equal(t, "../avatar.jpeg", nc.Avatar)
	assert.Equal(t, "painterqiao@gmail.com", nc.Email)

	arr := []string{K_title, K_subTitle, K_name, K_say, K_ipc, K_github, K_mail, K_avatar}

	wg := sync.WaitGroup{}
	wg.Add(8)
	for k := 0; k < 8; k++ {
		go func(i int) {
			for j := 0; j < 100; j++ {
				r := make([]byte, 5)
				_, _ = rand.Read(r)
				err := c.SetNormalConfig(arr[i], arr[i]+hex.EncodeToString(r))
				assert.Nil(t, err)
				get := c.GetNormalConfig()
				switch arr[i] {
				case K_title:
					assert.Equal(t, arr[i]+hex.EncodeToString(r), get.Title)
				case K_subTitle:
					assert.Equal(t, arr[i]+hex.EncodeToString(r), get.SubTitle)
				case K_name:
					assert.Equal(t, arr[i]+hex.EncodeToString(r), get.Name)
				case K_say:
					assert.Equal(t, arr[i]+hex.EncodeToString(r), get.Say)
				case K_ipc:
					assert.Equal(t, arr[i]+hex.EncodeToString(r), get.IPC)
				case K_github:
					assert.Equal(t, arr[i]+hex.EncodeToString(r), get.Github)
				case K_mail:
					assert.Equal(t, arr[i]+hex.EncodeToString(r), get.Email)
				case K_avatar:
					assert.Equal(t, arr[i]+hex.EncodeToString(r), get.Avatar)
				}
			}
			wg.Done()
		}(k) //避免闭包
	}
	wg.Wait()
	nc = c.GetNormalConfig()
	assert.True(t, strings.HasPrefix(nc.Avatar, K_avatar))
	assert.True(t, strings.HasPrefix(nc.Title, K_title))
	assert.True(t, strings.HasPrefix(nc.SubTitle, K_subTitle))
	assert.True(t, strings.HasPrefix(nc.Name, K_name))
	assert.True(t, strings.HasPrefix(nc.Say, K_say))
	assert.True(t, strings.HasPrefix(nc.IPC, K_ipc))
	assert.True(t, strings.HasPrefix(nc.Github, K_github))
	assert.True(t, strings.HasPrefix(nc.Email, K_mail))

	c.Close()
	err = c.Start(fileName)
	assert.Nil(t, err)
	c.Close()
}

func Test_Concur(t *testing.T) {
	start := time.Now()

	c := new(AppConfig)
	err := c.Start(fileName)
	assert.Nil(t, err)

	nc := c.GetNormalConfig()
	assert.Equal(t, "painter qiao", nc.Title)
	assert.Equal(t, "for dear & love", nc.SubTitle)
	assert.Equal(t, "painter qiao", nc.Name)
	assert.Equal(t, "a blog for dear \u0026 love", nc.Say)
	assert.Equal(t, "浙江备案", nc.IPC)
	assert.Equal(t, "https://github.com/painterQ", nc.Github)
	assert.Equal(t, "../avatar.jpeg", nc.Avatar)
	assert.Equal(t, "painterqiao@gmail.com", nc.Email)

	pwd, _ := c.GetPWD()
	s, err := base64.URLEncoding.DecodeString(pwd.Salt)
	assert.True(t, len(s) == 32)
	assert.Nil(t, err)

	tm := c.GetLastLogin()
	assert.Equal(t, tm.Unix(), int64(1580185668))

	arr := []string{K_title, K_subTitle, K_name, "time", "pwd"}

	wg := sync.WaitGroup{}
	wg.Add(5)
	for k := 0; k < 5; k++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				r := make([]byte, 5)
				_, _ = rand.Read(r)
				var err error
				switch arr[i] {
				case K_title:
					err = c.SetNormalConfig(K_title, K_title+hex.EncodeToString(r))
				case K_subTitle:
					err = c.SetNormalConfig(K_subTitle, K_subTitle+hex.EncodeToString(r))
				case K_name:
					err = c.SetNormalConfig(K_name, K_name+hex.EncodeToString(r))
				case "time":
					c.SetLastLogin(time.Now())
				case "pwd":
					c.SetPWD(pwd2key(hex.EncodeToString(r)))
				}
				assert.Nil(t, err)
			}
		}(k) //避免闭包
	}
	wg.Wait()
	nc = c.GetNormalConfig()
	assert.True(t, strings.HasPrefix(nc.Title, K_title))
	assert.True(t, strings.HasPrefix(nc.SubTitle, K_subTitle))
	assert.True(t, strings.HasPrefix(nc.Name, K_name))
	tmc := c.GetLastLogin()
	assert.True(t, tmc.After(start))
	pc, _ := c.GetPWD()
	s, err = base64.URLEncoding.DecodeString(pc.Salt)
	assert.True(t, len(s) == 32)
	assert.Nil(t, err)

	c.Close()
	err = c.Start(fileName)
	assert.Nil(t, err)
	c.Close()
}

func BenchmarkAppConfig_pwd2key(b *testing.B) {
	r := make([]byte, 6)
	_, _ = rand.Read(r)
	for i:=0;i<b.N;i++{
		pwd2key(hex.EncodeToString(r))
	}
}