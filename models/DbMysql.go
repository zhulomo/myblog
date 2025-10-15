package models

import (
	"log"

	"github.com/beego/beego/v2/client/orm"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//aliasName数据库别名
	aliasName := "default"
	driverName := "mysql"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	username := "root"
	password := "123456"
	dbName := "myBlog"

	//注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//创建链接
	dbConn := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	orm.RegisterDataBase(aliasName, driverName, dbConn)

	o := orm.NewOrm()

	if _, err := o.Raw("select 1").Exec(); err != nil {
		log.Println("数据库链接出错", err)

	}
	log.Println("数据库链接成功")

}

// func CreatTableUser() {
// 	sql := `CREATE TABLE  IF NOT EXISTS users(
// 		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
// 		username VARCHAR(64),
// 		password VARCHAR(64),
// 		status INT(4),
// 		createTime DATETIME
// 		);`

// 	ModifyDB(sql)
// }

// func ModifyDB(query string, args ...interface{}) (orm.Result, error)  {
// 	o := orm.NewOrm()
// 	res, err := o.Raw(query).Exec()
// 	if err != nil {
// 		log.Println("执行sql出错", query)
// 		return res, err
// 	}
// 	log.Println("执行成功", query)
// 	return res, err
// }
