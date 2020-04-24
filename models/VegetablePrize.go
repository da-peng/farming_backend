package models

import "github.com/astaxie/beego/orm"

//VegetablePrize 蔬菜价格
type VegetablePrize struct {
	// ID         int    `orm:"pk;auto;unique;column(user_id)" json:"id"`
	Vegetable  string `orm:"column(vegetable);size(100)" json:"vegetable"`
	Origin     string `orm:"column(origin);size(100)" json:"origin"`
	AvgPrice   string `orm:"column(AvgPrice);size(100)" json:"avgPrice"`
	DateTime   string `orm:"type(datetime);column(dateTime);size(100)" json:"dateTime"`
	CreateTime string `orm:"type(datetime);column(create_time);size(100);auto_now_add" json:"-"`
}

// TableName 表名
func (v *VegetablePrize) TableName() string {
	return "vegetable_prize"
}

// TableEngine 获取数据使用的引擎.
func (v *VegetablePrize) TableEngine() string {
	return "INNODB"
}

// NewVegetablePrize 新数据
func NewVegetablePrize() *VegetablePrize {

	return &VegetablePrize{}
}

// AddMulti 添加
func AddMulti(vs []VegetablePrize) (int64, error) {
	o := orm.NewOrm()
	//第一个参数 bulk 为并列插入的数量，第二个为对象的slice
	successNums, err := o.InsertMulti(20, vs)
	if err != nil {
		return 0, err
	}
	return successNums, nil
}
