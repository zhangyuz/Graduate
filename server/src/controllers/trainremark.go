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
params: player=$player&type=$type&page=$page&num=$num
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
  page, err := this.GetInt("page", 0)
  if err != nil {
    beego.Warn("Query train remark error, parse page error: ", err)
    this.Data["json"] = errors.Issue("Parse page error",
      errors.E_TYPE_SERVICE + errors.E_MODULE_PLAYER + errors.E_DETAIL_PARSE_PARAM_ERROR,
      this.Ctx.Request.URL.String())
    this.ServeJson()
    return
  }

  num, err := this.GetInt("num", 10)
  if err != nil {
    beego.Warn("Query train remark error, parse num error: ", err)
    this.Data["json"] = errors.Issue("Parse num error",
      errors.E_TYPE_SERVICE + errors.E_MODULE_PLAYER + errors.E_DETAIL_PARSE_PARAM_ERROR,
      this.Ctx.Request.URL.String())
    this.ServeJson()
    return
  }

  if page < 0 || num < 0 {
    beego.Warn("Query train remark error, wrong params: page, num = ", page, num)
    this.Data["json"] = errors.Issue("Page and num must be non-negative",
      errors.E_TYPE_SERVICE + errors.E_MODULE_PLAYER + errors.E_DETAIL_ILLEGAL_PARAM,
      this.Ctx.Request.URL.String())
    this.ServeJson()
    return
  }

  mType := this.GetString("type")
  player := this.GetString("player")

  switch(mType) {
    case "speed":
      this.Data["json"] = models.GetTrainSpeed(player, page, num)
    case "distance":
      this.Data["json"] = models.GetTrainDistance(player, page, num)
    default:
      this.Data["json"] = errors.Issue("Invalid type parameter",
        errors.E_TYPE_SERVICE + errors.E_MODULE_TRAINREMARK + errors.E_DETAIL_ILLEGAL_PARAM,
        this.Ctx.Request.URL.String())
  }

  beego.Info(fmt.Sprintf("Query train remark, req: %s, player: %s, type: %s, page: %d, num: %d",
    this.Ctx.Request.URL.String(), player, mType, page, num))

  this.ServeJson()
}