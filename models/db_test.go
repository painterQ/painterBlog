package models

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
	"time"
)

var impl = new(DocumentLevelDB)
var path string
var global int

func TestMain(m *testing.M) {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	path = "/tmp/" + hex.EncodeToString(b)
	err := impl.Init(path, func(dbPath string) {
		global = 1
	})
	if err != nil {
		panic(err)
	}
	m.Run()
	impl.Close()
	_ = os.RemoveAll(path)
}

func print() {
	iter := impl.db.NewIterator(nil, nil)
	for iter.Next() {
		k, v := iter.Key(), iter.Value()
		fmt.Printf("k:%v   |   v:%v\n", string(k), string(v))
	}
}

func TestInit(t *testing.T) {
	err := impl.Init(path, nil)
	assert.NotNil(t, err)
	assert.NotNil(t, impl.db)
	assert.Equal(t, path, impl.dbPath)
	assert.Equal(t, int32(1), impl.inited)
	assert.Equal(t, 1, global)
}

func TestPush(t *testing.T) {
	mate := &DocumentMate{
		ID:         "/doc/doc/1",
		Title:      "golang slice",
		SubTitle:   "blog",
		Tags:       []string{"blog","doc"},
		Attr:       1,
		LastTime:   time.Now(),
		Abstract:  "<p>first abstract, a doc for golang slice</p>",
	}
	err := impl.Push([]byte("<content>/doc/doc/1</content>"),mate)
	assert.Nil(t, err)
	mate = &DocumentMate{
		ID:         "/doc/doc/2",
		Title:      "golang map",
		SubTitle:   "blog",
		Tags:       []string{"blog","document"},
		Attr:       0,
		LastTime:   time.Now(),
		Abstract:   "<p>second abstract, a document for golang map</p>",
	}
	err = impl.Push([]byte("<content>/doc/doc/2</content>"),mate)
	assert.Nil(t, err)
	diff := impl.GetDocumentByTag([]string{"blog"})
	assert.Equal(t,2, len(diff))
	diff = impl.GetDocumentByTag([]string{"doc"})
	assert.Equal(t,1, len(diff))
	assert.Equal(t,"/doc/doc/1", diff[0])
	diff = impl.GetDocumentByTag([]string{"document"})
	assert.Equal(t,1, len(diff))
	assert.Equal(t,"/doc/doc/2", diff[0])

	diff = impl.GetDocumentByTag([]string{"document", "blog"})
	assert.Equal(t,1, len(diff))
	assert.Equal(t,"/doc/doc/2", diff[0])

	diff = impl.GetDocumentByTag([]string{"document", "doc"})
	assert.Equal(t,0, len(diff))

	err = impl.Push([]byte("<content>/doc/doc/3</content>"),&DocumentMate{
		ID:"/doc/doc|/3",
	})
	assert.NotNil(t, err)
}

func TestGetDocument(t *testing.T) {
	mate := &DocumentMate{
		ID:         "/doc/doc/1",
		Title:      "golang slice",
		SubTitle:   "blog",
		Tags:       []string{"blog","doc"},
		Attr:       1,
		LastTime:   time.Now(),
		Abstract:   "<p>first abstract, a doc for golang slice</p>",
	}
	err := impl.Push([]byte("<content>/doc/doc/1</content>"),mate)
	assert.Nil(t, err)
	c,err := impl.GetDocument([]byte("/doc/doc/1"))
	assert.Nil(t, err)
	assert.Equal(t, "<content>/doc/doc/1</content>",string(c))
	print()
	js,err :=impl.GetMate([]byte("/doc/doc/1"),2)
	assert.Nil(t,err)
	target := make([]*DocumentMate, 1)
	err = json.Unmarshal(js, &target)
	assert.Nil(t,err)
	assert.Equal(t, target[0].ID, "/doc/doc/1")

	mate = &DocumentMate{
		ID:         "/doc/doc/2", //todo 检查参数是否一致
		Title:      "golang map",
		SubTitle:   "blog",
		Tags:       []string{"blog","document"},
		Attr:       0,
		LastTime:   time.Now(),
		Abstract:   "<p>second abstract, a document for golang map</p>",
	}
	err = impl.Push([]byte("<content>/doc/doc/2</content>"),mate)
	target = make([]*DocumentMate, 2)
	js,err =impl.GetMate([]byte("/doc/doc/1"),2)
	assert.Nil(t,err)
	err = json.Unmarshal(js, &target)
	assert.Nil(t,err)
	assert.Equal(t, target[1].Tags[0], "blog")
}

func TestOpenAgain(t *testing.T) {
	impl.Close()
	impl2 := new(DocumentLevelDB)
	err := impl2.Init(path,nil)
	assert.Nil(t, err)
	impl2.Close()
}

func TestGetTag(t *testing.T) {
	for i:= 0;i< 10;i++{
		tmp := make([]byte,20)
		_,_ = rand.Read(tmp)
		err := impl.AddTag([]string{hex.EncodeToString(tmp)})
		assert.Nil(t, err)
	}
	g := impl.GetTag()
	assert.Equal(t, 10, len(g))
	for i:= 0;i <9;i++{
		assert.True(t, strings.Compare( g[i],g[i+1]) < 0)
	}
}