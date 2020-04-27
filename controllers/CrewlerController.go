package controllers

import (
	"encoding/json"
	"farming_backend/crewler"
	"strings"
)

// CrewlerController 爬虫控制器
type CrewlerController struct {
	BaseController
}

// CrewlerRange 爬虫范围
type CrewlerRange struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// SpiderVegetablePrice  爬取所有蔬菜价格
func (c *CrewlerController) SpiderVegetablePrice() {
	var crewlerRange CrewlerRange
	if c.Ctx.Input.IsPost() {
		data := c.Ctx.Input.RequestBody
		err := json.Unmarshal(data, &crewlerRange)
		if err != nil {
			c.ResponseJSON(20001, "请求数据异常")
		}
		startTime := strings.Split(crewlerRange.StartTime, " ")[0][2:]
		endTime := strings.Split(crewlerRange.EndTime, " ")[0][2:]
		crewler.VegetableCrewler(startTime, endTime)

		c.ResponseJSON(20004, "请求成功")

	}
}
