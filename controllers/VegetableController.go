package controllers

import (
	"encoding/json"
	"farming_backend/models"

	"github.com/astaxie/beego"
)

// VegetableController 爬虫控制器
type VegetableController struct {
	BaseController
}

// VegetableSearch 爬虫范围
type VegetableSearch struct {
	VegetableName string `json:"vegetable"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
}

// GetVegetablePrice 获取蔬菜价格
func (v *VegetableController) GetVegetablePrice() {
	var vegetableSearch VegetableSearch
	if v.Ctx.Input.IsPost() {
		data := v.Ctx.Input.RequestBody
		err := json.Unmarshal(data, &vegetableSearch)
		if err != nil {
			v.ResponseJSON(20001, "请求数据异常")
		}
		vegetablePriceList, err := models.NewVegetablePrize().GetVegetablePriceList(
			vegetableSearch.VegetableName,
			vegetableSearch.StartTime,
			vegetableSearch.EndTime,
		)
		if err != nil {
			beego.Info("蔬菜价格 ->", err)
			v.ResponseJSON(20002, "操作失败，数据异常")
		}
		v.ResponseJSON(20000, success, vegetablePriceList)
	}
}

// GetVegetableList 获取蔬菜名称列表
func (v *VegetableController) GetVegetableList() {

	if v.Ctx.Input.IsGet() {
		vegetableList, err := models.NewVegetablePrize().GetVegetableList()
		if err != nil {
			beego.Info("蔬菜列表 ->", err)
			v.ResponseJSON(20002, "操作失败，数据异常")
		}
		v.ResponseJSON(20000, success, vegetableList)

	}
}
