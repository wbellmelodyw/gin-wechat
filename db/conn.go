package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github/wbellmelodyw/gin-wechat/config"
	"xorm.io/core"
)

var dns string
var weChatDB *xorm.Engine

//程序开始执行
func init() {
	dns = fmt.Sprintf("root:a890418@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		"mysql-master",
		"3306",
		"wechat_todo",
		"utf8mb4")

	if err := initEngine; err != nil {
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
	weChatDB.Logger().SetLevel(core.LogLevel(logLevel))
	return
}
