package imageStore

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/nfnt/resize"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math/big"
	"os"
	"path"
	"sync"
)

var one = big.NewInt(1)
var zero = big.NewInt(0)

type ImageInfo struct {
	ID     []byte `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Src    string `json:"src"`
	Small  []byte `json:"small,omitempty"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type ImageStore interface {
	Release()
	GetImageByID(id []byte) *ImageInfo
	GetImageByName(name []byte) []byte
	GetAllImageInfo(webAddr string, s, e int64) [][]byte
	//SaveImage, if id is nil, use a random id
	SaveImage(data io.Reader, name string, id []byte) ([]byte, error)
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
	if db == nil {
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
		fmt.Println("### init index with " + ret.index.Text(10))
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
	err = json.Unmarshal(v, tmp)
	if err != nil {
		return nil
	}
	return tmp
}

func (i *levelDBImpl) GetImageByName(name []byte) []byte {
	panic("implement me")
}

func (i *levelDBImpl) GetAllImageInfo(webAddr string, s, e int64) [][]byte {
	if e < s {
		s, e = e, s
	}
	if s < 1{
		return nil
	}
	ret := make([][]byte, 0, e-s)
	iter := i.db.NewIterator(&util.Range{
		Start: big.NewInt(s).Bytes(),
		Limit: big.NewInt(e).Bytes(),
	}, nil)
	defer iter.Release()
	for iter.Next() {
		v := iter.Value()
		buf := make([]byte,len(v))
		copy(buf,v)
		ret = append(ret, buf)
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

func getSmallerChan(width uint, t string, img image.Image) chan *bytes.Buffer {
	ret := make(chan *bytes.Buffer, 1)
	go func() {
		m := resize.Resize(width, 0, img, resize.Lanczos3)
		buf := new(bytes.Buffer)
		switch t {
		case "jpeg":
			_ = jpeg.Encode(buf, m, &jpeg.Options{Quality: 50})
		case "png":
			_ = (&png.Encoder{CompressionLevel: png.BestSpeed}).Encode(buf, m)
		case "gif":
			_ = gif.Encode(buf, m, nil)
		}
		ret <- buf
	}()
	return ret
}

func (i *levelDBImpl) SaveImage(data io.Reader, name string, id []byte) (JSON []byte, err error) {
	//open a temp file
	nonce := make([]byte, 32)
	_, _ = rand.Read(nonce)
	var imgInfo *ImageInfo = nil
	tempName := path.Join(i.globalDBPath, hex.EncodeToString(nonce))
	file, err := os.OpenFile(tempName, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0666)
	if err != nil {
		return nil, errors.New("write to file error:" + err.Error())
	}

	//check image and write to temp file
	img, n, err := image.Decode(io.TeeReader(data, file))
	_ = file.Close()
	if err != nil { //png jpeg gif
		return nil, errors.New("image type error：" + err.Error())
	}

	size := img.Bounds().Size()
	var smallerChan chan *bytes.Buffer
	if (size.X > 200 || size.Y > 200) && n != "gif"{
		smallerChan = getSmallerChan(200, n, img)
	}


	//handle error
	absPath := ""
	defer func() {
		if err == nil && imgInfo != nil && JSON != nil && absPath != "" {
			_ = os.Rename(tempName, absPath)
		} else {
			err = fmt.Errorf("store image error: %v", err)
			_ = os.Remove(tempName)
		}
	}()

	//write db
	if id == nil {
		id = i.addIndexAtomic().Bytes()
		fmt.Println("现在image的index是", new(big.Int).SetBytes(id).Text(10))
	}
	src := path.Join(i.imagePath, hex.EncodeToString(id)+"."+n)
	absPath = path.Join(i.globalDBPath, src)

	imgInfo = &ImageInfo{
		ID:     id,
		Name:   name,
		Type:   n,
		Src:    src,
		Small:  nil,
		Width:  size.X,
		Height: size.Y,
	}

	if (size.X > 200 || size.Y > 200) && n != "gif"{
		imgInfo.Small = (<-smallerChan).Bytes()
	}


	JSON, err = json.Marshal(*imgInfo)
	if err != nil {
		return
	}
	err = i.db.Put(id, JSON, nil)
	fmt.Println("db put with id " + new(big.Int).SetBytes(id).Text(10))
	return
}

func AddWebDN2ImgArray(webDN string, JSON [][]byte) []byte {
	array := bytes.Join(JSON, []byte{','})
	return bytes.Join([][]byte{[]byte(`{"list":[`), array, []byte(`],"webDN":"`), []byte(webDN), []byte(`"}`)}, nil)
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
