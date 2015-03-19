package controllers

import (
  "encoding/base64"
  "encoding/json"

	"github.com/astaxie/beego"

	"controllers/errors"
  "models/types"
  "models"
)

type AuthController struct {
	beego.Controller
}

/**
GET: for redirect, return error message
{
  "ErrorMsg": "You need login first",
  "ErrorCode": "20007",
  "Request": "/api/player?page=0\u0026num=10"
}
*/
func (this *AuthController) Get() {
  beego.Info("req: ", this.GetString("req"))
  req, err := base64.URLEncoding.DecodeString(this.GetString("req"))
  var reqStr string
  if err != nil {
    reqStr = ""
  } else {
    reqStr = string(req)
  }
  this.Data["json"] = errors.Issue("You need login first",
    errors.E_TYPE_SERVICE + errors.E_MODULE_ANY + errors.E_DETAIL_NEED_LOGIN,
    reqStr)
  this.ServeJson()
}

/**
POST: for authentication
request body: username and password(json)
result: token is set in session
success:
{
  "success": true
}
fail:
{
  "success": false
}
*/
func (this *AuthController) Post() {
  var auth types.AuthInfo
  json.Unmarshal(this.Ctx.Input.RequestBody, &auth)

  type Success struct {
		Succ bool     `json:"success"`
	}
  if auth.Valid() && models.VerifyAdmin(auth.Username, auth.Password) {
    this.Data["json"] = &Success {true}
    this.SetSession("token", models.GetToken())
  } else {
    this.Data["json"] = &Success {false}
  }
  this.ServeJson()
}
