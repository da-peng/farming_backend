package controllers

import (
	"encoding/json"
	"io"

	"farming_backend/models"

	"github.com/astaxie/beego"
)

// BaseController .
type BaseController struct {
	beego.Controller
	User *models.User
}

// 判断用户是否登录.
func (a *BaseController) isUserLoggedIn() bool {
	return a.User != nil && a.User.UserID > 0
}

// Prepare 请求预处理
func (a *BaseController) Prepare() {

}

// ResponseJSON 响应
func (a *BaseController) ResponseJSON(statusCode int, msg string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["statuscode"] = statusCode
	jsonData["message"] = msg
	if len(data) > 0 && data[0] != nil {
		jsonData["data"] = data[0]
	}
	responseJSON, err := json.Marshal(jsonData)
	if err != nil {
		beego.Error(err)
	}

	a.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	a.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")

	io.WriteString(a.Ctx.ResponseWriter, string(responseJSON))
	a.StopRun()

}
