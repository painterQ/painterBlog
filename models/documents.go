package models

//DocumentDataBase document data base
type DocumentDataBase interface {
	Init(path string ,hookAfterInitDB func()error) (err error)
	Close()

	GetDocument(key []byte) (content []byte, err error)
	GetDocumentByTag(tag ...string) [][]byte
	//相同的key会覆盖
	Push(key ,content []byte, isDraft bool) error
}


