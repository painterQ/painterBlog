package models

import "github.com/syndtr/goleveldb/leveldb"

const (
	None = iota
	Top
)

type DocumentLevelDB struct {
	dbPath string
}

func (* DocumentLevelDB)Init(path string) (err error){
 leveldb.OpenFile()
}

func (* DocumentLevelDB)Close(){

}

func (* DocumentLevelDB)GetDocument(key []byte) (content []byte, err error){

}

func (* DocumentLevelDB)GetDocumentByTag(tag ...string) [][]byte{

}
//相同的key会覆盖
func (* DocumentLevelDB)Push(key ,content []byte, isDraft bool) error{

}
