package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"service/internal/domain/info/do"
)

var Db *gorm.DB

func init() {
	Db = NewDb()
}

func NewDb() *gorm.DB {
	username := "root"   //账号
	password := "qiuyue" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "test"     //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect mysql server error=%s", err.Error()))
	}
	db.AutoMigrate(&do.InfoDo{})
	return db
}
