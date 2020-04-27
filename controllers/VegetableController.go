package controllers

import "encoding/json"

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

	}
}
