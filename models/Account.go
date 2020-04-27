package models

import (
	"errors"
	"farming_backend/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// User 用户信息表
type User struct {
	UserID   int    `orm:"pk;auto;unique;column(user_id)" json:"uid"`
	UserName string `orm:"column(u_name);size(100)" json:"username"`
	Nick     string `orm:"column(nick);size(100)" json:"nick"`
	Password string `orm:"column(password);size(100)" json:"-"`
	Avatar   string `orm:"column(avatar);size(1000)" json:"avatar"`
	//enable 账号可用 disable账号不可用
	Status     string    `orm:"column(status);size(50)" json:"-"`
	CreateTime time.Time `orm:"type(datetime);column(create_time);auto_now_add" json:"-"`
}

// TableName 获取对应数据库表名.
func (user *User) TableName() string {

	return "users"
}

// TableUnique 多字段唯一键 唯一索引
func (user *User) TableUnique() [][]string {
	return [][]string{
		{"user_id", "u_name"},
	}
}

// TableEngine 获取数据使用的引擎.
func (user *User) TableEngine() string {
	return "INNODB"
}

// NewUser 新用户
func NewUser() *User {
	return &User{}
}

var (
	//ErrUserNoExist 用户不存在
	ErrUserNoExist = errors.New("用户不存在")
	//ErrAuthMethodInvalid 登录验证方式无效
	ErrAuthMethodInvalid = errors.New("登录验证方式无效")
	// ErrPassword 登录异常
	ErrPassword = errors.New("登录异常")
	defaultNick = "萌新" + utils.SixRandomNums()
)

const (
	defaultStatus = "enable"
	defaultAvatar = "http://e.hiphotos.baidu.com/nuomi/pic/item/d000baa1cd11728b395ef96ec0fcc3cec3fd2c0e.jpg"
)

// Login 登录
func (user *User) Login(uName string, password string) (*User, error) {
	o := orm.NewOrm()
	newUser := NewUser()
	// 验证用户是否存在
	err := o.Raw("select * from users where u_name = ? and status = 'enable' limit 1;", uName).QueryRow(newUser)
	if err != nil {
		logs.Error("用户登录 ->", err)
		return newUser, ErrUserNoExist
	}
	// 验证用户登录
	ok, err := utils.PasswordVerify(newUser.Password, password)
	if ok && err == nil {
		return newUser, nil
	}

	return newUser, ErrPassword
}

// Register 注册
func (user *User) Register() (int64, error) {
	o := orm.NewOrm()

	// 验证用户账号是否存在
	var list orm.ParamsList
	num, err := o.Raw("SELECT *  FROM users WHERE u_name = ?", user.UserName).ValuesFlat(&list)

	if err == nil && num > 0 {
		return 0, errors.New("用户账号已存在")
	}
	// md5后再md5
	user.Password = utils.CryptoPassword(user.Password)
	user.Avatar = defaultAvatar
	user.Nick = defaultNick
	user.Status = defaultStatus
	// 这里如果&user 会报错
	uid, err := o.Insert(user)

	if err != nil {
		beego.Error("保存用户数据失败 =>", err)
		return 0, errors.New("保存用户失败")
	}
	return uid, nil
}

// SaveOrUpdateUserInfo 保存更新用户信息
func (user *User) SaveOrUpdateUserInfo(cols ...string) error {
	o := orm.NewOrm()
	// err := o.Raw("select * from users where user_id = ? and status = 'enable' limit 1;", u.UserId).QueryRow(user)

	if _, err := o.Update(user, cols...); err != nil {
		beego.Error("保存用户信息失败=>", err)
		return errors.New("保存用户信息失败")
	}
	return nil
}

// GetUserInfo 获取用户信息
func (user *User) GetUserInfo() (*User, error) {
	newUser := NewUser()
	o := orm.NewOrm()
	err := o.Raw("select * from users where user_id = ? limit 1;", user.UserID).QueryRow(newUser)
	if err != nil {
		// logs.Error("获取用户信息 ->", err)
		return newUser, ErrUserNoExist
	}
	return newUser, nil
}
