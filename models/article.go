package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Category struct {
	Id   int    `orm:"column(id);auto" json:"id,omitempty"`
	Name string `orm:"column(name);size(128);unique" json:"name,omitempty"`
	// 不使用 reverse(many)，避免在 RunSyncdb 时出现问题
	// Article []*Article `orm:"reverse(many)" json:"-"`
}

type Tag struct {
	Id   int    `orm:"auto" json:"id,omitempty"`
	Name string `orm:"size(50);unique" json:"name,omitempty"`
	// 不使用 reverse(many)，避免在 RunSyncdb 时出现问题
	// Article []*Article `orm:"reverse(many)" json:"-"`
}

type Article struct {
	Id       int    `orm:"column(id);auto" json:"id"`
	Title    string `orm:"column(title);size(255)" json:"title"`
	Abstract string `orm:"column(abstract);size(255);null" json:"abstract,omitempty"`
	Content  string `orm:"column(content);type(longtext)" json:"content"`
	Author   string `orm:"column(author);size(100)" json:"author"`
	//Category   *Category `orm:"rel(fk);null" json:"category,omitempty"` // 外键关系，允许为 null
	//Tags       []*Tag    `orm:"rel(m2m)" json:"tags,omitempty"`         // 多对多关系
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"createTime"`
}

func init() {
	orm.RegisterModel(
		new(Category),
		new(Tag),
		new(Article),
	)
}

// get all articles
func GetAllArticles() ([]Article, error) {

	o := orm.NewOrm()

	var articles []Article

	_, err := o.QueryTable("article").All(&articles)

	return articles, err
}

// get article by id
func GetArticleById(id int) (*Article, error) {
	o := orm.NewOrm()
	article := Article{}
	err := o.QueryTable("article").Filter("id", id).One(&article)
	if err != nil {
		return nil, err
	}
	return &article, err

}

// add article
func ArticleInsert(article *Article) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(article)
}

// update article
func ArticleUpdate(article *Article) (int64, error) {
	o := orm.NewOrm()
	return o.Update(article)
}

// delete article
func ArticleDelete(id int) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("article").Filter("id", id).Delete()
}

// get article by page
func GetArticleByPage(page int, pageSize int) ([]Article, int64, error) {
	o := orm.NewOrm()
	var articles []Article
	//查询总数
	total, err := o.QueryTable("article").Count()
	if err != nil {
		return nil, 0, err
	}
	//分页查询
	_, err = o.QueryTable("article").OrderBy("-id").
		Limit(pageSize, (page-1)*pageSize).
		All((&articles))
	if err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}
