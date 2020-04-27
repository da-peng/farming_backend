package routers

import (
	"farming_backend/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/login", &controllers.AccountController{}, "post:Login")
	beego.Router("/user/get", &controllers.AccountController{}, "post:GetUserInfo")
	beego.Router("/register", &controllers.AccountController{}, "post:Register")
	beego.Router("/user/update", &controllers.AccountController{}, "post:SaveOrUpdateUserInfo")

	beego.Router("/crewler/vegetable", &controllers.CrewlerController{}, "post:SpiderVegetablePrice")
	// beego.Router("/statistics/vegetable/list", &controllers)
	// beego.Router("/statistics/weather/list", &controllers)
	// beego.Router("/statistics/today", &controllers)
}
