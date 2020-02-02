package models

import (
	"github.com/astaxie/beego"
	"github.com/painterQ/painterBlog/models/imageStore"
	"path"
)

var DocumentDataBaseSingleCase DocumentDataBase
var AuthorSingleCase Author
var ImageStoreSingleCase imageStore.ImageStore

const (
	DocDBPath     = "./doc"
	ImageBDPath = "./imageDB"
	ImagePath = "./image"

	DefaultDBPathConfig = "/tmp/painter"
	ConfigDBPath        = "db"
)

//GlobalDBPath absolute path for database
var GlobalDBPath string

func init() {
	GlobalDBPath = beego.AppConfig.DefaultString(ConfigDBPath, DefaultDBPathConfig)

	DocumentDataBaseSingleCase = &DocumentLevelDB{}
	err := DocumentDataBaseSingleCase.Init(path.Join(GlobalDBPath, DocDBPath),nil)
	if err != nil{
		panic(err)
	}

	err = AuthorSingleCase.Start("./conf/app.json")
	if err != nil{
		panic(err)
	}

	ImageStoreSingleCase = imageStore.New(ImageBDPath,ImagePath,GlobalDBPath)
}
