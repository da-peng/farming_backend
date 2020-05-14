package routers

import (
	"farming_backend/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/api/login", &controllers.AccountController{}, "post:Login")
	beego.Router("/api/user/get", &controllers.AccountController{}, "post:GetUserInfo")
	beego.Router("/api/register", &controllers.AccountController{}, "post:Register")
	beego.Router("/api/user/update", &controllers.AccountController{}, "post:SaveOrUpdateUserInfo")
	beego.Router("/api/crewler/vegetable", &controllers.CrewlerController{}, "post:SpiderVegetablePrice")

	beego.Router("/api/vegetable/list", &controllers.VegetableController{}, "get:GetVegetableList")
	beego.Router("/api/vegetable/price", &controllers.VegetableController{}, "post:GetVegetablePrice")
	// beego.Router("/statistics/vegetable/list", &controllers)
	// beego.Router("/statistics/weather/list", &controllers)
	// beego.Router("/statistics/today", &controllers)
}
