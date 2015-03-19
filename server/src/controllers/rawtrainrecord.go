package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"controllers/errors"
	"models"
	"models/types"
)

type RawTrainRecordController struct {
	beego.Controller
}

/**
POST: collect train record data
body: json, two types, append and flush
{
  "op" : $op(append, flush)
  "player": $player
  "data": $data
}
result:
{
  "success": $succ
}
*/
func (this *RawTrainRecordController) Post() {
  type PostData struct {
    Op string                 `json:"op"`
    Player string             `json:"player"`
    Data types.RawTrainRecord `json:"data"`
  }
  type Success struct {
    Succ bool     `json:"success"`
  }

  var data PostData
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
  switch(data.Op) {
    case "append":
      models.AppendRawTrainData(data.Player, &data.Data)
      this.Data["json"] = &Success {true}
    case "flush":
      err := models.FlushRawTrainData()
      if err != nil {
        beego.Warn("Flush train data error, database error: ", err)
        this.Data["json"] = errors.Issue("Flush raw train data failed",
          errors.E_TYPE_SERVICE + errors.E_MODULE_TRAINRECORD + errors.E_DETAIL_INTERNAL_DB_ERROR,
          this.Ctx.Request.URL.String())
      } else {
        this.Data["json"] = &Success {true}
      }
    default:
      this.Data["json"] = errors.Issue("parameter op needed",
        errors.E_TYPE_SERVICE + errors.E_MODULE_TRAINRECORD + errors.E_DETAIL_ILLEGAL_PARAM,
        this.Ctx.Request.URL.String())
  }
  this.ServeJson()
}
