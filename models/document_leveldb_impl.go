package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"os"
)

const (
	None = iota
	Top
)

type DocumentLevelDB struct {
	dbPath string
	db *leveldb.DB
	tagMapDB *leveldb.DB
}

func (ddb * DocumentLevelDB)Init(path string ,hookAfterInitDB func()error) (err error){
	create := func(newDB *leveldb.DB) error{
		ddb.dbPath = path
		ddb.db = newDB
		if hookAfterInitDB != nil{
			return hookAfterInitDB()
		}
		return nil
	}

 	db, err := leveldb.OpenFile(path,nil)
	switch  {
	case errors.IsCorrupted(err):
		logs.Error("database (%v) corrupted, process try to recover: %v",path, err.Error())
		db,err = leveldb.RecoverFile(path, nil)
		if err != nil{
			logs.Error("try to recover fail: %v",path, err.Error())
			return err
		}
		return create(db)
	case err == os.ErrExist :
		logs.Warn("database (%v) is exist, use old database",path)
		fallthrough
	case err == nil:
		return create(db)
	default:
		logs.Error("opening database (%v) encountered unknown error: %v",path, err.Error())
		return err
	}
}

func (ddb * DocumentLevelDB)Close(){
	_ = ddb.db.Close()
}

func (ddb * DocumentLevelDB)GetDocument(key []byte) (content []byte, err error){
	content,err = ddb.db.Get(key, nil)
	switch err {
	case leveldb.ErrNotFound:
		return nil, os.ErrNotExist
	case nil:
		fallthrough
	default:
		return
	}
}



func (ddb * DocumentLevelDB)GetDocumentByTag(tag ...string) [][]byte{
	return nil
}
//相同的key会覆盖
func (ddb * DocumentLevelDB)Push(key ,content []byte, isDraft bool) (err error){
	return ddb.db.Put(key, content, nil)
}
