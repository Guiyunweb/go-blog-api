package model

import (
	"github.com/Guiyunweb/shiki/conf"
	"github.com/Guiyunweb/shiki/log"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var db *xorm.Engine

func Init() {
	var err error
	// 读取配置文件配置
	dbConf := conf.LoadDatabase()
	engine, err := xorm.NewEngine(dbConf.DbType, dbConf.DefaultDsn)
	if err != nil {
		log.Panic(err)
	}

	engine.ShowSQL(true)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)
	//engine.SetConnMaxLifetime(100)
	//配置检测连接是否有效
	engine.Query("SELECT 1 FROM DUAL")
	db = engine
}

func CloseDB() {
	defer db.Close()
}
