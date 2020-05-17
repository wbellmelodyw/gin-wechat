package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github/wbellmelodyw/gin-wechat/config"
	"xorm.io/core"
)

var dns string
var weChatDB *xorm.Engine

//程序开始执行
func init() {
	dns = fmt.Sprintf("root:890418@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		"172.21.0.3", //"mysql-master",
		"3306",
		"wechat_todo",
		"utf8mb4")

	if err := initEngine(); err != nil {
		panic(err)
	}

	if err := weChatDB.Ping(); err != nil {
		panic(err)
	}
}

func initEngine() (err error) {
	weChatDB, err = xorm.NewEngine("mysql", dns)
	if err != nil {
		return
	}
	weChatDB.SetMaxIdleConns(2)
	weChatDB.SetMaxOpenConns(10)

	showSQL := config.GetBool("show_sql")
	logLevel := config.MustInt("log_level", 1)

	weChatDB.ShowSQL(showSQL)
	//weChatDB.SetLogger(logger.Module("db").Sugar())
	weChatDB.Logger().SetLevel(core.LogLevel(logLevel))
	return
}
