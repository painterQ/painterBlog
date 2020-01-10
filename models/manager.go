package models

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net"
	"sync"
	"time"
)

type SSHKey struct {
	PublicKey []byte //rsa ecc
	KeyName   string //64byte之内
	LastTime  time.Time
	LastIP    net.IP
}

type Twitter struct { // twitter信息
	Card    string
	Site    string
	Image   string
	Address string
}

var ManagerSingleCase *Manager

type ManagerUpdater map[string]interface{}

type Manager struct {
	// 账户名
	username string
	// 账户
	email string
	// 手机号
	phoneN string
	// 住址
	address string
	//最后使用的公钥
	password string
	//Twitter
	twitter *Twitter
	// 博客信息
	lastTime time.Time
	// 最后登录ip
	loginIP net.IP
	// 当前token
	token string
	lock  sync.Locker
}

func (m *Manager) Twitter() *Twitter {
	return m.twitter
}

func (m *Manager) SetTwitter(twitter *Twitter) {
	m.twitter = twitter
}

func (m *Manager) Token() string {
	return m.token
}

func (m *Manager) SetToken(token string) {
	m.token = token
}

func (m *Manager) LoginIP() net.IP {
	return m.loginIP
}

func (m *Manager) SetLoginIP(loginIP net.IP) {
	m.loginIP = loginIP
}

func (m *Manager) LastTime() time.Time {
	return m.lastTime
}

func (m *Manager) SetLastTime(lastTime time.Time) {
	m.lastTime = lastTime
}

func (m *Manager) Password() string {
	return m.password
}

func (m *Manager) SetPassword(password string) {
	m.password = password
}

func (m *Manager) Address() string {
	return m.address
}

func (m *Manager) SetAddress(address string) {
	m.address = address
}

func (m *Manager) PhoneN() string {
	return m.phoneN
}

func (m *Manager) SetPhoneN(phoneN string) {
	m.phoneN = phoneN
}

func (m *Manager) Email() string {
	return m.email
}

func (m *Manager) SetEmail(email string) {
	m.email = email
}

func (m *Manager) Username() string {
	return m.username
}

func (m *Manager) SetUsername(username string) {
	m.username = username
}

func (m Manager) NewUpdate() ManagerUpdater {
	logs.Error("Manager update")
	return make(map[string]interface{}, 9)
}

func (m ManagerUpdater) Update(k string, v interface{}) ManagerUpdater {
	switch k {
	case "username", "email", "phoneNumber", "address", "token":
		t, ok := v.(string)
		if !ok {
			//todo
		}
		m[k] = t
	case "twitter":
		t, ok := v.(*Twitter)
		if !ok {
			//todo
		}
		m[k] = t
	case "lastTime":
		t, ok := v.(time.Time)
		if !ok {
			//todo
		}
		m[k] = t
	case "loginIP":
		t, ok := v.(net.IP)
		if !ok {
			//todo
		}
		m[k] = t
	default:
		m["error"] = fmt.Sprintf("ManagerUpdater, default 分支:%v", k)
	}
	return m
}

func (m ManagerUpdater) EndUpdate() error {
	if err, ok := m["error"]; ok {
		return fmt.Errorf("非法的修改：%v", err)
	}
	ManagerSingleCase.lock.Lock()
	defer ManagerSingleCase.lock.Unlock()
	for k, v := range m {
		switch k {
		case "twitter":
			t, ok := v.(*Twitter)
			if !ok {
				return fmt.Errorf("update username err:(%v)is not string", v)
			}
			ManagerSingleCase.twitter = t
		case "username":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update username err:(%v)is not string", v)
			}
			ManagerSingleCase.username = t
		case "email":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update username err:(%v)is not string", v)
			}
			ManagerSingleCase.email = t
		case "phoneNumber":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update phoneN err:(%v)is not string", v)
			}
			ManagerSingleCase.phoneN = t
		case "address":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update address err:(%v)is not string", v)
			}
			ManagerSingleCase.address = t
		case "token":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update token err:(%v)is not string", v)
			}
			ManagerSingleCase.token = t
		case "lastTime":
			t, ok := v.(time.Time)
			if !ok {
				return fmt.Errorf("update token err:(%v)is not time.Time", v)
			}
			ManagerSingleCase.lastTime = t
		case "loginIP":
			t, ok := v.(net.IP)
			if !ok {
				return fmt.Errorf("update loginIP err:(%v)is not time.Time", v)
			}
			ManagerSingleCase.loginIP = t
		default:
		}
	}
	err := Save(ManagerSingleCase)
	if err != nil {
		//todo
	}
	return nil
}

func (m *Manager) Encode() (v []byte) {
	value := make(map[string]interface{}, 9)
	buf := bytes.NewBuffer(make([]byte, 0, 512))
	value["username"] = m.username
	value["email"] = m.email
	value["phoneN"] = m.phoneN
	value["address"] = m.address
	value["password"] = m.password
	value["lastTime"] = m.lastTime
	value["loginIP"] = m.loginIP
	value["token"] = m.token
	value["twitter"] = m.twitter
	err := gob.NewEncoder(buf).Encode(value)
	if err != nil {
		logs.Error("Manager gob Encode err :", err)
	}
	return buf.Bytes()
}

//len(v) == 0 ===> EOF
func (m *Manager) Decode(v []byte) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("panic when manager decode,%v", r)
		}
	}()
	dc := gob.NewDecoder(bytes.NewBuffer(v))
	value := make(map[string]interface{}, 9)
	err = dc.Decode(&value)
	if err != nil {
		return fmt.Errorf("gob decode err:" + err.Error())
	}
	m.username = value["username"].(string)
	m.email = value["email"].(string)
	m.phoneN = value["phoneN"].(string)
	m.address = value["address"].(string)
	m.password = value["password"].(string)
	m.lastTime = value["lastTime"].(time.Time)
	m.loginIP = value["loginIP"].(net.IP)
	m.token = value["token"].(string)
	tt := value["twitter"].(Twitter)
	m.twitter = &tt
	return nil
}

func (m *Manager) GetKey() (k []byte) {
	return []byte("manager")
}

func (m *Manager) DBName() string {
	return "manager"
}
