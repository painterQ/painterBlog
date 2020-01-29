package models

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

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

	openDB := func(path string) (newDB *leveldb.DB) {
		db, err := leveldb.OpenFile(path, &opt.Options{
			Comparer: new(DocumentComparer),
		})
		switch {
		case errors.IsCorrupted(err):
			logs.Error("database (%v) corrupted, process try to recover: %v", path, err.Error())
			db, err = leveldb.RecoverFile(path, nil)
			if err != nil {
				logs.Error("try to recover fail: %v", path, err.Error())
				return nil
			}
			return db
		case err == os.ErrExist:
			logs.Warn("database (%v) is exist, use old database", path)
			fallthrough
		case err == nil:
			return db
		default:
			logs.Error("opening database (%v) encountered unknown error: %v", path, err.Error())
			return nil
		}
	}

	a.db = openDB(path)
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
