package routers

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func filterUser(ctx *context.Context) {
	token := ctx.Input.Header("token")
	beego.Info("用户token ->", token)
	if token == "" {
		jsonData := make(map[string]interface{}, 3)
		jsonData["errcode"] = 20401
		jsonData["message"] = "请登录后再操作"
		returnJSON, _ := json.Marshal(jsonData)
		ctx.ResponseWriter.Write(returnJSON)
		return
	}

	if ctx.Input.IsPost() {
		requestBody := ctx.Input.RequestBody
		m := make(map[string]interface{})
		err := json.Unmarshal(requestBody, &m)
		if err != nil {
			jsonData := make(map[string]interface{}, 3)
			jsonData["errcode"] = 20402
			jsonData["message"] = "请求数据异常"
			returnJSON, _ := json.Marshal(jsonData)
			ctx.ResponseWriter.Write(returnJSON)
		}
	}
}

func filterRequestInfo(ctx *context.Context) {
	requestBody := ctx.Input.RequestBody
	beego.Info("请求接口 ->", ctx.Input.URL())
	beego.Info("请求参数 ->", string(requestBody))
}

func filterContentType(ctx *context.Context) {
	contentType := ctx.Input.Header("Content-Type")
	if !strings.Contains(strings.ToLower(contentType), "application/json") {
		jsonData := make(map[string]interface{}, 3)
		jsonData["errcode"] = 20403
		jsonData["message"] = "请求数据异常"
		returnJSON, _ := json.Marshal(jsonData)
		ctx.ResponseWriter.Write(returnJSON)
	}
}

func finishRouter(ctx *context.Context) {
	ctx.ResponseWriter.Header().Add("Farming-Version", "1.0")
	ctx.ResponseWriter.Header().Add("Farming-Site", "https://www.suwan.club")
}

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, filterRequestInfo)
	beego.InsertFilter("*", beego.BeforeRouter, filterContentType)
	beego.InsertFilter("/user", beego.BeforeRouter, filterUser)
	beego.InsertFilter("/user/*", beego.BeforeRouter, filterUser)
	beego.InsertFilter("/statistics", beego.BeforeRouter, filterUser)
	beego.InsertFilter("/statistics/*", beego.BeforeRouter, filterUser)
}
