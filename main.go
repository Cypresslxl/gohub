package main

import (
	"flag"
	"fmt"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {
	// 配置初始化，依赖命令行 --env 参数

	var env string
	// StringVar defines a string flag with specified name, default value, and usage string.
	// The argument p points to a string variable in which to store the value of the flag.
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	// fmt.Println("---env --- %s: ", env)
	flag.Parse()
	config.InitConfig(env)

	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化日志库
	bootstrap.SetupLogger()
	// 初始化数据库
	bootstrap.SetupDB()
	// 初始化 Redis
	bootstrap.SetupRedis()
	// 初始化路由绑定
	bootstrap.SetupRoute(router)
	// logger.Dump(captcha.NewCaptcha().VerifyCaptcha("9hZ0bCaMBdc0oCAFj0fy", "242703"), "正确的答案")
	// logger.Dump(captcha.NewCaptcha().VerifyCaptcha("4EAztsuaTDrotxaUjoEg", "000000"), "错误的答案")

	//test
	// fmt.Println("name  : %s", config.GetString("app.name"))
	// fmt.Println("level : %d", config.GetInt("app.level"))
	// fmt.Println("debug : %s", config.GetBool("app.debug"))

	// 运行服务
	err := router.Run(":" + config.GetString("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
