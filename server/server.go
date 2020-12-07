package server

import (
	"github.com/Guiyunweb/shiki/conf"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"os"
	"blog-api/server/http"
)

func StartServer() {
	app := iris.New()
	app.Use(logger.New())
	// 跨域配置
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
	}))
	app.AllowMethods(iris.MethodOptions)

	// 读取路由
	http.LoadRouter(app)

	// 读取配置
	server := conf.LoadServer()
	err := app.Run(iris.Addr(server.HttpPost))
	if err != nil {
		os.Exit(-1)
	}
}
