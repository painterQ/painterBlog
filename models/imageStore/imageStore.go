package imageStore

import (
	"bytes"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math/big"
	"os"
	"path"
	"sync"
)

var one = big.NewInt(1)
var zero = big.NewInt(0)

type ImageInfo struct {
	ID   []byte	`json:"id"`
	Name string	`json:"name"`
	Type string	`json:"type"`
	Src  string	`json:"src"`
	//todo 缩略图
}

type ImageStore interface {
	Release()
	GetImageByID(id []byte) *ImageInfo
	GetImageByName(name []byte) []byte
	GetAllImageInfo(s, e int64) []*ImageInfo
	//SaveImage, if id is nil, use a random id
	SaveImage(data []byte, name, imageType string, id []byte) (*ImageInfo, error)
	RemoveImage(id []byte) error
}

type levelDBImpl struct {
	db           *leveldb.DB
	imagePath    string
	index        *big.Int //point to top
	indexAtom    sync.RWMutex
	globalDBPath string
}

func openDB(path string, ErrorIfExist bool, flagOld *bool) *leveldb.DB {
	if flagOld != nil {
		*flagOld = false
	}
	db, err := leveldb.OpenFile(path, &opt.Options{
		Comparer:     new(ImageCmp),
		ErrorIfExist: ErrorIfExist,
	})
	switch {
	case errors.IsCorrupted(err):
		logs.Error("database (%v) corrupted, process try to recover: %v", path, err.Error())
		db, err = leveldb.RecoverFile(path, nil)
		if err != nil {
			logs.Error("try to recover fail: %v", path, err.Error())
			return nil
		}
		if flagOld != nil {
			*flagOld = true
		}
		return openDB(path, false, nil)
	case err == os.ErrExist:
		logs.Warn("database (%v) is exist, use old database", path)
		if flagOld != nil {
			*flagOld = true
		}
		return openDB(path, false, nil)
	case err == nil:
		return db
	default:
		logs.Error("opening database (%v) encountered unknown error: %v", path, err.Error())
		return nil
	}
}

//New generate a ImageStore
//mateDBPath & imagePath: relative path, relative to the path globalDBPath
func New(mateDBPath, imagePath, globalDBPath string) ImageStore {
	flagOld := false
	db := openDB(path.Join(globalDBPath, mateDBPath), true, &flagOld)
	if db == nil{
		return nil
	}
	ret := &levelDBImpl{
		db:           db,
		imagePath:    imagePath,
		globalDBPath: globalDBPath}
	if !flagOld {
		ret.index = big.NewInt(0)
		ierr := ret.db.Put([]byte{0}, ret.index.Bytes(), nil)
		if ierr != nil {
			panic("init ImageStore error:" + ierr.Error())
		}
	} else {
		ierr := ret.UpdateIndexFormDb()
		fmt.Println("### init index with "+ ret.index.Text(10))
		if ierr != nil {
			panic("init ImageStore error:" + ierr.Error())
		}
	}

	imageDirPath := path.Join(globalDBPath, imagePath)
	fi, err := os.Stat(imageDirPath)
	switch {
	case err == nil:
		return ret
	case os.IsNotExist(err):
		merr := os.Mkdir(imageDirPath, 0777)
		if merr != nil {
			panic("create dir:" + imageDirPath + " error:" + merr.Error())
		}
	case !fi.IsDir():
		panic(imageDirPath + " is not a dir")
	default:
		panic("open " + imageDirPath + " error:" + err.Error())
	}
	return ret
}

func (i *levelDBImpl) Release() {
	_ = i.db.Close()
}

func (i *levelDBImpl) GetImageByID(id []byte) *ImageInfo {
	v, err := i.db.Get(id, nil)
	if err != nil {
		return nil
	}
	tmp := new(ImageInfo)
	_, err = asn1.Unmarshal(v, tmp)
	if err != nil {
		return nil
	}
	return tmp
}

func (i *levelDBImpl) GetImageByName(name []byte) []byte {
	panic("implement me")
}

func (i *levelDBImpl) GetAllImageInfo(s, e int64) []*ImageInfo {
	if e < s {
		s, e = e, s
	}
	ret := make([]*ImageInfo, 0, e-s)
	iter := i.db.NewIterator(&util.Range{
		Start: big.NewInt(s).Bytes(),
		Limit: big.NewInt(e).Bytes(),
	}, nil)
	defer iter.Release()
	for iter.Next() {
		v := iter.Value()
		tmp := new(ImageInfo)
		_, ierr := asn1.Unmarshal(v, tmp)
		if ierr != nil {
			continue
		}
		ret = append(ret, tmp)
	}
	return ret
}

func (i *levelDBImpl) addIndexAtomic() *big.Int {
	//写穿，先写数据库后写内存
	i.indexAtom.Lock()
	defer i.indexAtom.Unlock()
	n := new(big.Int).Add(i.index, one)
	err := i.db.Put([]byte{0}, n.Bytes(), nil)
	if err != nil {
		panic(err)
	}

	i.index = n
	return i.index
}

func (i *levelDBImpl) getIndexAtomic() *big.Int {
	//直接获取内存的值
	i.indexAtom.RLock()
	defer i.indexAtom.RUnlock()
	z := big.NewInt(0)
	return z.Add(i.index, z)
}

func (i *levelDBImpl) UpdateIndexFormDb() error {
	i.indexAtom.Lock()
	defer i.indexAtom.Unlock()
	v, err := i.db.Get([]byte{0}, nil)
	if err != nil {
		return err
	}
	i.index = new(big.Int).SetBytes(v)
	return nil
}

func (i *levelDBImpl) SaveImage(data []byte, name, imageType string, id []byte) (*ImageInfo, error) {
	//check image
	_, n, e := image.Decode(bytes.NewBuffer(data))
	if n != imageType || e != nil { //png jpeg gif
		return nil, errors.New("image type error")
	}

	//write db
	if id == nil {
		id = i.addIndexAtomic().Bytes()
		fmt.Println("现在image的index是",new(big.Int).SetBytes(id).Text(10))
	}
	src := path.Join(i.imagePath, hex.EncodeToString(id)+"."+imageType)
	ret := &ImageInfo{
		ID:   id,
		Name: name,
		Type: imageType,
		Src:  src,
	}
	value, _ := asn1.Marshal(*ret)
	err := i.db.Put(id, value, nil)
	fmt.Println("db put with id "+ new(big.Int).SetBytes(id).Text(10))
	if err != nil {
		return nil, errors.New("db put error:" + err.Error())
	}

	//write file
	absPath := path.Join(i.globalDBPath, src)
	file, err := os.OpenFile(absPath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		_ = i.db.Delete(ret.ID, nil)
		return nil, errors.New("write to file error:" + err.Error())
	}
	defer func() {
		_ = file.Close()
	}()

	_, _ = file.Seek(0, 0)
	sizeOfFile, err := file.Write(data)
	if err != nil {
		_ = i.db.Delete(ret.ID, nil)
		return nil, errors.New("write to file error:" + err.Error())
	}
	err = file.Truncate(int64(sizeOfFile))
	if err != nil {
		_ = i.db.Delete(ret.ID, nil)
		return nil, errors.New("write to file error:" + err.Error())
	}

	return ret, nil
}

func (i *levelDBImpl) RemoveImage(id []byte) error {
	info := i.GetImageByID(id)
	if info == nil {
		return nil
	}
	err := i.db.Delete(id, nil)
	if err != nil {
		return err
	}
	return os.Remove(path.Join(i.globalDBPath, info.Src))
}
