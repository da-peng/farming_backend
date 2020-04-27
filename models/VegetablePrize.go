package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//VegetablePrize 蔬菜价格
type VegetablePrize struct {
	ID         int       `orm:"pk;auto;unique;column(id)" json:"-"`
	Vegetable  string    `orm:"column(vegetable);size(100)" json:"vegetable"`
	Origin     string    `orm:"column(origin);size(100)" json:"origin"`
	AvgPrice   string    `orm:"column(AvgPrice);size(100)" json:"avgPrice"`
	DateTime   time.Time `orm:"type(datetime);column(dateTime)" json:"dateTime"`
	CreateTime time.Time `orm:"type(datetime);column(create_time);auto_now_add" json:"-"`
}

// TableName 表名
func (v *VegetablePrize) TableName() string {
	return "vegetable_prize"
}

// TableEngine 获取数据使用的引擎.
func (v *VegetablePrize) TableEngine() string {
	return "INNODB"
}

// // TableIndex 多字段索引
// func (v *VegetablePrize) TableIndex() [][]string {
// 	return [][]string{
// 		{"vegetable", "dateTime"},
// 	}
// }

// TableUnique 多字段唯一键 唯一索引
func (v *VegetablePrize) TableUnique() [][]string {
	return [][]string{
		{"vegetable", "dateTime"},
	}
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

// GetVegetablePriceList 获取蔬菜价格信息
func GetVegetablePriceList(vegetableName string, startTime string, endTime string) {

}

// GetVegetableList 获取蔬菜名列表
func GetVegetableList() {

}
