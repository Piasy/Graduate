package controllers

import (
  "fmt"

	"github.com/astaxie/beego"

	"controllers/errors"
	"models"
)

type TrainHistoryController struct {
	beego.Controller
}

/**
GET: request the history train data(processed remark)
params: player=$player&page=$page&num=$num
result:
{
	"result": [{}, {}, ...],
	"total_number": 2,
  "before": 0,
  "current": 1,
  "next": 0
}
*/
func (this *TrainHistoryController) Get() {
  page, err := this.GetInt("page", 0)
  if err != nil {
    beego.Warn("Query train history error, parse page error: ", err)
    this.Data["json"] = errors.Issue("Parse page error",
      errors.E_TYPE_SERVICE + errors.E_MODULE_PLAYER + errors.E_DETAIL_PARSE_PARAM_ERROR,
      this.Ctx.Request.URL.String())
    this.ServeJson()
    return
  }

  num, err := this.GetInt("num", 10)
  if err != nil {
    beego.Warn("Query train history error, parse num error: ", err)
    this.Data["json"] = errors.Issue("Parse num error",
      errors.E_TYPE_SERVICE + errors.E_MODULE_PLAYER + errors.E_DETAIL_PARSE_PARAM_ERROR,
      this.Ctx.Request.URL.String())
    this.ServeJson()
    return
  }

  if page < 0 || num < 0 {
    beego.Warn("Query train history error, wrong params: page, num = ", page, num)
    this.Data["json"] = errors.Issue("Page and num must be non-negative",
      errors.E_TYPE_SERVICE + errors.E_MODULE_PLAYER + errors.E_DETAIL_ILLEGAL_PARAM,
      this.Ctx.Request.URL.String())
    this.ServeJson()
    return
  }

  player := this.GetString("player")

  this.Data["json"] = models.QueryTrainRecord(player, page, num)

  beego.Info(fmt.Sprintf("Query train history, req: %s, player: %s, page: %d, num: %d",
    this.Ctx.Request.URL.String(), player, page, num))

  this.ServeJson()
}
