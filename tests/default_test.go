package test
//
//import (
//	"fmt"
//	_ "github.com/painterQ/painterBlog/routers"
//	"net/http"
//	"net/http/httptest"
//	"os"
//	"path/filepath"
//	"testing"
//
//	"github.com/astaxie/beego"
//	. "github.com/smartystreets/goconvey/convey"
//)
//
//func init() {
//	file, _ := os.Getwd()
//	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
//	fmt.Println("$$$$",apppath)
//	beego.TestBeegoInit(apppath)
//}
//
//// TestBeego is a sample to run an endpoint test
//func TestBeego(t *testing.T) {
//	r, _ := http.NewRequest("GET", "/", nil)
//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)
//
//	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())
//
//	Convey("Subject: Test Station Endpoint\n", t, func() {
//		Convey("Status Code Should Be 200", func() {
//			So(w.Code, ShouldEqual, 200)
//		})
//		Convey("The Result Should Not Be Empty", func() {
//			So(w.Body.Len(), ShouldBeGreaterThan, 0)
//		})
//	})
//}
