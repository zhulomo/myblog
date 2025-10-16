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

func GetArticleById(id int) (*Article, error) {
	o := orm.NewOrm()
	article := Article{}
	err := o.QueryTable("article").Filter("id", id).One(&article)
	if err != nil {
		return nil, err
	}
	return &article, err

}

func ArticleInsert(article *Article) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(article)
}

func ArticleUpdate(article *Article) (int64, error) {
	o := orm.NewOrm()
	return o.Update(article)
}
func ArticleDelete(article *Article) (int64, error) {
	o := orm.NewOrm()
	return o.Delete(article)
}
