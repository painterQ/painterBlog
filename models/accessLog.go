package models

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/util"
	"net"
	"path"
	"sync"
	"sync/atomic"
	"time"
)

var AccessLogSingleCase AccessLog

const AccessDBPath = "./access"

func init() {
	AccessLogSingleCase = new(AccessLogLevelDBImpl)
	dbPath := beego.AppConfig.DefaultString(ConfigDBPath, DefaultDBPathConfig)
	dbPath = path.Join(dbPath, AccessDBPath)
	err := AccessLogSingleCase.Start(dbPath)
	if err != nil {
		panic(err)
	}
}

type Access struct {
	IP    net.IP
	time  time.Time
	docID string
}

type Login struct {
	IP   net.IP
	time time.Time
}

type AccessLog interface {
	Login(net.IP)
	Access(ip net.IP, docID string)
	GetLatestAccess() []string
	GetLatestLogin() []string
	Start(path string) error
	Stop()
}

var loginKey = []byte("login")
var accessKey = []byte("access")

type AccessLogLevelDBImpl struct {
	db     *leveldb.DB
	inited int32
}

func (a *AccessLogLevelDBImpl) Login(ip net.IP) {
	go func() {
		key := bytes.Join([][]byte{loginKey, []byte(ip.String())}, nil)
		l := Login{
			IP:   ip,
			time: time.Now(),
		}
		js, _ := json.Marshal(l)
		_ = a.db.Put(key, js, nil)
	}()
}

func (a *AccessLogLevelDBImpl) Access(ip net.IP, docID string) {
	go func() {
		key := bytes.Join([][]byte{accessKey, []byte(ip.String())}, nil)
		access := Access{
			IP:    ip,
			time:  time.Now(),
			docID: docID,
		}
		js, _ := json.Marshal(access)
		_ = a.db.Put(key, js, nil)
	}()
}

func (a *AccessLogLevelDBImpl) GetLatestAccess() []string {
	return a.get(accessKey)
}

func (a *AccessLogLevelDBImpl) get(prefix []byte) []string {
	iter := a.db.NewIterator(util.BytesPrefix(prefix), nil)
	defer iter.Release()
	var ret = make([]string, 0, 20)
	i := 20
	for iter.Next() && i > 0 {
		i--
		v := iter.Value()
		ret = append(ret, string(v))
	}
	return ret
}

func (a *AccessLogLevelDBImpl) GetLatestLogin() []string {
	return a.get(loginKey)
}

func (a *AccessLogLevelDBImpl) Start(path string) error {
	if !atomic.CompareAndSwapInt32(&a.inited, 0, 1) {
		return errors.New("already init")
	}
	a.db = openDB(path, nil, false, nil)
	if a.db == nil {
		return errors.New("open db " + path + " error")
	}
	return nil
}

func (a *AccessLogLevelDBImpl) Stop() {
	_ = a.db.Close()
	atomic.StoreInt32(&a.inited, 0)
}

//************登录控制，同一个IP只能登录五次，否则一段时间改IP不能使用************

type entry struct {
	key   string
	value bool
}

type LoginTimes struct {
	ipList  [5]entry
	pointer int

	atom int32 //0 空闲，1使用

	banPool        [100]string
	pointerBanPool int
}

var LoginTimesLog LoginTimes
var once sync.Once

func (l *LoginTimes) Push(ip net.IP, success bool) bool {
	once.Do(l.clear)
	for {
		if atomic.CompareAndSwapInt32(&(l.atom), 0, 1) {
			break
		}
	}
	defer atomic.StoreInt32(&(l.atom), 0)

	for i := range l.banPool {
		if l.banPool[i] == ip.String() {
			return false
		}
	}

	l.ipList[l.pointer].key = ip.String()
	l.ipList[l.pointer].value = success

	l.pointer++
	l.pointer %= 5

	j := 0
	for i := 0; i < 5; i++ {
		if l.ipList[i].key == l.ipList[0].key && !l.ipList[i].value {
			j++
		}
	}

	ban := j < 5
	if !ban {
		l.banPool[l.pointerBanPool] = ip.String()
		l.pointerBanPool++
		l.pointerBanPool %= 100
	}
	return ban
}

func (l *LoginTimes) clear() {
	go func() {
		for {
			<-time.Tick(24 * time.Hour)
			for i := range l.banPool {
				l.banPool[i] = ""
			}
			logs.Warn("Clear IP blacklist")
		}
	}()
}
