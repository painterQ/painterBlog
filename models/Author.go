package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/painterQ/painterBlog/models/internal"
	"golang.org/x/crypto/scrypt"
	"net"
	"time"
)

var AuthorSingleCase Author

func init() {
	err := AuthorSingleCase.Start("./conf/app.json")
	if err != nil{
		panic(err)
	}
}

type Author struct {
	config internal.AppConfig
}

func (a *Author)Start(confPath string) error{
	return a.config.Start(confPath)
}

func (a *Author)Close(confPath string) {
	a.config.Close()
}

//MarshalJSON Marshaler
//{title: "Painter Qiao",
//          subTitle: "for dear & love",
//          avatar: "./avatar.jpeg",
//          lastLogin: 123213213,
//          name: "Painter Qiao",
//			ipc: "",
//			github: "",
//          say: "a blog for dear & love"
//          email: ""}
func (a *Author) MarshalJSON() ([]byte, error) {
	nc := a.config.GetNormalConfig()
	li := a.config.GetLastLogin()
	s := fmt.Sprintf(`{"title":"%s","subTitle":"%s","name":"%s","say":"%s","ipc":"%s","github":"%s","avatar":"%s","lastLogin": %d,"email":"%s"}`,
		nc.Title, nc.SubTitle, nc.Name, nc.Say, nc.IPC, nc.Github, nc.Avatar, li.Unix(), nc.Email)
	return []byte(s), nil
}

//SetNormal set normal config
func (a *Author) SetNormal(in ...string) error {
	return a.config.SetNormalConfig(in...)
}

//SetAvatar Set Avatar
func (a *Author) SetAvatar(avatar string) error {
	return a.config.SetNormalConfig(internal.K_avatar,avatar)
}

func (a *Author) SetEmail(email string) error{
	return a.config.SetNormalConfig(internal.K_mail,email)
}

//Login login
//todo 当前先传递口令明文，保存scrypt
//return bool value and a token
//如果返回true，这把token设置到cookie即可
func (a *Author) Login(pwd string, ip net.IP) (bool, string) {
	jsonKey, _ := a.config.GetPWD()
	ret := scryptMatch(pwd, jsonKey)
	if ret { //只有成功才能生成token
		a.config.SetLastLogin(time.Now())
		token := a.genToken(ip)
		a.config.SetToken(token)
		return true, token
	}

	return false, ""
}

func (a *Author) TokenMatch(token string, fromIP net.IP) bool {
	if token != a.config.Token {
		return false
	}
	return internal.CheckToken(token, a.tokenKeyFromPWD(), &fromIP)
}

//ChangePwd Change pwd
func (a *Author) ChangePwd(pwd string) {

	salt := make([]byte, 32)
	_, _ = rand.Read(salt)
	key, _ := scrypt.Key([]byte(pwd), salt, internal.ScryptInitN,
		internal.ScryptInitR, internal.ScryptInitP, internal.ScryptLength)

	a.config.SetPWD(&internal.ScryptJSON{
		Salt: base64.URLEncoding.EncodeToString(salt),
		Key:  base64.URLEncoding.EncodeToString(key),
	})
}

//ChangePwd Change pwd
func (a *Author) GetEmail() string {
	return a.config.GetNormalConfig().Email
}

func (a *Author) genToken(ip net.IP) string {
	return internal.GenToken(ip, a.tokenKeyFromPWD())
}

//**************tool****************

func scryptMatch(pwd string, key *internal.ScryptJSON) bool {
	salt, err := base64.URLEncoding.DecodeString(key.Salt)
	if err != nil {
		return false
	}
	_key, _ := scrypt.Key([]byte(pwd), salt, internal.ScryptInitN,
		internal.ScryptInitR, internal.ScryptInitP, internal.ScryptLength)
	return base64.URLEncoding.EncodeToString(_key) == key.Key
}

func (a *Author)tokenKeyFromPWD() []byte {
	keyJSON, _ := a.config.GetPWD()
	return []byte(keyJSON.Salt)
}
