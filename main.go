package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/app/cmd"
	"gohub/app/cmd/make"
	"gohub/bootstrap"
	"gohub/pkg/console"
	"os"

	btsConfig "gohub/config"
	"gohub/pkg/config"
)

func init() {
	//Load the configuration information in the 'config' directory
	btsConfig.Initialize()
}

func main() {
	//The main entry of the application. By default, the cmd.CmdServe command is invoked
	var rootCmd = &cobra.Command{
		Use:   "Playground",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		// 会被子命令继承的前置 Run : PersistentPreRun func(cmd *Command, args []string)
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// Registration global parameter configuration initialization, dependent on the command line --env 参数
			config.InitConfig(cmd.Env)

			//Init Logger,Database,Redis,Cache
			bootstrap.SetupLogger()
			bootstrap.SetupDB()
			bootstrap.SetupRedis()
			bootstrap.SetupCache()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
		cmd.CmdCache,
		//cmd.CmdTestCommand,
	)

	// 配置默认运行 Web 服务
	//The reason why the cmd.CmdServe command is set as the default command to be executed when no subcommand is provided is due to the following line of code:
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe) //这里设置了默认的执行命令

	// Register global parameter，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	//Execute uses the args (os.Args[1:] by default) and run through the command tree finding appropriate matches for commands and then corresponding flags.
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}

// 配置初始化，依赖命令行 --env 参数

//var env string
//// StringVar defines a string flag with specified name, default value, and usage string.
//// The argument p points to a string variable in which to store the value of the flag.
//flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
//// fmt.Println("---env --- %s: ", env)
//flag.Parse()
//config.InitConfig(env)
//
//// new 一个 Gin Engine 实例
//router := gin.New()
//
//// 初始化日志库
//bootstrap.SetupLogger()
//// 初始化数据库
//bootstrap.SetupDB()
//// 初始化 Redis
//bootstrap.SetupRedis()
//// 初始化路由绑定
//bootstrap.SetupRoute(router)

//router.GET("/test_auth", middlewares.AuthJWT(), func(c *gin.Context) {
//	userModel := auth.CurrentUser(c)
//	response.Data(c, userModel)
//})

//router.GET("/test_guest", middlewares.GuestJWT(), func(c *gin.Context) {
//	c.String(http.StatusOK, "Hello guest")
//})

// logger.Dump(captcha.NewCaptcha().VerifyCaptcha("9hZ0bCaMBdc0oCAFj0fy", "242703"), "正确的答案")
// logger.Dump(captcha.NewCaptcha().VerifyCaptcha("4EAztsuaTDrotxaUjoEg", "000000"), "错误的答案")

// sms.NewSMS().Send("15679195296", sms.Message{
// Template: config.GetString("sms.aliyun.template_code"),
// Data:     map[string]string{"code": "123456"},
// })
// verifycode.NewVerifyCode().SendSMS("15679195296")

//test
// fmt.Println("name  : %s", config.GetString("app.name"))
// fmt.Println("level : %d", config.GetInt("app.level"))
// fmt.Println("debug : %s", config.GetBool("app.debug"))

// 运行服务
//err := router.Run(":" + config.GetString("app.port"))
//if err != nil {
//	// 错误处理，端口被占用了或者其他错误
//	fmt.Println(err.Error())
//}
//}
