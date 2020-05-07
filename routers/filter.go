package routers

import (
	"encoding/json"
	"farming_backend/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func errorRequest(ctx *context.Context, errcode int, message string) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["errcode"] = errcode
	jsonData["message"] = message
	returnJSON, _ := json.Marshal(jsonData)
	ctx.ResponseWriter.Write(returnJSON)
}

func filterUser(ctx *context.Context) {
	token := ctx.Input.Header("token")
	beego.Info("用户token ->", token)
	if token == "" {
		errorRequest(ctx, 20401, "请登录后再操作")
		return
	} else {
		user := models.NewUser()
		userID, _ := strconv.Atoi(token)
		user.UserID = userID
		u, err := user.GetUserInfo()
		if err != nil {
			errorRequest(ctx, 20401, "请登录后再操作")
			return
		}
		beego.Info("用户 ->", u.UserName)
	}

	if ctx.Input.IsPost() {
		requestBody := ctx.Input.RequestBody
		m := make(map[string]interface{})
		err := json.Unmarshal(requestBody, &m)
		if err != nil {
			errorRequest(ctx, 20402, "请求数据异常")
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
		errorRequest(ctx, 20403, "请求数据异常")
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
	beego.InsertFilter("/crewler", beego.BeforeRouter, filterUser)
	beego.InsertFilter("/crewler/*", beego.BeforeRouter, filterUser)
	beego.InsertFilter("/vegetable", beego.BeforeRouter, filterUser)
	beego.InsertFilter("/vegetable/*", beego.BeforeRouter, filterUser)

}
