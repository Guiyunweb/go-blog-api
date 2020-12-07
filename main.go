package main

import (
	"blog-api/cmd"
	"github.com/Guiyunweb/shiki/conf"
)

func main() {
	// 读取配置
	if err := conf.Init(); err != nil {
		panic(err)
	}
	// 启动服务
	cmd.Run()
}
