package internal

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

const (
	ScryptInitN  = 262144
	ScryptInitR  = 8
	ScryptInitP  = 4
	ScryptLength = 32
)

const (
	K_title    = "title"
	K_subTitle = "subTitle"
	K_name     = "name"
	K_say      = "say"
	K_ipc      = "ipc"
	K_github   = "github"
	K_mail     = "email"
	K_avatar   = "avatar"
)

type NormalConfig struct {
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	Name     string `json:"name"`
	Say      string `json:"say"`
	IPC      string `json:"ipc"`
	Github   string `json:"github"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

type ScryptJSON struct {
	Salt string `json:"salt"`
	Key  string `json:"key"`
}

type AppConfig struct {
	NormalConfig NormalConfig `json:"normalConfig"`
	PWD          ScryptJSON   `json:"pwd"`
	LastLogin    int64        `json:"lastLogin"`
	Token        string       `json:"token"`

	lock           sync.RWMutex `json:"-"`
	jsonFileHandle *os.File     `json:"-"`
	channel        chan bool
	quit           chan bool
}

func (a *AppConfig) Start(filePath string) error {
	cf, err := os.OpenFile(filePath, os.O_RDWR|os.O_SYNC, 0777)
	if err != nil {
		return fmt.Errorf("open app.json file:" + err.Error())
	}
	a.jsonFileHandle = cf

	jsBytes, err := ioutil.ReadAll(a.jsonFileHandle)
	err = json.Unmarshal(jsBytes, a)
	if err != nil {
		return fmt.Errorf("open app.json file:" + err.Error())
	}

	if len(a.NormalConfig.Avatar) == 0 {
		return fmt.Errorf("init error, get Avatar error : get empty")
	}

	//pwd
	if len(a.PWD.Key) == 0 {
		return fmt.Errorf("init error, get init pwd error : get empty")
	}

	key, isExtKey := a.GetPWD()
	if !isExtKey {
		a.PWD = *(pwd2key(key.Key))
	}

	a.channel = make(chan bool, 128)
	a.quit = make(chan bool)
	go a.write()
	a.channel <- true
	return nil
}

func (a *AppConfig) Close() {
	close(a.quit)
	_ = a.jsonFileHandle.Close()
}

func (a *AppConfig) synchronize() {
	go func() {
		a.channel <- true
	}()
}

func (a *AppConfig) write() {
	for {
		select {
		case <-a.channel:
			length := len(a.channel)
			for i := 0; i < length; i++ {
				<-a.channel
			}

			bs, _ := json.MarshalIndent(a, "", "    ")
			_, _ = a.jsonFileHandle.Seek(0, 0)
			n, _ := a.jsonFileHandle.Write(bs)
			_ = a.jsonFileHandle.Truncate(int64(n))
		case <-a.quit:
			break
		}
	}
}

func (a *AppConfig) GetNormalConfig() *NormalConfig {
	a.lock.RLock()
	defer a.lock.RUnlock()
	ret := new(NormalConfig)
	*ret = a.NormalConfig
	return ret
}

func (a *AppConfig) GetLastLogin() time.Time {
	a.lock.RLock()
	defer a.lock.RUnlock()

	return time.Unix(a.LastLogin, 0)
}

func (a *AppConfig) GetPWD() (*ScryptJSON, bool) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	ret := new(ScryptJSON)
	*ret = a.PWD
	return ret, len(a.PWD.Salt) != 0
}

func (a *AppConfig) SetToken(token string) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.Token = token
	a.synchronize()
}

func (a *AppConfig) GetToken() string {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.Token
}

func (a *AppConfig) SetNormalConfig(in ...string) error {
	if len(in)%2 != 0 {
		return fmt.Errorf("odd input")
	}

	a.lock.Lock()
	defer a.lock.Unlock()

	tmp := a.NormalConfig
	for i := 0; i < len(in); i += 2 {
		key, value := in[i], in[i+1]
		switch key {
		case K_title:
			tmp.Title = value
		case K_subTitle:
			tmp.SubTitle = value
		case K_name:
			tmp.Name = value
		case K_say:
			tmp.Say = value
		case K_ipc:
			tmp.IPC = value
		case K_github:
			tmp.Github = value
		case K_mail:
			tmp.Email = value
		case K_avatar:
			tmp.Avatar = value
		default:
			return fmt.Errorf("unknown key:" + key)
		}
	}

	a.NormalConfig = tmp
	a.synchronize()
	return nil
}

func (a *AppConfig) SetLastLogin(in time.Time) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.LastLogin = in.Unix()
	a.synchronize()
}

//SetPWD it will clear token
func (a *AppConfig) SetPWD(in *ScryptJSON) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.PWD = *in
	a.Token = ""
	a.synchronize()
}

//输出是json编码
func pwd2key(pwd string) *ScryptJSON {
	s := make([]byte, 32)
	_, _ = rand.Read(s)
	key, _ := scrypt.Key([]byte(pwd), s, ScryptInitN, ScryptInitR, ScryptInitP, ScryptLength)
	ret := new(ScryptJSON)
	ret.Key = base64.URLEncoding.EncodeToString(key)
	ret.Salt = base64.URLEncoding.EncodeToString(s)
	return ret
}
