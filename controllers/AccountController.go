package controllers

import (
	"encoding/json"
	"farming_backend/models"

	"github.com/astaxie/beego"
)

// AccountController 用户登录与注册
type AccountController struct {
	BaseController
}

// User 用户登录信息
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// UserInfo 用户信息
type UserInfo struct {
	UserID   int    `json:"uid"`
	UserName string `json:"u_name"`
	Nick     string `json:"nick"`
	Avatar   string `json:"avatar"`
}

const (
	success = "操作成功"
)

// Login 用户登录
func (a *AccountController) Login() {
	var user User
	// a.BaseController.Prepare()
	if a.Ctx.Input.IsPost() {
		data := a.Ctx.Input.RequestBody
		err := json.Unmarshal(data, &user)
		if err != nil {
			a.ResponseJSON(20001, "请求数据异常")
		}
		user, err := models.NewUser().Login(user.UserName, user.Password)
		if err == nil {
			a.ResponseJSON(20002, success, user)
		} else {
			beego.Error("用户登录 ->", err)
			a.ResponseJSON(20003, "账号或密码错误")
		}

	}
}

// GetUserInfo 获取用户信息
func (a *AccountController) GetUserInfo() {
	var userInfo UserInfo
	// a.BaseController.Prepare()
	if a.Ctx.Input.IsPost() {
		data := a.Ctx.Input.RequestBody
		err := json.Unmarshal(data, &userInfo)
		if err != nil {
			a.ResponseJSON(20001, "请求数据异常")
		}
		newUser := models.NewUser()
		var user *models.User

		if userInfo.UserID != 0 {
			newUser.UserID = userInfo.UserID
			user, err = newUser.GetUserInfo()
			if err != nil {
				beego.Error("查询用户 ->", err)
				a.ResponseJSON(20004, "用户不存在")
			}
		}
		if userInfo.UserName != "" {
			newUser.UserName = userInfo.UserName
			user, err = newUser.GetUserInfo()
			if err != nil {
				beego.Error("查询用户 ->", err)
				a.ResponseJSON(20004, "用户不存在")
			}
		}
		a.ResponseJSON(20004, success, user)

	}

}

// Register 注册
func (a *AccountController) Register() {
	var user User
	// a.BaseController.Prepare()
	if a.Ctx.Input.IsPost() {
		body := a.Ctx.Input.RequestBody

		err := json.Unmarshal(body, &user)
		if err != nil {
			a.ResponseJSON(20001, "请求数据异常")
		}
		newUser := models.NewUser()
		newUser.UserName = user.UserName
		newUser.Password = user.Password
		var uid int64
		if uid, err = newUser.Register(); err != nil {
			beego.Error("用户注册 ->", err)
			a.ResponseJSON(20004, "注册失败，请联系系统管理员处理")
		}
		data := map[string]int64{"uid": uid}
		a.ResponseJSON(20004, "注册成功", data)
	}
}

//SaveOrUpdateUserInfo 保存更新用户信息
func (a *AccountController) SaveOrUpdateUserInfo() {
	var userInfo UserInfo
	// a.BaseController.Prepare()
	if a.Ctx.Input.IsPost() {
		data := a.Ctx.Input.RequestBody
		err := json.Unmarshal(data, &userInfo)
		if err != nil {
			a.ResponseJSON(20001, "请求数据异常")
		}
		// DVO 转 DTO 再更新数据库
		newUser := models.NewUser()
		newUser.UserID = userInfo.UserID
		// 验证用户是否存在
		newUser, err = newUser.GetUserInfo()
		if err != nil {
			beego.Error("查询用户 ->", err)
			a.ResponseJSON(20004, "更新失败，请联系系统管理员处理")
		}
		var cols []string
		if userInfo.Nick != "" {
			newUser.Nick = userInfo.Nick
			cols = append(cols, "nick")
		}
		if userInfo.Avatar != "" {
			newUser.Avatar = userInfo.Avatar
			cols = append(cols, "avatar")
		}

		if err := newUser.SaveOrUpdateUserInfo(cols...); err != nil {
			beego.Error("用户更新信息 ->", err)
			a.ResponseJSON(20004, "更新失败，请联系系统管理员处理")
		}

		a.ResponseJSON(20005, success)

	}
}
