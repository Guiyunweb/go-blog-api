package conf

import (
	"blog-api/cache"
	"blog-api/model"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Info() {

	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()

}
