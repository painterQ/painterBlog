package models

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/scrypt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	ConfigNormalConfig = "normalConfig"
	ConfigAvatar       = "avatar"
	ConfigLastLogin    = "lastLogin"
	ConfigPwd          = "pwd"
	ConfigTokenKey     = "tokenKey"

	scryptInitN  = 262144
	scryptInitR  = 8
	scryptInitP  = 4
	scryptLength = 32
)

//AuthorSingle author single case
var AuthorSingle = new(Author)
var tokenKey []byte

func init() {
	var err error
	//last login
	AuthorSingle.LastLogin, err = beego.AppConfig.Int64(ConfigLastLogin)
	if err != nil {
		panic("init error, get AuthorSingle.LastLogin error : " + err.Error())
	}

	//normal config
	var tmp string
	tmp = beego.AppConfig.String(ConfigNormalConfig)

	js := strings.Split(tmp, "|")
	if len(js) != 4 {
		panic("init error, get normal config error : please use 'title | subTitle | name | say'")
	}
	AuthorSingle.Title, AuthorSingle.SubTitle, AuthorSingle.Name, AuthorSingle.Say = js[0], js[1], js[2], js[3]

	//avatar
	AuthorSingle.Avatar = beego.AppConfig.String(ConfigAvatar)
	if AuthorSingle.Avatar == "" {
		panic("init error, get AuthorSingle.Avatar error : get empty")
	}

	//pwd
	pwd := beego.AppConfig.String(ConfigPwd)
	if pwd == "" {
		panic("init error, get init pwd error : get empty")
	}

	init := isKeyJson(pwd)
	if !init {
		AuthorSingle.keyJson = pwd
	} else {
		err = beego.AppConfig.Set(ConfigPwd, pwd2key(pwd))
		if err != nil {
			panic("set init pwd error:" + err.Error())
		}
	}

	//tokenKey
	tokenKey = []byte(beego.AppConfig.String(ConfigTokenKey))
	if len(tokenKey) == 0 {
		panic("get token key error, get empty")
	}
}

type Author struct {
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	Name     string `json:"name"`
	Say      string `json:"say"`

	Avatar    string `json:"avatar"`
	LastLogin int64  `json:"lastLogin"`

	atom     int32  `json:"-"` //空闲0， 正在使用1
	keyJson  string `json:"-"`
	token    string `json:"-"`
	tokenKey string `json:"-"`
}

//SetNormal set normal config
func (a *Author) SetNormal(title, subTitle, name, say string) bool {
	for {
		if atomic.CompareAndSwapInt32(&a.atom, 0, 1) {
			break
		}
	}
	defer atomic.StoreInt32(&a.atom, 1)

	input := strings.Join([]string{title, subTitle, name, say}, "|")

	backup := beego.AppConfig.String(ConfigNormalConfig)
	err := beego.AppConfig.Set(ConfigNormalConfig, input)
	if err != nil {
		_ = beego.AppConfig.Set(ConfigNormalConfig, backup)
		return false
	}

	a.Title = title
	a.SubTitle = subTitle
	a.Name = name
	a.Say = say
	return true
}

//SetAvatar Set Avatar
func (a *Author) SetAvatar(avatar string) bool {
	for {
		if atomic.CompareAndSwapInt32(&a.atom, 0, 1) {
			break
		}
	}
	defer atomic.StoreInt32(&a.atom, 1)

	backup := beego.AppConfig.String(ConfigAvatar)
	err := beego.AppConfig.Set(ConfigAvatar, avatar)
	if err != nil {
		_ = beego.AppConfig.Set(ConfigAvatar, backup)
		return false
	}

	a.Avatar = avatar
	return true
}

//Login login
//todo 当前先传递口令明文，保存scrypt
//return bool value and a token
//如果返回true，这把token设置到cookie即可
func (a *Author) Login(pwd string, ip net.IP) (bool, string) {
	for {
		if atomic.CompareAndSwapInt32(&a.atom, 0, 1) {
			break
		}
	}
	defer atomic.StoreInt32(&a.atom, 0)

	t := time.Now().Unix()
	atomic.StoreInt64(&a.LastLogin, t)

	ret := scryptMatch(pwd, a.keyJson)
	if ret {
		t := time.Now().Unix()
		_ = beego.AppConfig.Set(ConfigLastLogin, strconv.FormatInt(t, 10))
		a.LastLogin = t

		a.token = genToken(ip, tokenKey)
		return true, a.token
	}

	return false, ""
}

func (a *Author) TokenMatch(token string, fromIP net.IP) bool {
	return checkToken(token, tokenKey, &fromIP)
}

//ChangePwd Change pwd
func (a *Author) ChangePwd(pwd string) bool {
	for {
		if atomic.CompareAndSwapInt32(&a.atom, 0, 1) {
			break
		}
	}
	defer atomic.StoreInt32(&a.atom, 1)

	err := beego.AppConfig.Set(ConfigPwd, pwd2key(pwd))
	if err != nil {
		return false
	}

	a.token = ""
	return true
}

//**************tool****************
type scryptJSON struct {
	Salt []byte `json:"salt"`
	Key  []byte `json:"key"`
}

//输出是json编码
func pwd2key(pwd string) string {
	s := make([]byte, 32)
	_, _ = rand.Read(s)
	key, _ := scrypt.Key([]byte(pwd), s, scryptInitN, scryptInitR, scryptInitP, scryptLength)
	ret := new(scryptJSON)
	ret.Key = key
	ret.Salt = s
	code, _ := json.Marshal(*ret)
	return string(code)
}

func scryptMatch(pwd string, key string) bool {
	js := new(scryptJSON)
	err := json.Unmarshal([]byte(key), js)
	if err != nil {
		return false
	}
	_key, _ := scrypt.Key([]byte(pwd), js.Salt, scryptInitN, scryptInitR, scryptInitP, scryptLength)
	return bytes.Equal(_key, js.Key)
}

var regexpKeyJson, _ = regexp.Compile(`\{.*key.*salt.*\}`)

func isKeyJson(key string) bool {
	return regexpKeyJson.MatchString(key)
}
