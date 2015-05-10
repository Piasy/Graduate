package controllers

import (
  "fmt"

	"github.com/astaxie/beego"

	"controllers/errors"
	"models"
)

type TrainRemarkController struct {
	beego.Controller
}

/**
GET: request the processed remark
params: player=$player&page=$page&num=$num
result:
{
	"result": [...],
	"total_number": 2,
  "before": 0,
  "current": 1,
  "next": 0
}
*/
func (this *TrainRemarkController) Get() {
  player := this.GetString("player")
  if len(player) != 40 {
    beego.Warn("Query train remark error, wrong params: player = ", player)
    this.Data["json"] = errors.Issue("Invalid player",
      errors.E_TYPE_SERVICE + errors.E_MODULE_TRAINREMARK + errors.E_DETAIL_ILLEGAL_PARAM,
      this.Ctx.Request.URL.String())
    this.ServeJson()
    return
  }

  page, err := this.GetInt("page", 0)
	if err != nil {
		beego.Warn("Get players error, parse page error: ", err)
		this.Data["json"] = errors.Issue("Parse page error",
			errors.E_TYPE_SERVICE + errors.E_MODULE_TRAINREMARK + errors.E_DETAIL_PARSE_PARAM_ERROR,
			this.Ctx.Request.URL.String())
		this.ServeJson()
		return
	}

	num, err := this.GetInt("num", 10)
	if err != nil {
		beego.Warn("Get players error, parse num error: ", err)
		this.Data["json"] = errors.Issue("Parse num error",
			errors.E_TYPE_SERVICE + errors.E_MODULE_TRAINREMARK + errors.E_DETAIL_PARSE_PARAM_ERROR,
			this.Ctx.Request.URL.String())
		this.ServeJson()
		return
	}

  this.Data["json"] = models.GetTrainRecord(player, page, num)
  beego.Info(fmt.Sprintf("Query train remark, req: %s, player: %s, page: %d, num: %d",
    this.Ctx.Request.URL.String(), player, page, num))

  this.ServeJson()
}
