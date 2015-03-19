package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	_ "routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file,
		".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestMain is a sample to run an endpoint test
func Test_PlayersController_Need_Auth(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/player?page=0&num=10", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_PlayersController_Need_Auth", "Code[%d]\n%s", w.Code,
		w.Body.String())

	Convey("Subject: Test Player Endpoint\n", t, func() {
    Convey("Status code should be 302", func() {
      So(w.Code, ShouldEqual, 302)
    })
    Convey("The redirect location should be \"/api/auth?req=L2FwaS9wbGF5ZXI_cGFnZT0wJm51bT0xMA==\"",
					func() {
      So(w.HeaderMap.Get("Location"), ShouldEqual, "/api/auth?req=L2FwaS9wbGF5ZXI_cGFnZT0wJm51bT0xMA==")
    })
	})
}
