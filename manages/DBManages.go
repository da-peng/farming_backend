package manages

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"farming_backend/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// RegisterDataBase 初始化数据库
func RegisterDataBase() {
	beego.Info("正在初始化数据库配置.")
	orm.DefaultTimeLoc = time.Local
	// mysql 配置
	host := beego.AppConfig.String("db_host")
	database := beego.AppConfig.String("db_database")
	username := beego.AppConfig.String("db_username")
	password := beego.AppConfig.String("db_password")
	port := beego.AppConfig.String("db_port")
	timezone := beego.AppConfig.String("timezone")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s", username, password, host, port, database, url.QueryEscape(timezone))
	// 设置最大空闲连接
	// 设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	if err := orm.RegisterDataBase("default", "mysql", dataSource, maxIdle, maxConn); err != nil {
		beego.Error("注册默认数据库失败->", err)
		os.Exit(1)
	}

}

// RegisterModel 注册模型
func RegisterModel() {
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.VegetablePrize))
	// 参数三：true更新表
	// 第二个参数 是否强制更新 当为 true的时候，会执行drop table后再建表 ，（每次启动项目会将原先的表和数据都删除
	orm.RunSyncdb("default", false, true)
}

// RegisterLogger 注册日志
func RegisterLogger() {

}
