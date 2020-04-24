package models

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//CrewlerInfo 爬虫列表
type CrewlerInfo struct {
	CrewlerName  string `orm:"column(crewler_name);size(100)" json:"crewlerName"`
	Config       string `orm:"column(config);size(100)" json:"config"`
	LastExecTime string `orm:"type(datetime);column(create_time);size(100)" json:"lastExecTime"`
	CreateTime   string `orm:"type(datetime);column(create_time);size(100);auto_now_add" json:"-"`
}

// TableName 表名
func (v *CrewlerInfo) TableName() string {
	return "vegetable_prize"
}

// TableEngine 获取数据使用的引擎.
func (v *CrewlerInfo) TableEngine() string {
	return "INNODB"
}

// NewCrewlerInfo 新数据
func NewCrewlerInfo() *CrewlerInfo {

	return &CrewlerInfo{}
}

// SaveOrUpdateCrewlerInfo  保存或更新 爬虫配置信息
func (v *CrewlerInfo) SaveOrUpdateCrewlerInfo(cols ...string) (int64, error) {
	o := orm.NewOrm()
	// 判断是否存在
	num, err := o.QueryTable(v.TableName()).Filter("crewlerName", v.CrewlerName).Count()

	if err == nil && num > 0 {
		nums, err := o.Update(v, cols...)
		if err != nil {
			beego.Error("保存用户信息失败=>", err)
			return 0, errors.New("保存用户信息失败")
		}
		return nums, nil
	}

	if err == nil && num == 0 {
		id, err := o.Insert(v)
		if err != nil {
			beego.Error("保存用户数据失败 =>", err)
			return 0, errors.New("保存用户失败")
		}
		return id, nil
	}

	return 0, err

}

// GetCrewlerInfo  保存或更新 爬虫配置信息
func (v *CrewlerInfo) GetCrewlerInfo() (*CrewlerInfo, error) {
	o := orm.NewOrm()
	// 判断是否存在
	err := o.Raw("SELECT * FROM vegetable_prize WHERE crewler_name = ?", v.CrewlerName).QueryRow(v)

	if err != nil {
		return nil, errors.New("此爬虫不存在")
	}
	return v, nil

}
