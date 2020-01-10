package models

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/syndtr/goleveldb/leveldb"
	"path"
)

type Storer interface {
	DBName() string
	GetKey() []byte
	Encode() []byte
	Decode([]byte) error
}

var dbPath string

func Save(storer Storer) error {

	db, err := leveldb.OpenFile(path.Join(dbPath, storer.DBName()), nil)

	//延迟关闭db流,必须的操作
	defer db.Close()
	if err != nil {
		fmt.Println("#####", err)
		return err
	}
	key, value := storer.GetKey(), storer.Encode()
	if len(key) == 0 {
		return fmt.Errorf("key is empty")
	}
	//向db中插入键值对
	err = db.Put(key, value, nil)
	if err != nil {
		logs.Error("db put error:" +err.Error())
		return err
	}
	logs.Info("put, key: %v, value:%v", len(key),len(value))
	return nil
}

func Query(rec Storer) (err error) {
	//todo reflect
	//if rec.Type().Kind() != reflect.Ptr {
	//	dec.err = errors.New("gob: attempt to decode into a non-pointer")
	//	return dec.err
	//}
	switch r := (interface{}(rec)).(type) {
	case *Article:
		var db *leveldb.DB
		db, err = leveldb.OpenFile(path.Join(dbPath, r.DBName()), nil)
		if err != nil { //todo leveldb.ErrNotFound
			return err
		}
		defer db.Close()
		var value []byte
		value, err = db.Get(r.GetKey(), nil)
		if err != nil {
			return err
		}
		err = r.Decode(value)
		if err != nil {
			return err
		}
	case *Manager:
		var db *leveldb.DB
		db, err = leveldb.OpenFile(path.Join(dbPath, r.DBName()), nil)
		if err != nil {
			return err
		}
		defer db.Close()
		var value []byte
		value, err = db.Get(r.GetKey(), nil)
		if err != nil {
			return err
		}
		err = r.Decode(value)
		if err != nil {
			return err
		}
	case *Blog:
		var db *leveldb.DB
		db, err = leveldb.OpenFile(path.Join(dbPath, r.DBName()), nil)
		if err != nil {
			return err
		}
		defer db.Close()
		var value []byte
		value, err = db.Get(r.GetKey(), nil)
		if err != nil {
			return err
		}
		err = r.Decode(value)
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown rec type, must be pointer")
	}
	return
}

func QueryArt(s,e int) ([]Article, error) {
	var db *leveldb.DB
	db, err := leveldb.OpenFile(path.Join(dbPath, "article"), nil)
	if err != nil { //todo leveldb.ErrNotFound
		logs.Error("query article err:", err)
		return nil, err
	}
	defer db.Close()
	sByte,eByte := make([]byte,8),make([]byte,8)
	binary.BigEndian.PutUint64(sByte,uint64(s))
	binary.BigEndian.PutUint64(eByte,uint64(e))
	//iter := db.NewIterator(&util.Range{
	//	Start: sByte,
	//	Limit: eByte,
	//}, nil)

	arts := make([]Article,3,64)

	//v ,i := iter.Value() ,0
	//logs.Error("查询文章：",iter.Error())
	//for iter.Valid(){
	//	if iter.Last() {
	//		break
	//	}
	//	if len(v) == 0 {
	//		continue
	//	}
	//	logs.Error("查询文章：",i)
	//	arts = append(arts ,new(Article))
	//	err = arts[i].Decode(v)
	//	i ++
	//	if err != nil {
	//		return nil, err
	//	}
	//	iter.Next()
	//	v = iter.Value()
	//}
	return arts, nil
}
