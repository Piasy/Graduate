package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	"time"
	"encoding/json"
	"bytes"
	"strings"
	"fmt"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"

	_ "routers"
	"models/types"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file,
		".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

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

func Test_PlayersController_Insert_Query_Update(t *testing.T) {
	//auth
	b, _ := json.Marshal(&mocked_auth_info)
	r, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(b))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_PlayersController_Insert_Query_Update(Auth)", "Code[%d]\n%s", w.Code,
		w.Body.String())

	Convey("Subject: Test Player Insert, Query and Update Endpoint\n", t, func() {
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
	cookie := w.HeaderMap.Get(http.CanonicalHeaderKey("Set-Cookie"))


	//insert
	mocked_player.Name = fmt.Sprintf("%s_%d", mocked_player.Name, time.Now().UnixNano())
	b, _ = json.Marshal(&mocked_player)
	r, _ = http.NewRequest("POST", "/api/player", bytes.NewBuffer(b))
	r.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_PlayersController_Insert_Query_Update(Insert)", "Code[%d]\n%s", w.Code,
		w.Body.String())

	Convey("Subject: Test Player Insert, Query and Update Endpoint\n", t, func() {
    Convey("Status code should be 200", func() {
      So(w.Code, ShouldEqual, 200)
    })
		Convey("Response should be a QueryResult", func() {
			type Success struct {
				Succ bool `json:"success"`
				Id string	`json:"id"`
			}
			var result Success
			json.Unmarshal(w.Body.Bytes(), &result)
      So(result.Succ, ShouldBeTrue)
			mocked_player.ObjId = bson.ObjectIdHex(result.Id)
    })
	})


	//query
	r, _ = http.NewRequest("GET", "/api/player?page=0&num=20", nil)
	r.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_PlayersController_Insert_Query_Update(Query)", "Code[%d]\n%s", w.Code,
		w.Body.String())

	Convey("Subject: Test Player Insert, Query and Update Endpoint\n", t, func() {
    Convey("Status code should be 200", func() {
      So(w.Code, ShouldEqual, 200)
    })
		Convey("Response should be a QueryResult", func() {
			type QueryResult struct {
			  Result []types.Player  `json:"result"`
			  TotalNum int        `json:"total_number"`
			  Before int          `json:"before"`        //-1 means this is the first page
			  Current int         `json:"current"`
			  Next int            `json:"next"`          //-1 means this is the last page
			}
			var result QueryResult
			json.Unmarshal(w.Body.Bytes(), &result)
			exist := false
			for _, one := range result.Result {
				if one.Equals(&mocked_player) {
					exist = true
					break
				}
			}
      So(exist, ShouldBeTrue)
    })
	})


	//update
	org_player := mocked_player
	mocked_player.Name = fmt.Sprintf("%s_%d", mocked_player.Name, time.Now().UnixNano())
	b, _ = json.Marshal(&mocked_player)
	r, _ = http.NewRequest("POST", "/api/player", bytes.NewBuffer(b))
	r.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_PlayersController_Insert_Query_Update(Update)", "Code[%d]\n%s", w.Code,
		w.Body.String())

	Convey("Subject: Test Player Insert, Query and Update Endpoint\n", t, func() {
		Convey("Status code should be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Response should be a QueryResult", func() {
			type Success struct {
				Succ bool `json:"success"`
				Id string	`json:"id"`
			}
			var result Success
			json.Unmarshal(w.Body.Bytes(), &result)
      So(result.Succ, ShouldBeTrue)
			So(mocked_player.ObjId.Hex() == result.Id, ShouldBeTrue)
		})
	})


	//check
	r, _ = http.NewRequest("GET", "/api/player?page=0&num=20", nil)
	r.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test_PlayersController_Insert_Query_Update(Check)", "Code[%d]\n%s", w.Code,
		w.Body.String())

	Convey("Subject: Test Player Insert, Query and Update Endpoint\n", t, func() {
		Convey("Status code should be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Response should be a QueryResult", func() {
			type QueryResult struct {
			  Result []types.Player  `json:"result"`
			  TotalNum int        `json:"total_number"`
			  Before int          `json:"before"`        //-1 means this is the first page
			  Current int         `json:"current"`
			  Next int            `json:"next"`          //-1 means this is the last page
			}
			var result QueryResult
			json.Unmarshal(w.Body.Bytes(), &result)
			success := false
			for _, one := range result.Result {
				if one.Equals(&mocked_player) {
					success = true
					break
				}
			}
			for _, one := range result.Result {
				if one.Equals(&org_player) {
					success = false
					break
				}
			}
			So(success, ShouldBeTrue)
		})
	})
}
