package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego"
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
	successNums, err := o.InsertMulti(1, vs)
	if err != nil {
		return 0, err
	}
	return successNums, nil
}

// GetVegetablePriceList 获取蔬菜价格信息
func (v *VegetablePrize) GetVegetablePriceList(vegetableName string, StartTime string, EndTime string) ([]*VegetablePrize, error) {
	o := orm.NewOrm()
	var vegetablePrize []*VegetablePrize

	_, err1 := time.Parse("2006-01-02 15:04:05", StartTime)
	_, err2 := time.Parse("2006-01-02 15:04:05", EndTime)
	if err1 != nil || err2 != nil {
		beego.Info("数据不是时间类型")
		return nil, errors.New("数据不是时间类型")
	}

	num, err := o.Raw("SELECT * FROM vegetable_prize WHERE vegetable = ? And dateTime > ? And dateTime < ?", vegetableName, StartTime, EndTime).QueryRows(&vegetablePrize)
	if err != nil {
		beego.Info("Returned Rows Num: %d, %s", num, err)
		return nil, err
	}
	return vegetablePrize, nil
}

//VegetablePrizeDTO 蔬菜价格
type VegetablePrizeDTO struct {
	Vegetable []string `json:"vegetable"`
}

// GetVegetableList 获取蔬菜名列表
func (v *VegetablePrize) GetVegetableList() (*VegetablePrizeDTO, error) {
	vegetablePrize := &VegetablePrizeDTO{}
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)

	logDay := yesTime.Format("2006-01-02")

	o := orm.NewOrm()
	var list orm.ParamsList
	// num, err := o.QueryTable(v.TableName()).Filter("dateTime", logDay+" 08:00:00").All(&vegetablePrize, "vegetable")
	num, err := o.Raw("SELECT vegetable  FROM vegetable_prize WHERE dateTime = ?", logDay+" 08:00:00").ValuesFlat(&list)
	if err != nil {
		beego.Info("Returned Rows Num: %d, %s", num, err)
		return nil, err
	}

	if err == nil && num > 0 {
		beego.Info(list)
	}
	var ret []string

	for _, i := range list {
		j := convertString(i)
		ret = append(ret, string(j))
	}
	// beego.Info(ret)

	vegetablePrize.Vegetable = ret

	return vegetablePrize, nil
}

func convertString(e interface{}) string {
	var s string
	switch v := e.(type) {
	case int:
		// fmt.Println("整型", v)
		s = strconv.Itoa(v)
		// fmt.Println(s)
	case string:
		s = v
	}
	return s
}
