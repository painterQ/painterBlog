package models

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"net"
	"os"
	"path"
	"sync"
	"time"
)

func init() {
	gob.Register(time.Time{})
	gob.Register(net.IP{})
	gob.Register(Twitter{})
}

func init() {
	/**********  CDN   ***********/
	CDNSingleCase = &CDN{
		UseCDN: beego.AppConfig.DefaultBool("useCDN", false),
	}
	if CDNSingleCase.UseCDN {
		CDNSingleCase.Bucket = beego.AppConfig.String("bucket")
		CDNSingleCase.Domain = beego.AppConfig.String("domain")
		CDNSingleCase.AccessKey = beego.AppConfig.String("accesskey")
		CDNSingleCase.SecretKey = beego.AppConfig.String("secretkey")
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dbPath = path.Join(pwd, beego.AppConfig.String("dbPath"))
	f, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(dbPath, 0777)
		if err != nil {
			panic(err)
		}
	}
	if !f.IsDir() {
		panic(dbPath + " is not dir!")
	}
	if f.Mode()&0700 != 0700 {
		panic(dbPath + " permission deny, please 'chmod'")
	}
}

func init() {
	//todo 已经登录过，数据库有数据，从数据库加载数据
	newManager := new(Manager)
	if err := Query(newManager); err != nil {
		logs.Info("实例化ManagerSingleCase，数据来自配置文件, 打开数据库错误：" + err.Error())
		ManagerSingleCase = &Manager{
			username: beego.AppConfig.String("username"),
			email:    beego.AppConfig.String("email"),
			phoneN:   beego.AppConfig.String("phonenumber"),
			address:  beego.AppConfig.String("address"),
			password: beego.AppConfig.String("password"),
			lastTime: time.Time{},
			loginIP:  net.IPv4(127, 0, 0, 1),
			token:    "",
			twitter: &Twitter{
				Card:    beego.AppConfig.String("twitter.card"),
				Site:    beego.AppConfig.String("twitter.site"),
				Image:   beego.AppConfig.String("twitter.image"),
				Address: beego.AppConfig.String("twitter.address"),
			},
		}
	} else {
		logs.Info("实例化ManagerSingleCase，数据来自数据库")
		ManagerSingleCase = newManager
	}
	ManagerSingleCase.lock = new(sync.Mutex)

	newBlog := new(Blog)
	if err := Query(newBlog); err != nil {
		logs.Info("实例化BlogSingleCase，数据来自配置文件")
		logs.Info("打开数据库错误：" + err.Error())
		BlogSingleCase = &Blog{
			blogName:    beego.AppConfig.String("blogname"),
			subTitle:    beego.AppConfig.String("subtitle"),
			beiAn:       beego.AppConfig.String("beian"),
			bTitle:      beego.AppConfig.String("btitle"),
			copyright:   beego.AppConfig.String("copyright"),
			copyYear:    beego.AppConfig.String("copyYear"),
			domain:      beego.AppConfig.String("domain"),
			seriesSay:   "seriesSay",
			archivesSay: "archivesSay",
		}
	} else {
		logs.Info("实例化BlogSingleCase，数据来自数据库")
		BlogSingleCase = newBlog
	}
	BlogSingleCase.lock = new(sync.Mutex)
}

// 数据库及表名
const (
	DB                 = "eiblog"
	COLLECTION_ACCOUNT = "account"
	COLLECTION_ARTICLE = "article"
	COUNTER_SERIE      = "serie"
	COUNTER_ARTICLE    = "article"
	SERIES_MD          = "series_md"
	ARCHIVE_MD         = "archive_md"
	ADD                = "add"
	DELETE             = "delete"
)
