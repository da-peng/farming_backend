package main

import (
	_ "farming_backend/routers"

	"farming_backend/manages"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	manages.RegisterDataBase()

	manages.RegisterModel()
	manages.RegisterLogger()

	beego.Run()

}
