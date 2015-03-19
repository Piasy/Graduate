package routers

import (
	"encoding/base64"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"controllers"
	"models"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, loginChecker)
  beego.Router("/api/player", &controllers.PlayersController{})
	beego.Router("/api/auth", &controllers.AuthController{})
	beego.Router("/api/rawtrainrecord", &controllers.RawTrainRecordController{})
}

var loginChecker = func(ctx *context.Context) {
	token := ctx.Input.Session("token")

	if ctx.Request.URL.Path != "/api/auth" && token != models.GetToken() {
		ctx.Redirect(302, "/api/auth?req=" +
			base64.URLEncoding.EncodeToString([]byte(ctx.Request.URL.String())))
	}
}
