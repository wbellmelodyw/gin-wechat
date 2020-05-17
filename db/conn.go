package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github/wbellmelodyw/gin-wechat/config"
	"xorm.io/core"
)

var dns string
var WeChat *xorm.Engine

//程序开始执行
func init() {
	dns = fmt.Sprintf("root:890418@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		"mysql-master", //"172.21.0.3",
		"3306",
		"wechat_todo",
		"utf8mb4")

	if err := initEngine(); err != nil {
		panic(err)
	}

	if err := WeChat.Ping(); err != nil {
		panic(err)
	}
}

func initEngine() (err error) {
	WeChat, err = xorm.NewEngine("mysql", dns)
	if err != nil {
		return
	}
	WeChat.SetMaxIdleConns(2)
	WeChat.SetMaxOpenConns(10)

	showSQL := config.GetBool("show_sql")
	logLevel := config.MustInt("log_level", 1)

	WeChat.ShowSQL(showSQL)
	//weChatDB.SetLogger(logger.Module("db").Sugar())
	WeChat.Logger().SetLevel(core.LogLevel(logLevel))
	return
}
