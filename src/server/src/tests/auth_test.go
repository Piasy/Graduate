package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
  "encoding/json"
  "bytes"
  "strings"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"

	_ "routers"
  "controllers/errors"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file,
    ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func Test_AuthController_Need_Auth(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/auth", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_AuthController_Need_Auth", "Code[%d]\n%s",
    w.Code, w.Body.String())

	Convey("Subject: Test Need Auth Endpoint\n", t, func() {
    Convey("Status code should be 200", func() {
      So(w.Code, ShouldEqual, 200)
    })
    Convey("The result should not be empty", func() {
      So(w.Body.Len(), ShouldBeGreaterThan, 0)
    })
    Convey("The result should be an Error", func() {
      var result errors.Error
      json.Unmarshal(w.Body.Bytes(), &result)
      So(result == mocked_auth_err, ShouldBeTrue)
    })
	})
}

func Test_AuthController_Auth_Success(t *testing.T) {
  b, _ := json.Marshal(&mocked_auth_info)
	r, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(b))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_AuthController_Auth_Success", "Code[%d]\n%s",
    w.Code, w.Body.String())

	Convey("Subject: Test Auth Success Endpoint\n", t, func() {
    Convey("Status code should be 200", func() {
      So(w.Code, ShouldEqual, 200)
    })
    Convey("The result should not be empty", func() {
      So(w.Body.Len(), ShouldBeGreaterThan, 0)
    })
    Convey("The result should be an true Success", func() {
      type Success struct {
        Succ bool     `json:"success"`
      }
      var result Success
      json.Unmarshal(w.Body.Bytes(), &result)
      So(result == Success {true}, ShouldBeTrue)
    })
    Convey("The response should set session", func() {
      So(strings.Contains(w.HeaderMap.Get(
        http.CanonicalHeaderKey("Set-Cookie")), "beegosessionID"), ShouldBeTrue)
    })
	})
}
