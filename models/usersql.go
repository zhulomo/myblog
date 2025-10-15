package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Users struct {
	Id         int       `orm:"auto"`
	Username   string    `orm:"column(username);size(64)"`
	Password   string    `orm:"column(password);size(64)"`
	Status     int       `orm:"column(status);size(4)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Users))
}

// 通过用户名查询
func GetUserByName(username string) (*Users, error) {
	o := orm.NewOrm()

	//err := o.Read(&user, "Username")
	//orm.ErrNoRows 一个全局变量，值是一个错误值。 使用orm查询数据时，没有记录则返回这个错误
	//err == orm.ErrNoRows err是errors.errorString,不是orm.ErrNoRows
	// if err.Error() == "no row found" {
	// 	//数据库没有该用户名
	// 	return nil, nil
	// } else if err != nil {
	// 	fmt.Println(err.Error())
	// 	fmt.Println("o.Read err type:", reflect.TypeOf(err), "err:", err)
	// 	return nil, err
	exist := o.QueryTable("users").Filter("username", username).Exist()
	if exist {
		user := Users{Username: username}
		err := o.Read(&user, "Username")
		if err != nil {
			return nil, err // 这里 err 应该几乎不会出现
		}
		return &user, nil
	}

	return nil, nil

}
func GetUserPasswordByName(username string) (string, error) {
	var result []orm.Params
	o := orm.NewOrm()
	num, err := o.QueryTable("users").Filter("Username", username).Values(&result, "Password")
	if err == nil && num > 0 {
		password := result[0]["Password"].(string)

		return password, nil

	}
	return "", err
}

func InsertUser(user *Users) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(user)
}
