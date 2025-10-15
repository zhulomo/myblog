package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Article struct {
	Id         int       `orm:"auto"`
	Title      string    `orm:"column(title);size(64)"`
	Abstract   string    `orm:"column(abstract);size(64)"`
	Content    string    `orm:"column(content)type(text)"`
	Author     string    `orm:"column(author);size(64)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Article))
}

func GetAllArticles() ([]Article, error) {

	o := orm.NewOrm()

	var articles []Article

	_, err := o.QueryTable("article").All(&articles)

	return articles, err
}
