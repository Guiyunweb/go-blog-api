package main

import (
	"blog-api/conf"
	"blog-api/server"
)

func main() {
	// 从配置文件读取配置
	conf.Info()

	// 装载路由
	r := server.NewRouter()
	r.Run(":8866")
}
