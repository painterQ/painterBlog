package models

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego/logs"
	"sync"
)

var BlogSingleCase *Blog

type BlogUpdater map[string]interface{}

type Blog struct {
	// 博客名 blogName subTitle beiAn bTitle copyright
	blogName string
	// SubTitle
	subTitle string
	// 备案号
	beiAn string
	// 底部title
	bTitle string
	// 版权声明
	copyright string
	// 版权到期
	copyYear string
	// 域名
	domain string
	// 专题，倒序
	seriesSay string
	//Series    SortSeries
	// 归档描述
	archivesSay string
	//Archives    SortArchives
	// 忽略存储，前端界面全部缓存
	//PageSeries   string                  `bson:"-"` // 专题页面
	//PageArchives string                  `bson:"-"` // 归档页面
	//Tags         map[string]SortArticles `bson:"-"` // 标签 name->tag
	//Articles     SortArticles            `bson:"-"` // 所有文章
	//MapArticles  map[string]*Article     `bson:"-"` // url->Article
	//CH           chan string             `bson:"-"` // channel
	lock sync.Locker
}

func (m *Blog) Domain() string {
	return m.domain
}

func (m *Blog) SetDomain(domain string) {
	m.domain = domain
}

func (m *Blog) CopyYear() string {
	return m.copyYear
}

func (m *Blog) SetCopyYear(copyYear string) {
	m.copyYear = copyYear
}

func (m *Blog) ArchivesSay() string {
	return m.archivesSay
}

func (m *Blog) SetArchivesSay(archivesSay string) {
	m.archivesSay = archivesSay
}

func (m *Blog) SeriesSay() string {
	return m.seriesSay
}

func (m *Blog) SetSeriesSay(seriesSay string) {
	m.seriesSay = seriesSay
}

func (m *Blog) Copyright() string {
	return m.copyright
}

func (m *Blog) SetCopyright(copyright string) {
	m.copyright = copyright
}

func (m *Blog) BTitle() string {
	return m.bTitle
}

func (m *Blog) SetBTitle(bTitle string) {
	m.bTitle = bTitle
}

func (m *Blog) BeiAn() string {
	return m.beiAn
}

func (m *Blog) SetBeiAn(beiAn string) {
	m.beiAn = beiAn
}

func (m *Blog) SubTitle() string {
	return m.subTitle
}

func (m *Blog) SetSubTitle(subTitle string) {
	m.subTitle = subTitle
}

func (m *Blog) BlogName() string {
	return m.blogName
}

func (m *Blog) SetBlogName(blogName string) {
	m.blogName = blogName
}

func (m *Blog) NewUpdate() BlogUpdater {
	return make(map[string]interface{}, 9)
}

func (b BlogUpdater) Update(k string, v interface{}) BlogUpdater {
	switch k {
	case "blogName", "subTitle", "beiAn", "bTitle", "copyright", "archivesSay","seriesSay","domain","copyYear":
		t, ok := v.(string)
		if !ok {
			//todo
		}
		b[k] = t
	//case "lastTime":
	//	t, ok := v.(time.Time)
	//	if !ok {
	//		//todo
	//	}
	//	b[k] = t
	//case "loginIP":
	//	t, ok := v.(net.IP)
	//	if !ok {
	//		//todo
	//	}
	//	b[k] = t
	default:
		b["error"] = fmt.Sprintf("BlogUpdater, default 分支:%v", k)
	}
	return b
}

func (b BlogUpdater) EndUpdate() error {
	if err, ok := b["error"]; ok {
		return fmt.Errorf("非法的修改：%v", err)
	}
	BlogSingleCase.lock.Lock()
	defer BlogSingleCase.lock.Unlock()
	for k,v := range b {
		switch k {
		case "blogName":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update blogName err:(%v)is not string",v)
			}
			BlogSingleCase.blogName = t
		case "subTitle":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update subTitle err:(%v)is not string",v)
			}
			BlogSingleCase.subTitle = t
		case "domain":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update domain err:(%v)is not string",v)
			}
			BlogSingleCase.domain = t
		case "copyYear":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update copyYear err:(%v)is not string",v)
			}
			BlogSingleCase.copyYear = t
		case "beiAn":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update beiAn err:(%v)is not string",v)
			}
			BlogSingleCase.beiAn = t
		case "bTitle":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update bTitle err:(%v)is not string",v)
			}
			BlogSingleCase.bTitle = t
		case "copyright":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update copyright err:(%v)is not string",v)
			}
			BlogSingleCase.copyright = t
		case "archivesSay":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update archivesSay err:(%v)is not string",v)
			}
			BlogSingleCase.archivesSay = t
		case "seriesSay":
			t, ok := v.(string)
			if !ok {
				return fmt.Errorf("update seriesSay err:(%v)is not string",v)
			}
			BlogSingleCase.seriesSay = t
		//case "lastTime":
		//	t, ok := v.(time.Time)
		//	if !ok {
		//		return fmt.Errorf("update token err:(%v)is not time.Time",v)
		//	}
		//	BlogSingleCase.lastTime = t
		//case "loginIP":
		//	t, ok := v.(net.IP)
		//	if !ok {
		//		return fmt.Errorf("update loginIP err:(%v)is not time.Time",v)
		//	}
		//	BlogSingleCase.loginIP = t
		default:
		}
	}
	err := Save(BlogSingleCase)
	if err != nil{
		//todo
	}
	return nil
}

func (m *Blog) Encode() (v []byte) {
	value := make(map[string]interface{}, 9)
	buf := bytes.NewBuffer(make([]byte, 0, 512))

	value["blogName"] = m.blogName
	value["subTitle"] = m.subTitle
	value["beiAn"] = m.beiAn
	value["bTitle"] = m.bTitle
	value["copyright"] = m.copyright
	value["domain"] = m.domain
	value["copyYear"] = m.copyYear
	value["archivesSay"] = m.archivesSay
	value["seriesSay"] = m.seriesSay
	err := gob.NewEncoder(buf).Encode(value)
	if err != nil {
		logs.Error("Blog Encode Err:"+err.Error())
	}
	return buf.Bytes()
}

func (m *Blog) Decode(v []byte)(err error){
	defer func() {
		r := recover()
		if r != nil{
			err = fmt.Errorf("panic when decode,%v",r)
		}
	}()
	dc := gob.NewDecoder(bytes.NewBuffer(v))
	value := make(map[string]interface{}, 9)
	err = dc.Decode(&value)
	if err != nil {
		return err
	}
	m.blogName = value["blogName"].(string)
	m.subTitle = value["subTitle"].(string)
	m.beiAn = value["beiAn"].(string)
	m.bTitle = value["bTitle"].(string)
	m.copyright = value["copyright"].(string)
	m.domain = value["domain"].(string)
	m.copyYear = value["copyYear"].(string)
	m.archivesSay = value["archivesSay"].(string)
	m.seriesSay = value["seriesSay"] .(string)
	return nil
}

func (m *Blog)GetKey()(k []byte)  {
	return []byte("Blog")
}

func (m *Blog) DBName() string {
	return "Blog"
}
