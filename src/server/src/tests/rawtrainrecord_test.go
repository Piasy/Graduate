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
	"github.com/astaxie/beego/logs"
	. "github.com/smartystreets/goconvey/convey"

	_ "routers"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file,
		".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	beego.SetLevel(logs.LevelInfo)
}

func Test_RawTrainRecordController_Need_Auth(t *testing.T) {
	r, _ := http.NewRequest("POST", "/api/rawtrainrecord", bytes.NewBuffer([]byte(mocked_train_post_data_str)))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_RawTrainRecordController_Need_Auth", "Code[%d]\n%s", w.Code,
		w.Body.String())

	Convey("Subject: Test RawTrainRecord Endpoint\n", t, func() {
    Convey("Status code should be 302", func() {
      So(w.Code, ShouldEqual, 302)
    })
    Convey("The redirect location should be \"/api/auth?req=L2FwaS9yYXd0cmFpbnJlY29yZA==\"",
					func() {
      So(w.HeaderMap.Get("Location"), ShouldEqual, "/api/auth?req=L2FwaS9yYXd0cmFpbnJlY29yZA==")
    })
	})
}

func Test_RawTrainRecordController_POST(t *testing.T) {
	cookie := auth(t)
	appendRawData(cookie, t)
	appendRawData(cookie, t)
	appendRawData(cookie, t)
	flushRawData(cookie, t)
}

func auth(t interface {}) string {
	b, _ := json.Marshal(&mocked_auth_info)
	r, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(b))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_PlayersController_Insert_Query_Update(Auth)", "Code[%d]\n%s", w.Code,
		w.Body.String())

	Convey("Subject: Test RawTrainRecord Endpoint\n", t, func() {
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
	return w.HeaderMap.Get(http.CanonicalHeaderKey("Set-Cookie"))
}

func appendRawData(cookie string, t interface {}) {
	r, _ := http.NewRequest("POST", "/api/rawtrainrecord", bytes.NewBuffer([]byte(mocked_train_post_data_str)))
	r.Header.Set("Cookie", cookie)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_RawTrainRecordController_Need_Auth", "Code[%d]\n%s", w.Code,
		w.Body.String())

	Convey("Subject: Test RawTrainRecord Endpoint\n", t, func() {
		Convey("Status code should be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The redirect location should be \"/api/auth?req=L2FwaS9yYXd0cmFpbnJlY29yZA==\"",
					func() {
			type Success struct {
				Succ bool `json:"success"`
			}
			var result Success
			json.Unmarshal(w.Body.Bytes(), &result)
			So(result.Succ, ShouldBeTrue)
		})
	})
}

func flushRawData(cookie string, t interface {}) {
	r, _ := http.NewRequest("POST", "/api/rawtrainrecord", bytes.NewBuffer([]byte("{\"op\":\"flush\"}")))
	r.Header.Set("Cookie", cookie)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_RawTrainRecordController_Need_Auth", "Code[%d]\n%s", w.Code,
		w.Body.String())

	Convey("Subject: Test RawTrainRecord Endpoint\n", t, func() {
		Convey("Status code should be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The redirect location should be \"/api/auth?req=L2FwaS9yYXd0cmFpbnJlY29yZA==\"",
					func() {
			type Success struct {
				Succ bool `json:"success"`
			}
			var result Success
			json.Unmarshal(w.Body.Bytes(), &result)
			So(result.Succ, ShouldBeTrue)
		})
	})
}

func Benchmark_Raw_Data_Post(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		cookie := auth(b)
		i := 1
		for pb.Next() {
			if i % (10 * 60 * 45) == 0 {
				flushRawData(cookie, b)
			} else {
				appendRawData(cookie, b)
			}
			i++
    }
	})
}
