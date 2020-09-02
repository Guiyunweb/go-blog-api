package model

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

var DB *xorm.Engine

// Database 在中间件中初始化mysql链接
func Database(connString string) {

	engine, err := xorm.NewEngine("mysql", connString)
	if err != nil {
		panic(err)
	}

	// 启动日志
	engine.ShowSQL(true)

	// 设置空闲数大小
	engine.SetMaxIdleConns(20)

	//最大打开连接数
	engine.SetMaxOpenConns(20)

	//最大生存时间
	engine.SetConnMaxLifetime(time.Second * 30)

	DB = engine
}
