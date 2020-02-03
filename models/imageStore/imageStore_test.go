package imageStore

import (
	"bytes"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"path"
	"sync"
	"testing"
	"time"
)

var instance *levelDBImpl

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	instance = New("./db", "image", "./testdata/tmp").(*levelDBImpl)
	m.Run()
	instance.Release()
	_ = os.RemoveAll("./testdata/tmp")
}

func assertErr(err error) {
	if err != nil {
		panic(err)
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func getRandomImage() []byte {
	f, err := os.Open("./testdata/test.jpeg")
	assertErr(err)
	img, _, err := image.Decode(f)
	assertErr(err)
	all := img.Bounds()
	maxx, maxy := all.Max.X, all.Max.Y
	x2 := max(rand.Intn(maxx), 1)
	y2 := max(rand.Intn(maxy), 1)
	x1 := rand.Intn(x2)
	y1 := rand.Intn(y2)
	rect := image.Rect(x1, y1, x2, y2)

	rgbImg := img.(*image.YCbCr)
	subImage := rgbImg.SubImage(rect)

	buff := bytes.NewBuffer(nil)
	err = jpeg.Encode(buff, subImage, nil)
	assertErr(err)
	_ = f.Close()
	return buff.Bytes()
}

func TestStore(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				tmp := make([]byte, 5)
				_, _ = rand.Read(tmp)
				imginfo, err := instance.SaveImage(getRandomImage(), hex.EncodeToString(tmp), "jpeg", nil)
				assert.NotNil(t, imginfo)
				p := path.Join("./testdata/tmp", imginfo.Src)
				_, err = os.Stat(p)
				assert.Nil(t, err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	index := instance.getIndexAtomic().Uint64()
	assert.Equal(t, uint64(1000), index)

	v, err := instance.db.Get([]byte{0}, nil)
	assert.Nil(t, err)
	r := new(big.Int).SetBytes(v).Uint64()
	assert.Equal(t, uint64(1000), r)

	for i := int64(1); i < 1000; i++ {
		info := instance.GetImageByID(big.NewInt(i).Bytes())
		p := path.Join("./testdata/tmp", info.Src)
		_, err := os.Stat(p)
		assert.Nil(t, err)
	}
}

func TestAddIndexAtomic(t *testing.T) {
	start := instance.getIndexAtomic().Uint64()
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				instance.addIndexAtomic()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	index := instance.getIndexAtomic().Uint64()
	assert.Equal(t, uint64(1000000), index-start)
}

func TestLevelDBImpl_GetAllImageInfo(t *testing.T) {
	for j := 0; j < 100; j++ {
		tmp := make([]byte, 5)
		_, _ = rand.Read(tmp)
		imginfo, err := instance.SaveImage(getRandomImage(), hex.EncodeToString(tmp), "jpeg", nil)
		assert.NotNil(t, imginfo)
		assert.Nil(t, err)
		p := path.Join("./testdata/tmp", imginfo.Src)
		_, err = os.Stat(p)
		assert.Nil(t, err)
	}

	infos := instance.GetAllImageInfo(3, 67)
	assert.Equal(t, 67-3, len(infos))
	for i := 3; i < 67; i++ {
		p := path.Join("./testdata/tmp", infos[i-3].Src)
		_, err := os.Stat(p)
		assert.Nil(t, err)
	}
}

func TestLevelDBImpl_RemoveImage(t *testing.T) {
	start := instance.getIndexAtomic().Uint64() + 1
	for j := 0; j < 100; j++ {
		tmp := make([]byte, 5)
		_, _ = rand.Read(tmp)
		imginfo, err := instance.SaveImage(getRandomImage(), hex.EncodeToString(tmp), "jpeg", nil)
		assert.NotNil(t, imginfo)
		p := path.Join("./testdata/tmp", imginfo.Src)
		_, err = os.Stat(p)
		assert.Nil(t, err)
	}
	end := instance.getIndexAtomic().Uint64()
	infos := instance.GetAllImageInfo(int64(end+1), int64(start)) // include end
	assert.Equal(t, int(end-start+1), len(infos))
	err := instance.RemoveImage(infos[33].ID)
	assert.Nil(t, err)
	p := path.Join("./testdata/tmp", infos[33].Src)
	_, err = os.Stat(p)
	assert.True(t, os.IsNotExist(err))
	infos = instance.GetAllImageInfo(int64(start), int64(end+1)) // include end
	assert.Equal(t, int(end-start+1-1), len(infos))
}

func TestNew(t *testing.T) {
	instance.Release()
	instance = New("./db", "image", "./testdata/tmp").(*levelDBImpl)
	assert.NotNil(t, instance)
	instance.Release()
	list, err := ioutil.ReadDir("./testdata/tmp/db")
	assert.Nil(t, err)
	for i := range list {
		f, err := os.OpenFile(path.Join("./testdata/tmp/db", list[i].Name()), os.O_RDWR|os.O_SYNC, 0666)
		assert.Nil(t, err)
		err = f.Truncate(list[i].Size() / 2)
		assert.Nil(t, err)
		_, err = f.Write([]byte("some date"))
		assert.Nil(t, err)
		assert.Nil(t, f.Sync())
		assert.Nil(t, f.Close())
	}
	assert.Nil(t, New("./db", "image", "./testdata/tmp"))
	_ = os.RemoveAll("./testdata/tmp")
	instance = New("./db", "image", "./testdata/tmp").(*levelDBImpl)
}


func TestIndex(t *testing.T) {
	start := instance.getIndexAtomic().Uint64()
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				instance.addIndexAtomic()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	index := instance.getIndexAtomic().Uint64()
	assert.Equal(t, uint64(1000000), index-start)
	instance.Release()
	instance = New("./db", "image", "./testdata/tmp").(*levelDBImpl)
	assert.Equal(t,big.NewInt(int64(index)),instance.index)
	assert.Equal(t,index,instance.getIndexAtomic().Uint64())
}