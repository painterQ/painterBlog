package models

import (
	"errors"
	"fmt"
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
	if len(key) == 0{
		return fmt.Errorf("key is empty")
	}
	//向db中插入键值对
	err = db.Put(key, value, nil)
	if err != nil {
		fmt.Println("####", err)
		return err
	}
	return nil
}

func Query(rec Storer) (err error){
	//todo reflect
	//if rec.Type().Kind() != reflect.Ptr {
	//	dec.err = errors.New("gob: attempt to decode into a non-pointer")
	//	return dec.err
	//}
	switch r := (interface{}(rec)).(type) {
	case *Article:
		var db *leveldb.DB
		db, err = leveldb.OpenFile(path.Join(dbPath, r.DBName()), nil)
		if err != nil{  //todo leveldb.ErrNotFound
			return err
		}
		defer db.Close()
		var value []byte
		value, err = db.Get(r.GetKey(),nil)
		if err != nil{
			return err
		}
		err = r.Decode(value)
		if err != nil{
			return err
		}
	case *Manager:
		var db *leveldb.DB
		db, err = leveldb.OpenFile(path.Join(dbPath, r.DBName()), nil)
		if err != nil{
			return err
		}
		defer db.Close()
		var value []byte
		value, err = db.Get(r.GetKey(),nil)
		if err != nil{
			return err
		}
		err = r.Decode(value)
		if err != nil{
			return err
		}
	case *Blog:
		var db *leveldb.DB
		db, err = leveldb.OpenFile(path.Join(dbPath, r.DBName()), nil)
		if err != nil{
			return err
		}
		defer db.Close()
		var value []byte
		value, err = db.Get(r.GetKey(),nil)
		if err != nil{
			return err
		}
		err = r.Decode(value)
		if err != nil{
			return err
		}
	default:
		return errors.New("unknown rec type, must be pointer")
	}
	return
}
