package controllers

import (
	"fmt"
	"encoding/json"

	"github.com/astaxie/beego"

	"controllers/errors"
	"models"
	"models/types"
)

type PlayersController struct {
	beego.Controller
}

/**
GET: query players
params: page, num
result: json
error happened:
{
	"error": "Error message",
	"error_code": $code,
	"request": "/api/player?page=$page&num=$num"
}
success:
{
	"result": [{}, {}, ...],
	"total_number": 2,
  "before": 0,
  "current": 1,
  "next": 0
}
*/
func (this *PlayersController) Get() {
	page, err := this.GetInt("page", 0)
	if err != nil {
		beego.Warn("Get players error, parse page error: ", err)
		mErr := errors.Issue("Parse page error",
			errors.E_TYPE_SERVICE + errors.E_MODULE_PLAYER + errors.E_DETAIL_PARSE_PARAM_ERROR,
			this.Ctx.Request.URL.String())
		this.Data["json"] = &mErr
		this.ServeJson()
		return
	}

	num, err := this.GetInt("num", 10)
	if err != nil {
		beego.Warn("Get players error, parse num error: ", err)
		mErr := errors.Issue("Parse num error",
			errors.E_TYPE_SERVICE + errors.E_MODULE_PLAYER + errors.E_DETAIL_PARSE_PARAM_ERROR,
			this.Ctx.Request.URL.String())
		this.Data["json"] = &mErr
		this.ServeJson()
		return
	}

	beego.Info(fmt.Sprintf("Query player, req: %s, page: %d, num: %d",
		this.Ctx.Request.URL.String(), page, num))

	if page < 0 || num < 0 {
		mErr := errors.Issue("Page and num must be non-negative",
			errors.E_TYPE_SERVICE + errors.E_MODULE_PLAYER + errors.E_DETAIL_ILLEGAL_PARAM,
			this.Ctx.Request.URL.String())
		this.Data["json"] = &mErr
		this.ServeJson()
		return
	}

	this.Data["json"] = models.QueryPlayer(page, num)
	this.ServeJson()
}

/**
POST: insert or update player(ObjId not empty means update)
request body: player info(json)
result: json
success:
{
	"success": true,
	"id": $ObjId
}
*/
func (this *PlayersController) Post() {
	var player types.Player
	json.Unmarshal(this.Ctx.Input.RequestBody, &player)
	if !player.Valid() {
		mErr := errors.Issue("Player info incomplete",
			errors.E_TYPE_SERVICE + errors.E_MODULE_PLAYER + errors.E_DETAIL_INCOMPLETE_DATA,
			this.Ctx.Request.URL.String())
		this.Data["json"] = &mErr
		this.ServeJson()
		return
	}

	var err error
	if player.ObjId != "" {
		err = models.UpdatePlayer(&player)
	} else {
		err = models.InsertPlayer(&player)
	}

	if err != nil {
		mErr := errors.Issue("Database operate error",
			errors.E_TYPE_SERVICE + errors.E_MODULE_PLAYER + errors.E_DETAIL_INTERNAL_DB_ERROR,
			this.Ctx.Request.URL.String())
		this.Data["json"] = &mErr
		this.ServeJson()
		return
	}

	type Success struct {
		Succ bool `json:"success"`
		Id string	`json:"id"`
	}
	//this.Data["json"] = &Success {true, player.ObjId.Hex()}
	this.Data["json"] = &player
	this.ServeJson()
}
