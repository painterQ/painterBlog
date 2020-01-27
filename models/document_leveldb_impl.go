package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"gopkg.in/fatih/set.v0"
	"os"
	"strings"
	"sync/atomic"
)

const (
	None = iota
	Top
)

const (
	mateMaxNum = 20
)

var (
	docPreFix  = []byte("internal__doc::")  //0
	matePreFix = []byte("internal__mate::") //1
	tagPreFix  = []byte("internal__tag::")  //2
)

type DocumentLevelDB struct {
	dbPath   string
	db       *leveldb.DB
	docTotal int
	inited   int32
}

func (ddb *DocumentLevelDB) Init(dbPath string, hookAfterInitDB func(dbPath string)) (err error) {
	if !atomic.CompareAndSwapInt32(&ddb.inited, 0, 1) {
		return errors.New("already init")
	}
	flagOld := false
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
			flagOld = true
			logs.Warn("database (%v) is exist, use old database", path)
			fallthrough
		case err == nil:
			return db
		default:
			logs.Error("opening database (%v) encountered unknown error: %v", path, err.Error())
			return nil
		}
	}

	documentDB := openDB(dbPath)
	if documentDB == nil {
		return errors.New("init error")
	}

	ddb.dbPath = dbPath
	ddb.db = documentDB

	if !flagOld {
		//插入分界的
		_ = ddb.db.Put(matePreFix, matePreFix, nil)
		_ = ddb.db.Put(tagPreFix, tagPreFix, nil)
	}

	//初始化docTotal
	ddb.getSize()

	if hookAfterInitDB != nil {
		hookAfterInitDB(dbPath)
	}
	return nil

}

func (ddb *DocumentLevelDB) getSize() {
	iter := ddb.db.NewIterator(&util.Range{
		Start: matePreFix,
		Limit: tagPreFix,
	}, nil)
	defer iter.Release()
	ddb.docTotal = -1
	for iter.Next() {
		ddb.docTotal++
	}
}

func (ddb *DocumentLevelDB) Close() {
	_ = ddb.db.Close()
}

func (ddb *DocumentLevelDB) GetMate(key []byte, length int) ([]byte, error) {
	//todo get all
	key = nil

	// if length > mateMaxNum {
	// 	return nil, errors.New("params error")
	// }
	iter := ddb.db.NewIterator(&util.Range{
		Start: append(matePreFix, key...),
		Limit: tagPreFix,
	}, nil)
	defer iter.Release()
	buf := make([][]byte, length)

	var i int
	for i = 0; i < length && iter.Next(); i++ {
		tmp := iter.Value()
		buf[i] = make([]byte, len(tmp))
		copy(buf[i], tmp)
	}

	err := iter.Error()
	if err != nil {
		logs.Error("GetMate error %v", err.Error())
		return nil, err
	}
	if i == 0 {
		return []byte("{list:[],total:0}"), nil
	}
	tmp := bytes.Join(buf[1:i], []byte{','})
	return bytes.Join([][]byte{[]byte(`{"list":[`),
		tmp,
		[]byte(`],`),
		[]byte(fmt.Sprintf(`"total":%d}`, ddb.docTotal))},
		nil), nil
}

func (ddb *DocumentLevelDB) GetDocument(key []byte) (content []byte, err error) {
	content, err = ddb.db.Get(addDocPrefix(key), nil)
	switch err {
	case leveldb.ErrNotFound:
		return nil, os.ErrNotExist
	case nil:
		fallthrough
	default:
		return
	}
}

func (ddb *DocumentLevelDB) GetDocumentByTag(tag []string) []string {
	if len(tag) > 4 || len(tag) < 1 {
		return nil
	}

	setAll := set.New(set.NonThreadSafe)
	v1, err1 := ddb.db.Get(addTagPrefix([]byte(tag[0])), nil)
	if err1 != nil {
		return nil
	}
	s1 := bytes.Split(v1, []byte{'|'})
	for i := range s1 {
		setAll.Add(string(s1[i]))
	}

	for i := 1; i < len(tag); i++ {
		v, err := ddb.db.Get(addTagPrefix([]byte(tag[i])), nil)
		if err != nil {
			continue
		}
		s := bytes.Split(v, []byte{'|'})
		tmp := set.New(set.NonThreadSafe)
		for i := range s {
			tmp.Add(string(s[i]))
		}
		setAll = set.Intersection(setAll, tmp)
	}
	list := setAll.List()
	ret := make([]string, setAll.Size())
	for i := range list {
		ret[i] = list[i].(string)
	}
	return ret
}

//相同的key会覆盖
//自动维护  calibration
//使用事务
func (ddb *DocumentLevelDB) Push(content []byte, mate *DocumentMate) (err error) {
	if strings.Contains(mate.ID, "|") {
		return errors.New("paras error")
	}
	key := []byte(mate.ID)

	mateByte, _ := mate.Encode()
	tx, err := ddb.db.OpenTransaction()
	defer func() {
		if err != nil {
			tx.Discard()
		} else {
			err = tx.Commit()
			ddb.getSize()
		}
	}()

	err = tx.Put(addDocPrefix(key), content, nil)
	if err != nil{
		return
	}
	err = tx.Put(addMatePrefix(key), mateByte, nil)
	if err != nil{
		return
	}

	for i := range mate.Tags {
		tagKey := addTagPrefix([]byte(mate.Tags[i]))
		ret, _ := tx.Has(tagKey, nil)
		var v []byte
		if ret { //tag 已经存在
			v, _ = tx.Get(tagKey, nil)
			if !bytes.Contains(v, key) {
				v = bytes.Join([][]byte{v, {'|'}, key}, nil)
			}
			err = tx.Put(tagKey, v, nil)
			if err != nil {
				return
			}
		} else {
			err = tx.Put(tagKey, key, nil)
			if err != nil {
				return
			}
		}

	}
	return
}

func (ddb *DocumentLevelDB)AddTag(t []string)  error{
	batch := new(leveldb.Batch)
	for i:= range t{
		tag := []byte(t[i])
		if len(tag) == 0{
			continue
		}
		batch.Put(addTagPrefix(tag), tag)
	}
	return ddb.db.Write(batch, nil)
}

func (ddb *DocumentLevelDB) GetTag()[]string{
	ret := make([]string, 0, 10)
	iter := ddb.db.NewIterator(&util.Range{
		Start: tagPreFix,
		Limit: nil,
	},nil)
	defer iter.Release()
	i := 0
	iter.Next()
	for iter.Next(){
		tmpSplit := bytes.Split(iter.Key(),[]byte("::"))
		if len(tmpSplit) < 2{
			continue
		}
		ret = append(ret,string(tmpSplit[1]))
		i++
		if i== len(ret){
			tmp := make([]string, len(ret), len(ret) * 2)
			copy(tmp, ret)
			ret = tmp
		}
	}
	return ret
}

func addMatePrefix(key []byte) []byte {
	return bytes.Join([][]byte{matePreFix, key}, nil)
}

func addDocPrefix(key []byte) []byte {
	return bytes.Join([][]byte{docPreFix, key}, nil)
}

func addTagPrefix(key []byte) []byte {
	return bytes.Join([][]byte{tagPreFix, key}, nil)
}
