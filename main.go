package main

import (
	"farming_backend/manages"
	_ "farming_backend/routers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func main() {
	manages.RegisterDataBase()

	manages.RegisterModel()
	manages.RegisterLogger()
	beego.Run()

}
